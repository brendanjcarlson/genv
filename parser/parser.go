package parser

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

type KeyValue struct {
	Key   string
	Value string
	Type  string
}

type Mode int

const (
	KEY Mode = iota
	VALUE
	TYPE
)

func ParseFiles(args *Args) (kvs map[string]*KeyValue) {
	kvs = make(map[string]*KeyValue)

	for _, file := range args.Input { // start file
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("Failed to open file: %s", file)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		k := new(bytes.Buffer)
		v := new(bytes.Buffer)
		t := new(bytes.Buffer)

		for scanner.Scan() { // start scan
			b := scanner.Bytes()

			if len(b) == 0 {
				continue
			}
			if b[0] == HASH {
				continue
			}

			mode := KEY
			doubles := 0
			singles := 0

			for i := 0; i < len(b); i++ { // start line
				var c byte
				var p byte
				c = b[i]
				if i > 0 {
					p = b[i-1]
				}

				switch mode { // start mode
				case KEY:
					if c == SPACE {
						continue
					} else if c == EQUAL {
						mode = VALUE
					} else {
						k.WriteByte(c)
					}

				case VALUE:
					if c == DOUBLE_QUOTE && p != BACKSLASH {
						doubles++
					}
					if c == SINGLE_QUOTE && p != BACKSLASH {
						singles++
					}
					if c == SPACE && !(singles%2 == 1 || doubles%2 == 1) {
						continue
					}
					if c == HASH && !(singles%2 == 1 || doubles%2 == 1) {
						mode = TYPE
						continue
					}

					v.WriteByte(c)

				case TYPE:
					if !(c == SPACE || c == HASH) {
						t.WriteByte(c)
					}

				} // end mode
			} // end line

			kvs[k.String()] = &KeyValue{
				Key:   k.String(),
				Value: v.String(),
				Type:  GetType(t.String()),
			}

			k.Reset()
			v.Reset()
			t.Reset()

		} // end scan
	} // end file

	return
}
