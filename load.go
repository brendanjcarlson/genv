package genv

import (
	"os"

	"github.com/brendanjcarlson/genv/parser"
)

// Load reads env files and loads the variables into the current process.
//
// Call this function as close to the start of your main function as possible.
//
// Calling load without arguments will attempt to load the .env file in the current path.
//
// Order matters.
//
// Variables set previously will be OVERRIDDEN if set in a subsequent file.
//
// See package 'autoload' to make your life even easier.
func Load(filenames ...string) error {
	if err := load(filenames...); err != nil {
		return err
	}
	return nil
}

// Calls Load and panics if there is an error.
func LoadOrPanic(filenames ...string) {
	if err := load(filenames...); err != nil {
		panic(err)
	}
}

func load(filenames ...string) error {
	p := parser.NewParser(filenames...)
	result, err := p.Parse()
	if err != nil {
		return err
	}
	for k, v := range result.Vars {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}
