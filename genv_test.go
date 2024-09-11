package genv

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func Test_Load_SingleFile(t *testing.T) {
	want := map[string]string{
		"KEY":      "value",
		"INT_KEY":  "16",
		"EXPANDED": "\"foo 16\"",
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

func Test_Get(t *testing.T) {
	if err := Load("./testdata/.env"); err != nil {
		t.Fatalf("%v\n", err)
	}

	str, err := Get[string]("KEY")
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if str != "value" {
		t.Fatalf("want %s, got %s", "value", str)
	}

	byt, err := Get[[]byte]("KEY")
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if !bytes.Equal([]byte("value"), byt) {
		t.Fatalf("want %v, got %v", []byte("value"), byt)
	}

	intv, err := Get[int]("INT_KEY")
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if intv != 16 {
		t.Fatalf("want %d, got %d", 16, intv)
	}

	boolv, err := Get[bool]("BOOL_KEY")
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if !boolv {
		t.Fatalf("want %t, got %t", true, boolv)
	}
}

func Test_GetStruct(t *testing.T) {
	type Config struct {
		StringKey string `genv:"KEY"`
		IntKey    int    `genv:"INT_KEY"`
		BoolKey   bool   `genv:"BOOL_KEY"`
	}

	var cfg Config
	err := GetStruct(&cfg)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	fmt.Printf("%#v\n", cfg)
}
