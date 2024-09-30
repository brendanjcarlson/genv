package parser

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type ParseResult struct {
	mu   sync.Mutex
	Vars map[string]string
}

type Parser struct {
	filenames []string
	result    *ParseResult
}

func NewParser(filenames ...string) *Parser {
	return &Parser{
		filenames: filenames,
		result: &ParseResult{
			Vars: make(map[string]string),
		},
	}
}

func (p *Parser) Parse() (result *ParseResult, err error) {
	for _, file := range p.filenames {
		if err := p.parseFile(file); err != nil {
			return result, fmt.Errorf("parse: %s: %w", file, err)
		}
	}

	pattern := regexp.MustCompile(`\$\{([a-zA-Z0-9_]*)\}`)
	for k, v := range p.result.Vars {
		for {
			match := pattern.FindStringSubmatch(v)
			if len(match) != 2 {
				break
			}
			ev, ok := p.result.Vars[match[1]]
			if !ok {
				return nil, errors.New("unset expand")
			}
			v = strings.ReplaceAll(v, match[0], ev)
			p.result.mu.Lock()
			p.result.Vars[k] = v
			p.result.mu.Unlock()
		}
	}
	return p.result, err
}

func (p *Parser) parseFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		line = bytes.TrimSpace(line)
		if (len(line) > 0 && line[0] == '#') || len(line) == 0 {
			continue
		}
		lp := NewLineParser(line)
		key, value, err := lp.parse()
		if err != nil {
			return err
		}
		if key == "" || value == "" {
			continue
		}
		p.result.mu.Lock()
		p.result.Vars[key] = value
		p.result.mu.Unlock()
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Parser) expand() error {
	return nil
}

type LineParser struct {
	line     []byte
	position int
	cursor   int
	curr     byte
}

func NewLineParser(line []byte) *LineParser {
	p := &LineParser{line: line}
	p.consume()
	return p
}

func (p *LineParser) consume() {
	if p.cursor >= len(p.line) {
		p.curr = 0
	} else {
		p.curr = p.line[p.cursor]
	}
	p.position = p.cursor
	p.cursor++
}

func (p *LineParser) parse() (key, value string, err error) {
	p.skipWhitespace()

	if p.curr == '#' {
		return key, value, nil
	}

	p.skipWhitespace()

	key, err = p.consumeKey()
	if err != nil {
		return key, value, err
	}

	p.skipWhitespace()

	if p.curr == '=' {
		p.consume()
	} else {
		return key, value, fmt.Errorf("unexpected char: %s, expected %s", string(p.curr), "=")
	}

	value, err = p.consumeValue()
	if err != nil {
		return key, value, err
	}

	return key, value, nil
}

func (p *LineParser) skipWhitespace() {
	for p.curr == ' ' {
		p.consume()
	}
}

func (p *LineParser) consumeKey() (string, error) {
	start := p.position
	for {
		if p.curr == 0 {
			return "", errors.New("unexpected end of line")
		}
		if p.curr == '=' {
			break
		}
		p.consume()
	}
	return string(p.line[start:p.position]), nil
}

func (p *LineParser) consumeValue() (string, error) {
	start := p.position
	for {
		if p.curr == 0 {
			break
		}
		p.consume()
	}

	return strings.TrimSpace(strings.Trim(string(p.line[start:p.position]), "\"")), nil
}
