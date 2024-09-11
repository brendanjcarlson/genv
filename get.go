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

// Get retrieves an environment variable from the current process.
//
// Returns the value cast to the given type parameter or an error if
// the variable is not set or cannot be cast to the given type.
//
// All simple types are supported.
//
// See GetStruct for loading variables into a struct.
//
// Use:
//
//	secret, err := Get[string]("SECRET_KEY")
//	if errors.Is(err, ErrNotSet) {
//	   ...
//	} else if errors.Is(err, ErrCannotCast) {
//	   ...
//	}
//
//	timeoutMillis, err := Get[int]("TIMEOUT_MILLIS")
//	if errors.Is(err, ErrNotSet) {
//	   ...
//	} else if errors.Is(err, ErrCannotCast) {
//	   ...
//	}
func Get[T any](key string) (value T, err error) {
	raw, ok := os.LookupEnv(key)
	if !ok {
		return value, fmt.Errorf("genv: %w: %q", ErrNotSet, key)
	}

	return cast[T](key, raw)
}

// GetOrDefault retrieves an environment variable from the current process.
//
// Returns the value cast to the given type parameter or the fallback value
// if the variable is not set or cannot be cast to the given type.
//
// Use:
//
//	secret := GetOrDefault[string]("SECRET_KEY", "super-secret-key")
//	timeoutMillis := GetOrDefault[int]("TIMEOUT_MILLIS", 500)
func GetOrDefault[T any](key string, fallback T) (value T) {
	value, err := Get[T](key)
	if err != nil {
		return fallback
	}
	return value
}

// GetOrPanic retrieves an environment variable from the current process.
//
// Returns the value cast to the given type parameter or panics
// if the variable is not set or cannot be cast to the given type.
//
// Use:
//
//	secret := GetOrPanic[string]("SECRET_KEY")
func GetOrPanic[T any](key string) (value T) {
	value, err := Get[T](key)
	if err != nil {
		panic(err)
	}
	return value
}

// GetStruct retrieves an environment variable from the current process
// and loads them into a struct.
//
// Attemps to cast and assign values to struct fields annotated by the tag:
//
// `genv:"KEY_NAME"`
//
// Returns an error if the variable is not set, the argument is not a struct,
// the tagged field is not exported, or if the variable cannot be cast to the given type.
//
// Use:
//
//	type Config struct {
//	    StringValue string `genv:"STRING_KEY"`
//	    IntValue    int    `genv:"INT_KEY"`
//	    BoolValue   bool   `genv:"BOOL_KEY"`
//	}
//
//	var cfg Config
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
