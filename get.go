package genv

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var (
	ErrNotSet             = errors.New("environment variable not set")
	ErrNotPointer         = errors.New("T is not a pointer")
	ErrNotPointerToStruct = errors.New("T is not a pointer to a struct")
	ErrCannotSetField     = errors.New("cannot set field, field must be exported")
	ErrUnsupportedType    = errors.New("unsupported type")
	ErrCannotCast         = errors.New("environment variable cannot be cast to target type")
)

func Get[T any](key string) (value T, err error) {
	raw, ok := os.LookupEnv(key)
	if !ok {
		return value, fmt.Errorf("genv: environment variable not set: %q", key)
	}

	return cast[T](key, raw)
}

func GetOrDefault[T any](key string, fallback T) (value T) {
	value, err := Get[T](key)
	if err != nil {
		return fallback
	}
	return value
}

func GetOrPanic[T any](key string) (value T) {
	value, err := Get[T](key)
	if err != nil {
		panic(err)
	}
	return value
}

// Use:
//
//	type Config struct {
//	    StringValue string `genv:"STRING_KEY"`
//	    IntValue    int    `genv:"INT_KEY"`
//	    BoolValue   bool   `genv:"BOOL_KEY"`
//	}
//
//	var cfg Config
//
//	if err := GetStruct(&cfg); err != nil {
//	   ...
//	}
func GetStruct[T any](value *T) (err error) {
	typ := reflect.TypeOf(value)
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("genv: %w", ErrNotPointer)
	}

	el := typ.Elem()
	if el.Kind() != reflect.Struct {
		return fmt.Errorf("genv: %w", ErrNotPointerToStruct)
	}

	val := reflect.ValueOf(value).Elem()

	for i := range el.NumField() {
		field := el.Field(i)
		key := field.Tag.Get("genv")
		if key == "" {
			continue
		}

		fieldVal := val.Field(i)
		if !fieldVal.CanSet() {
			return fmt.Errorf("genv: %w: %s.%s", ErrCannotSetField, el.Name(), field.Name)
		}

		switch fieldVal.Kind() {
		case reflect.String:
			s, err := Get[string](key)
			if err != nil {
				return err
			}
			fieldVal.SetString(s)
		case reflect.Bool:
			b, err := Get[bool](key)
			if err != nil {
				return err
			}
			fieldVal.SetBool(b)
		case reflect.Int:
			i, err := Get[int](key)
			if err != nil {
				return err
			}
			fieldVal.SetInt(int64(i))
		case reflect.Float64:
			f, err := Get[float64](key)
			if err != nil {
				return err
			}
			fieldVal.SetFloat(f)
		}
	}

	return nil
}

// TODO: extend types
func cast[T any](key, raw string) (value T, err error) {
	switch t := any(value).(type) {
	case string:
		return any(raw).(T), nil
	case bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return value, fmt.Errorf("genv: %w: %q %T", ErrCannotCast, key, t)
		}
		return any(b).(T), nil
	case int:
		i, err := strconv.Atoi(raw)
		if err != nil {
			return value, fmt.Errorf("genv: %w: %q %T", ErrCannotCast, key, t)
		}
		return any(i).(T), nil
	case float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return value, fmt.Errorf("genv: %w: %q %T", ErrCannotCast, key, t)
		}
		return any(f).(T), nil
	default:
		return value, fmt.Errorf("genv: %w: %T", ErrUnsupportedType, t)
	}
}
