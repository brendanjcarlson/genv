package genv

import (
	"os"
	"testing"
)

func TestLoadSingleFile(t *testing.T) {
	want := map[string]string{
		"KEY":                "value",
		"BOOL_KEY":           "true",
		"INT_KEY":            "16",
		"FLOAT_KEY":          "12.34",
		"EXPANDED_KEY":       "foo value",
		"MULTI_EXPANDED_KEY": "true 16 12.34",
	}

	if err := Load("./testdata/.env"); err != nil {
		t.Fatalf("%v", err)
	}

	for k, v := range want {
		got := os.Getenv(k)
		if got != v {
			t.Fatalf("want %s=%s, got %s=%s", k, v, k, got)
		}
	}
}

func TestLoadMultipleFiles(t *testing.T) {
	if err := Load("./testdata/.env", "./testdata/.env2"); err != nil {
		t.Fatalf("%v", err)
	}

	_, ok := os.LookupEnv("KEY")
	if !ok {
		t.Fatalf("%s was not set", "KEY")
	}

	_, ok = os.LookupEnv("KEY2")
	if !ok {
		t.Fatalf("%s was not set", "KEY2")
	}
}

func TestLoadOrPanic(t *testing.T) {
	t.Run("does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should not have panicked")
			}
		}()
		LoadOrPanic("./testdata/.env")
	})

	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("should have panicked")
			}
		}()

		LoadOrPanic("not a real filepath")
	})
}
