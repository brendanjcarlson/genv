package genv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Get[T any](key string) (value T, err error) {
	raw, ok := os.LookupEnv(key)
	if !ok {
		return value, fmt.Errorf("genv: environment variable not set: %q", key)
	}

	return cast[T](key, raw)
}

// TODO: extend types
func cast[T any](key, raw string) (value T, err error) {
	switch t := any(value).(type) {
	case string:
		return any(raw).(T), nil
	case int:
		i, err := strconv.Atoi(raw)
		if err != nil {
			return value, fmt.Errorf("genv: environment variable %q cannot be cast to %T", key, t)
		}
		return any(i).(T), nil
	case bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return value, fmt.Errorf("genv: environment variable %q cannot be cast to %T", key, t)
		}
		return any(b).(T), nil
	case []byte:
		return any([]byte(raw)).(T), nil
	default:
		return value, fmt.Errorf("genv: unsupported type: %T", t)
	}
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
		return fmt.Errorf("genv: T is not a pointer")
	}

	el := typ.Elem()
	if el.Kind() != reflect.Struct {
		return fmt.Errorf("genv: T is not a pointer to a struct")
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
			return fmt.Errorf("genv: cannot set field %s", field.Name)
		}

		switch fieldVal.Kind() {
		case reflect.String:
			v, err := Get[string](key)
			if err != nil {
				return err
			}
			fieldVal.SetString(v)
		case reflect.Int:
			v, err := Get[int](key)
			if err != nil {
				return err
			}
			fieldVal.SetInt(int64(v))
		case reflect.Bool:
			v, err := Get[bool](key)
			if err != nil {
				return err
			}
			fieldVal.SetBool(v)
		}
	}

	return nil
}
