package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		input  []string
		result *ParseResult
	}{
		{
			input: []string{"../testdata/.env", "../testdata/.env2"},
			result: &ParseResult{
				Vars: map[string]string{
					`KEY`:                `value`,
					`BOOL_KEY`:           `true`,
					`INT_KEY`:            `16`,
					`FLOAT_KEY`:          `12.34`,
					`EXPANDED_KEY`:       `foo value`,
					`MULTI_EXPANDED_KEY`: `true 16 12.34`,
					`KEY2`:               `value2`,
				},
			},
		},
	}

	for _, tc := range testcases {
		p := NewParser(tc.input...)
		result, err := p.Parse()
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if result == nil {
			t.Fatalf("result is nil")
		}
		for wantKey, wantValue := range tc.result.Vars {
			gotValue, ok := result.Vars[wantKey]
			if !ok {
				t.Fatalf("missing key: %q", wantKey)
			}
			if wantValue != gotValue {
				t.Fatalf("wrong value: want %q, got %q", wantValue, gotValue)
			}
		}
	}
}

func TestLineParser(t *testing.T) {
	testcases := []struct {
		input     string
		wantKey   string
		wantValue string
	}{
		{`FOO=bar`, `FOO`, `bar`},
		{`EXPANDED=${FOO}`, `EXPANDED`, `${FOO}`},
		{`QUOTED="hey there"`, `QUOTED`, `hey there`},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			p := NewLineParser([]byte(tc.input))
			gotKey, gotValue, err := p.parse()
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			if gotKey != tc.wantKey {
				t.Fatalf("result key: want %s, got %s", tc.wantKey, gotKey)
			}
			if gotValue != tc.wantValue {
				t.Fatalf("result value: want %s, got %s", tc.wantValue, gotValue)
			}
		})
	}
}
