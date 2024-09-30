package genv

import (
	"errors"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("invalid key", func(t *testing.T) {
		var want error = ErrNotSet

		_, got := Get[string]("INVALID_KEY")
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("string key", func(t *testing.T) {
		var key string = "TEST_GET_STRING_KEY"
		var want string = "string"

		os.Setenv(key, want)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[string](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %sgot %s", want, got)
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
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %tgot %t", want, got)
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
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("int8 key", func(t *testing.T) {
		var key string = "TEST_GET_INT8_KEY"
		var value string = "123"
		var want int8 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[int8](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("int16 key", func(t *testing.T) {
		var key string = "TEST_GET_INT16_KEY"
		var value string = "123"
		var want int16 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[int16](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("int32 key", func(t *testing.T) {
		var key string = "TEST_GET_INT32_KEY"
		var value string = "123"
		var want int32 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[int32](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("int64 key", func(t *testing.T) {
		var key string = "TEST_GET_INT64_KEY"
		var value string = "123"
		var want int64 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[int64](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("uint key", func(t *testing.T) {
		var key string = "TEST_GET_UINT_KEY"
		var value string = "123"
		var want uint = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[uint](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("uint8 key", func(t *testing.T) {
		var key string = "TEST_GET_UINT8_KEY"
		var value string = "123"
		var want uint8 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[uint8](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("uint16 key", func(t *testing.T) {
		var key string = "TEST_GET_UINT16_KEY"
		var value string = "123"
		var want uint16 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[uint16](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("uint32 key", func(t *testing.T) {
		var key string = "TEST_GET_UINT32_KEY"
		var value string = "123"
		var want uint32 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[uint32](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("uint64 key", func(t *testing.T) {
		var key string = "TEST_GET_UINT64_KEY"
		var value string = "123"
		var want uint64 = 123

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[uint64](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %dgot %d", want, got)
		}
	})

	t.Run("float32 key", func(t *testing.T) {
		var key string = "TEST_GET_FLOAT32_KEY"
		var value string = "12.3"
		var want float32 = 12.3

		os.Setenv(key, value)
		t.Cleanup(func() { os.Unsetenv(key) })

		got, err := Get[float32](key)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %fgot %f", want, got)
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
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %fgot %f", want, got)
		}
	})
}

func TestGetOrDefault(t *testing.T) {
	os.Setenv("TEST_GET_OR_DEFAULT", "test_get_or_default")
	t.Cleanup(func() { os.Unsetenv("TEST_GET_OR_DEFAULT") })

	t.Run("returns value", func(t *testing.T) {
		got := GetOrDefault("TEST_GET_OR_DEFAULT", "fallback")
		if "test_get_or_default" != got {
			t.Fatalf("want %sgot %s", "test_get_or_default", got)
		}
	})

	t.Run("returns fallback", func(t *testing.T) {
		got := GetOrDefault("INVALID_KEY", "fallback")
		if "fallback" != got {
			t.Fatalf("want %sgot %s", "fallback", got)
		}
	})
}

func TestGetOrPanic(t *testing.T) {
	os.Setenv("TEST_GET_OR_PANIC", "test_get_or_panic")
	t.Cleanup(func() { os.Unsetenv("TEST_GET_OR_PANIC") })

	t.Run("does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should not have panicked")
			}
		}()
		got := GetOrPanic[string]("TEST_GET_OR_PANIC")
		if "test_get_or_panic" != got {
			t.Fatalf("want %sgot %s", "test_get_or_panic", got)
		}
	})

	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("should have panicked")
			}
		}()
		GetOrPanic[string]("INVALID KEY")
	})
}

func TestGetStruct(t *testing.T) {
	var stringKey string = "TEST_GET_STRUCT_STRING_KEY"
	var stringValue string = "string_value"
	var stringWant string = "string_value"

	var boolKey string = "TEST_GET_STRUCT_BOOL_KEY"
	var boolValue string = "true"
	var boolWant bool = true

	var intKey string = "TEST_GET_STRUCT_INT_KEY"
	var intValue string = "123"
	var intWant int = 123

	var int8Key string = "TEST_GET_STRUCT_INT8_KEY"
	var int8Value string = "123"
	var int8Want int8 = 123

	var int16Key string = "TEST_GET_STRUCT_INT16_KEY"
	var int16Value string = "123"
	var int16Want int16 = 123

	var int32Key string = "TEST_GET_STRUCT_INT32_KEY"
	var int32Value string = "123"
	var int32Want int32 = 123

	var int64Key string = "TEST_GET_STRUCT_INT64_KEY"
	var int64Value string = "123"
	var int64Want int64 = 123

	var uintKey string = "TEST_GET_STRUCT_UINT_KEY"
	var uintValue string = "123"
	var uintWant uint = 123

	var uint8Key string = "TEST_GET_STRUCT_UINT8_KEY"
	var uint8Value string = "123"
	var uint8Want uint8 = 123

	var uint16Key string = "TEST_GET_STRUCT_UINT16_KEY"
	var uint16Value string = "123"
	var uint16Want uint16 = 123

	var uint32Key string = "TEST_GET_STRUCT_UINT32_KEY"
	var uint32Value string = "123"
	var uint32Want uint32 = 123

	var uint64Key string = "TEST_GET_STRUCT_UINT64_KEY"
	var uint64Value string = "123"
	var uint64Want uint64 = 123

	var float64Key string = "TEST_GET_STRUCT_FLOAT64_KEY"
	var float64Value string = "12.3"
	var float64Want float64 = 12.3

	os.Setenv(stringKey, stringValue)
	os.Setenv(boolKey, boolValue)
	os.Setenv(intKey, intValue)
	os.Setenv(int8Key, int8Value)
	os.Setenv(int16Key, int16Value)
	os.Setenv(int32Key, int32Value)
	os.Setenv(int64Key, int64Value)
	os.Setenv(uintKey, uintValue)
	os.Setenv(uint8Key, uint8Value)
	os.Setenv(uint16Key, uint16Value)
	os.Setenv(uint32Key, uint32Value)
	os.Setenv(uint64Key, uint64Value)
	os.Setenv(float64Key, float64Value)
	t.Cleanup(func() {
		os.Unsetenv(stringKey)
		os.Unsetenv(boolKey)
		os.Unsetenv(intKey)
		os.Unsetenv(int8Key)
		os.Unsetenv(int16Key)
		os.Unsetenv(int32Key)
		os.Unsetenv(int64Key)
		os.Unsetenv(uintKey)
		os.Unsetenv(uint8Key)
		os.Unsetenv(uint16Key)
		os.Unsetenv(uint32Key)
		os.Unsetenv(uint64Key)
		os.Unsetenv(float64Key)
	})

	type Config struct {
		StringValue  string  `genv:"TEST_GET_STRUCT_STRING_KEY"`
		BoolValue    bool    `genv:"TEST_GET_STRUCT_BOOL_KEY"`
		IntValue     int     `genv:"TEST_GET_STRUCT_INT_KEY"`
		Int8Value    int8    `genv:"TEST_GET_STRUCT_INT8_KEY"`
		Int16Value   int16   `genv:"TEST_GET_STRUCT_INT16_KEY"`
		Int32Value   int32   `genv:"TEST_GET_STRUCT_INT32_KEY"`
		Int64Value   int64   `genv:"TEST_GET_STRUCT_INT64_KEY"`
		UintValue    uint    `genv:"TEST_GET_STRUCT_UINT_KEY"`
		Uint8Value   uint8   `genv:"TEST_GET_STRUCT_UINT8_KEY"`
		Uint16Value  uint16  `genv:"TEST_GET_STRUCT_UINT16_KEY"`
		Uint32Value  uint32  `genv:"TEST_GET_STRUCT_UINT32_KEY"`
		Uint64Value  uint64  `genv:"TEST_GET_STRUCT_UINT64_KEY"`
		Float64Value float64 `genv:"TEST_GET_STRUCT_FLOAT64_KEY"`
	}

	t.Run("ok", func(t *testing.T) {
		wantCfg := Config{
			StringValue:  stringWant,
			BoolValue:    boolWant,
			IntValue:     intWant,
			Int8Value:    int8Want,
			Int16Value:   int16Want,
			Int32Value:   int32Want,
			Int64Value:   int64Want,
			UintValue:    uintWant,
			Uint8Value:   uint8Want,
			Uint16Value:  uint16Want,
			Uint32Value:  uint32Want,
			Uint64Value:  uint64Want,
			Float64Value: float64Want,
		}

		var gotCfg Config
		err := GetStruct(&gotCfg)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if gotCfg != wantCfg {
			t.Fatalf("want %+vgot %+v", wantCfg, gotCfg)
		}
	})

	t.Run("not a struct pointer", func(t *testing.T) {
		var notAStruct string
		var want error = ErrNotPointerToStruct

		got := GetStruct(&notAStruct)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})
}

func TestGetStructNested(t *testing.T) {
	var serverHostKey string = "TEST_GET_STRUCT_NESTED_SERVERHOST"
	var serverHostValue string = "127.0.0.1"
	var serverHostWant string = "127.0.0.1"

	var serverPortKey string = "TEST_GET_STRUCT_NESTED_SERVERPORT"
	var serverPortValue string = "8080"
	var serverPortWant string = "8080"

	var databaseHostKey string = "TEST_GET_STRUCT_NESTED_DATABASE_HOST"
	var databaseHostValue string = "localhost"
	var databaseHostWant string = "localhost"

	var databasePortKey string = "TEST_GET_STRUCT_NESTED_DATABASE_PORT"
	var databasePortValue string = "5342"
	var databasePortWant string = "5342"

	var databaseUserKey string = "TEST_GET_STRUCT_NESTED_DATABASE_USER"
	var databaseUserValue string = "postgres"
	var databaseUserWant string = "postgres"

	var databasePasswordKey string = "TEST_GET_STRUCT_NESTED_DATABASE_PASSWORD"
	var databasePasswordValue string = "password"
	var databasePasswordWant string = "password"

	var databaseDbnameKey string = "TEST_GET_STRUCT_NESTED_DATABASE_DBNAME"
	var databaseDbnameValue string = "test"
	var databaseDbnameWant string = "test"

	os.Setenv(serverHostKey, serverHostValue)
	os.Setenv(serverPortKey, serverPortValue)
	os.Setenv(databaseHostKey, databaseHostValue)
	os.Setenv(databasePortKey, databasePortValue)
	os.Setenv(databaseUserKey, databaseUserValue)
	os.Setenv(databasePasswordKey, databasePasswordValue)
	os.Setenv(databaseDbnameKey, databaseDbnameValue)
	t.Cleanup(func() {
		os.Unsetenv(serverHostKey)
		os.Unsetenv(serverPortKey)
		os.Unsetenv(databaseHostKey)
		os.Unsetenv(databasePortKey)
		os.Unsetenv(databaseUserKey)
		os.Unsetenv(databasePasswordKey)
		os.Unsetenv(databaseDbnameKey)
	})

	type ServerConfig struct {
		Host string `genv:"TEST_GET_STRUCT_NESTED_SERVERHOST"`
		Port string `genv:"TEST_GET_STRUCT_NESTED_SERVERPORT"`
	}

	type DatabaseConfig struct {
		Host     string `genv:"TEST_GET_STRUCT_NESTED_DATABASE_HOST"`
		Port     string `genv:"TEST_GET_STRUCT_NESTED_DATABASE_PORT"`
		User     string `genv:"TEST_GET_STRUCT_NESTED_DATABASE_USER"`
		Password string `genv:"TEST_GET_STRUCT_NESTED_DATABASE_PASSWORD"`
		Dbname   string `genv:"TEST_GET_STRUCT_NESTED_DATABASE_DBNAME"`
	}

	type AppConfig struct {
		Server ServerConfig
		DB     DatabaseConfig
	}

	t.Run("ok", func(t *testing.T) {
		wantCfg := AppConfig{
			Server: ServerConfig{
				Host: serverHostWant,
				Port: serverPortWant,
			},
			DB: DatabaseConfig{
				Host:     databaseHostWant,
				Port:     databasePortWant,
				User:     databaseUserWant,
				Password: databasePasswordWant,
				Dbname:   databaseDbnameWant,
			},
		}

		var gotCfg AppConfig
		err := GetStruct(&gotCfg)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if gotCfg != wantCfg {
			t.Fatalf("want %+vgot %+v", wantCfg, gotCfg)
		}
	})
}

func TestCast(t *testing.T) {
	t.Run("ok string", func(t *testing.T) {
		var input string = "input"
		var want string = "input"

		got, err := cast[string]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %sgot %s", want, got)
		}
	})

	t.Run("ok bool", func(t *testing.T) {
		var input string = "true"
		var want bool = true

		got, err := cast[bool]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok bool", func(t *testing.T) {
		var input string = "not a bool"
		var want error = ErrCannotCast

		_, got := cast[bool]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok int", func(t *testing.T) {
		var input string = "123"
		var want int = 123

		got, err := cast[int]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok int", func(t *testing.T) {
		var input string = "not an int"
		var want error = ErrCannotCast

		_, got := cast[int]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok int8", func(t *testing.T) {
		var input string = "123"
		var want int8 = 123

		got, err := cast[int8]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok int8", func(t *testing.T) {
		var input string = "not an int8"
		var want error = ErrCannotCast

		_, got := cast[int8]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok int16", func(t *testing.T) {
		var input string = "123"
		var want int16 = 123

		got, err := cast[int16]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok int16", func(t *testing.T) {
		var input string = "not an int16"
		var want error = ErrCannotCast

		_, got := cast[int16]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok int32", func(t *testing.T) {
		var input string = "123"
		var want int32 = 123

		got, err := cast[int32]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok int32", func(t *testing.T) {
		var input string = "not an int32"
		var want error = ErrCannotCast

		_, got := cast[int32]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok int64", func(t *testing.T) {
		var input string = "123"
		var want int64 = 123

		got, err := cast[int64]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok int64", func(t *testing.T) {
		var input string = "not an int64"
		var want error = ErrCannotCast

		_, got := cast[int64]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok uint", func(t *testing.T) {
		var input string = "123"
		var want uint = 123

		got, err := cast[uint]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok uint", func(t *testing.T) {
		var input string = "not a uint"
		var want error = ErrCannotCast

		_, got := cast[uint]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok uint8", func(t *testing.T) {
		var input string = "123"
		var want uint8 = 123

		got, err := cast[uint8]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok uint8", func(t *testing.T) {
		var input string = "not a uint8"
		var want error = ErrCannotCast

		_, got := cast[uint8]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok uint16", func(t *testing.T) {
		var input string = "123"
		var want uint16 = 123

		got, err := cast[uint16]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok uint16", func(t *testing.T) {
		var input string = "not a uint16"
		var want error = ErrCannotCast

		_, got := cast[uint16]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok uint32", func(t *testing.T) {
		var input string = "123"
		var want uint32 = 123

		got, err := cast[uint32]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok uint32", func(t *testing.T) {
		var input string = "not a uint32"
		var want error = ErrCannotCast

		_, got := cast[uint32]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok uint64", func(t *testing.T) {
		var input string = "123"
		var want uint64 = 123

		got, err := cast[uint64]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok uint64", func(t *testing.T) {
		var input string = "not a uint64"
		var want error = ErrCannotCast

		_, got := cast[uint64]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok float32", func(t *testing.T) {
		var input string = "12.3"
		var want float32 = 12.3

		got, err := cast[float32]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok float32", func(t *testing.T) {
		var input string = "not an float32"
		var want error = ErrCannotCast

		_, got := cast[float32]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("ok float64", func(t *testing.T) {
		var input string = "12.3"
		var want float64 = 12.3

		got, err := cast[float64]("", input)
		if err != nil {
			t.Fatalf("should not errorgot %v", err)
		}
		if want != got {
			t.Fatalf("want %vgot %v", want, got)
		}
	})

	t.Run("not ok float64", func(t *testing.T) {
		var input string = "not an float64"
		var want error = ErrCannotCast

		_, got := cast[float64]("", input)
		if !errors.Is(got, want) {
			t.Fatalf("want %vgot %v", want, got)
		}
	})
}
