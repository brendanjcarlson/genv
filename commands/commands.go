package commands

import (
	"log"
	"os"
)

type Command string

const (
	HELP           Command = "help"
	HELP_SHORT     Command = "h"
	GENERATE       Command = "generate"
	GENERATE_SHORT Command = "g"
)

var validationMap = map[Command]bool{
	HELP:           true,
	HELP_SHORT:     true,
	GENERATE:       true,
	GENERATE_SHORT: true,
}

var usageMap = map[Command]string{
	HELP:           "Show this help message",
	HELP_SHORT:     "Show this help message",
	GENERATE:       "Generate code from the input file(s)/directory",
	GENERATE_SHORT: "Generate code from the input file(s)/directory",
}

func (c Command) String() string {
	return string(c)
}

func (c Command) ValidateOrExit() {
	if !validationMap[c] {
		log.Printf("Unknown command: %s\n", c.String())
		c.ShowUsage()
		os.Exit(1)
	}
}

func (c Command) ShowUsage() {
	log.Printf("Usage: genv %s\n", c.String())
	for k, v := range usageMap {
		log.Printf("%s: %s\n", k.String(), v)
	}
}
