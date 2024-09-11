package genv

import (
	"os"
	"testing"
)

func Test_Load_SingleFile(t *testing.T) {
	want := map[string]string{
		"KEY":          "value",
		"BOOL_KEY":     "true",
		"INT_KEY":      "16",
		"FLOAT_KEY":    "12.34",
		"QUOTED_KEY":   "string",
		"EXPANDED_KEY": "foo 16",
	}

	if err := Load("./testdata/.env"); err != nil {
		t.Fatalf("%v\n", err)
	}

	for k, v := range want {
		got := os.Getenv(k)
		if got != v {
			t.Fatalf("want %s=%s, got %s=%s", k, v, k, got)
		}
	}
}

func Test_Load_MultipleFiles(t *testing.T) {
	if err := Load("./testdata/multi/one.env", "./testdata/multi/two.env"); err != nil {
		t.Fatalf("%v\n", err)
	}

	_, ok := os.LookupEnv("FROM_ONE")
	if !ok {
		t.Fatalf("%s was not set\n", "FROM_ONE")
	}

	_, ok = os.LookupEnv("FROM_TWO")
	if !ok {
		t.Fatalf("%s was not set\n", "FROM_TWO")
	}
}

func Test_LoadOrPanic(t *testing.T) {
	t.Run("does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("should not have panicked")
			}
		}()

		LoadOrPanic("./testdata/.env")
	})

	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("should have panicked")
			}
		}()

		LoadOrPanic("not a real filepath")
	})
}

func Test_expand(t *testing.T) {
	vars := map[string]*entry{
		"TEST_EXPAND_ONE": {"test_expand_one", "string"},
		"TEXT_EXPAND_TWO": {"${TEST_EXPAND_ONE} test_expand_two", "string"},
	}

	var wantValue string = "test_expand_one test_expand_two"

	err := expand(vars)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	got, ok := vars["TEXT_EXPAND_TWO"]
	if !ok {
		t.Errorf("map should have key %q", "TEST_EXPAND_TWO")
	}

	gotValue := got.value

	if wantValue != gotValue {
		t.Errorf("want %s\ngot %s", wantValue, gotValue)
	}
}
