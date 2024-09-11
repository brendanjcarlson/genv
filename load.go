package genv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var ErrNotEnvExt = errors.New("file does not have \".env\" extension")

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

type entry struct {
	value string
	typ   string
}

func load(filenames ...string) error {
	vars := make(map[string]*entry)

	if len(filenames) == 0 {
		if err := loadFile(".env", vars); err != nil {
			return err
		}
	}

	for _, filename := range filenames {
		if err := loadFile(filename, vars); err != nil {
			return err
		}
	}

	if err := expand(vars); err != nil {
		return err
	}

	for k, v := range vars {
		os.Setenv(k, v.value)
	}

	return nil
}

func loadFile(filename string, vars map[string]*entry) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("genv: %w", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			if err := f.Close(); err != nil {
				return fmt.Errorf("genv: close: %w", err)
			}
			return fmt.Errorf("genv: parse: %w", err)
		}
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		} else if line[0] == '#' {
			continue
		}

		// TODO: handle type annotations after value e.g. INT_KEY=16 #int for code gen
		// TODO: handle escaped characters
		key, rest, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		value, typ, found := strings.Cut(rest, "#")
		if !found {
			typ = "string"
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if value[0] == '"' && value[len(value)-1] == '"' {
			value = strings.Trim(value, "\"")
		}
		typ = strings.TrimSpace(typ)

		vars[key] = &entry{value, typ}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("genv: close: %w", err)
	}
	return nil
}

var expandRegexp = regexp.MustCompile(`(\${(.*)})`)

func expand(vars map[string]*entry) error {
	for _, v := range vars {
		matches := expandRegexp.FindAllStringSubmatch(v.value, -1)
		if len(matches) == 0 {
			continue
		}
		for _, match := range matches {
			if len(match) != 3 {
				continue
			}
			if ev, ok := vars[match[2]]; ok {
				v.value = strings.Replace(v.value, match[1], ev.value, 1)
			} else {
				return fmt.Errorf("genv: undeclared environment variable: %q", match[2])
			}
		}
	}
	return nil
}
