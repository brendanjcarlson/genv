package genv

import (
	"os"
	"testing"
)

func Test_Get(t *testing.T) {
	t.Run("invalid key", func(t *testing.T) {
		_, err := Get[string]("INVALID_KEY")
		if err == nil {
			t.Errorf("expected non nil err")
		}
	})

	t.Run("string key", func(t *testing.T) {
		var key string = "TEST_GET_STRING_KEY"
		var want string = "string"

		os.Setenv(key, want)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[string](key)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %s\ngot %s", want, got)
		}
	})

	t.Run("bool key", func(t *testing.T) {
		var key string = "TEST_GET_BOOL_KEY"
		var value string = "true"
		var want bool = true

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[bool](key)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %t\ngot %t", want, got)
		}
	})

	t.Run("int key", func(t *testing.T) {
		var key string = "TEST_GET_INT_KEY"
		var value string = "123"
		var want int = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[int](key)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %d\ngot %d", want, got)
		}
	})

	t.Run("float64 key", func(t *testing.T) {
		var key string = "TEST_GET_FLOAT64_KEY"
		var value string = "12.3"
		var want float64 = 12.3

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[float64](key)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %f\ngot %f", want, got)
		}
	})
}

func Test_GetOrDefault(t *testing.T) {
	os.Setenv("TEST_GET_OR_DEFAULT", "test_get_or_default")
	t.Cleanup(func() {
		os.Unsetenv("TEST_GET_OR_DEFAULT")
	})

	t.Run("returns value", func(t *testing.T) {
		got := GetOrDefault("TEST_GET_OR_DEFAULT", "fallback")
		if "test_get_or_default" != got {
			t.Errorf("want %s\ngot %s", "test_get_or_default", got)
		}
	})

	t.Run("returns fallback", func(t *testing.T) {
		got := GetOrDefault("INVALID_KEY", "fallback")
		if "fallback" != got {
			t.Errorf("want %s\ngot %s", "fallback", got)
		}
	})
}

func Test_GetOrPanic(t *testing.T) {
	os.Setenv("TEST_GET_OR_PANIC", "test_get_or_panic")
	t.Cleanup(func() {
		os.Unsetenv("TEST_GET_OR_PANIC")
	})

	t.Run("does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("should not have panicked")
			}
		}()
		got := GetOrPanic[string]("TEST_GET_OR_PANIC")
		if "test_get_or_panic" != got {
			t.Errorf("want %s\ngot %s", "test_get_or_panic", got)
		}
	})

	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("should have panicked")
			}
		}()
		GetOrPanic[string]("INVALID KEY")
	})
}

func Test_GetStruct(t *testing.T) {
	var stringKey string = "TEST_GET_STRUCT_STRING_KEY"
	var stringValue string = "string_value"
	var stringWant string = "string_value"

	var boolKey string = "TEST_GET_STRUCT_BOOL_KEY"
	var boolValue string = "true"
	var boolWant bool = true

	var intKey string = "TEST_GET_STRUCT_INT_KEY"
	var intValue string = "123"
	var intWant int = 123

	var float64Key string = "TEST_GET_STRUCT_FLOAT64_KEY"
	var float64Value string = "12.3"
	var float64Want float64 = 12.3

	os.Setenv(stringKey, stringValue)
	os.Setenv(boolKey, boolValue)
	os.Setenv(intKey, intValue)
	os.Setenv(float64Key, float64Value)
	t.Cleanup(func() {
		os.Unsetenv(stringKey)
		os.Unsetenv(boolKey)
		os.Unsetenv(intKey)
		os.Unsetenv(float64Key)
	})

	type Config struct {
		StringValue  string  `genv:"TEST_GET_STRUCT_STRING_KEY"`
		BoolValue    bool    `genv:"TEST_GET_STRUCT_BOOL_KEY"`
		IntValue     int     `genv:"TEST_GET_STRUCT_INT_KEY"`
		Float64Value float64 `genv:"TEST_GET_STRUCT_FLOAT64_KEY"`
	}

	wantCfg := Config{
		StringValue:  stringWant,
		BoolValue:    boolWant,
		IntValue:     intWant,
		Float64Value: float64Want,
	}

	var gotCfg Config
	err := GetStruct(&gotCfg)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if gotCfg != wantCfg {
		t.Errorf("want %+v\ngot %+v", wantCfg, gotCfg)
	}
}

func Test_cast(t *testing.T) {
	t.Run("ok string", func(t *testing.T) {
		var input string = "input"
		var want string = "input"

		got, err := cast[string]("", input)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %s\ngot %s", want, got)
		}
	})

	t.Run("ok bool", func(t *testing.T) {
		var input string = "true"
		var want bool = true

		got, err := cast[bool]("", input)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %v\ngot %v", want, got)
		}
	})

	t.Run("not ok bool", func(t *testing.T) {
		var input string = "not a bool"

		_, err := cast[bool]("", input)
		if err == nil {
			t.Errorf("expected err, got nil\n")
		}
	})

	t.Run("ok int", func(t *testing.T) {
		var input string = "123"
		var want int = 123

		got, err := cast[int]("", input)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %v\ngot %v", want, got)
		}
	})

	t.Run("not ok int", func(t *testing.T) {
		var input string = "not an int"

		_, err := cast[int]("", input)
		if err == nil {
			t.Errorf("expected err, got nil\n")
		}
	})

	t.Run("ok float", func(t *testing.T) {
		var input string = "12.3"
		var want float64 = 12.3

		got, err := cast[float64]("", input)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if want != got {
			t.Errorf("want %v\ngot %v", want, got)
		}
	})

	t.Run("not ok float", func(t *testing.T) {
		var input string = "not an float64"

		_, err := cast[float64]("", input)
		if err == nil {
			t.Errorf("expected err, got nil\n")
		}
	})
}
