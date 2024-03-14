package main

import (
	"log"
	"os"

	"github.com/brendanjcarlson/genv/commands"
	"github.com/brendanjcarlson/genv/generator"
	"github.com/brendanjcarlson/genv/parser"
)

const (
	PROGRAM_NAME    string = "genv"
	PROGRAM_VERSION string = "0.0.1"
)

func main() {
	args := parser.ParseArgs(os.Args)

	switch args.Command {
	case commands.HELP:
		log.Println("help")
	case commands.GENERATE:
		generate(args)
	default:
		log.Fatalf("Unknown command: %s", args.Command)
	}
}

func generate(args *parser.Args) {
	kvs := parser.ParseFiles(args)

	err := generator.Generate(args, kvs)
	if err != nil {
		log.Fatalf("Error generating files: %s", err)
	}
}
