package flags

import (
	"log"
	"os"
)

type Flag string

const (
	INPUT         Flag = "--input"
	INPUT_SHORT   Flag = "-i"
	OUTDIR        Flag = "--outdir"
	OUTDIR_SHORT  Flag = "-o"
	HELP          Flag = "--help"
	HELP_SHORT    Flag = "-h"
	PACKAGE       Flag = "--package"
	PACKAGE_SHORT Flag = "-p"
)

var validationMap = map[Flag]bool{
	INPUT:         true,
	INPUT_SHORT:   true,
	OUTDIR:        true,
	OUTDIR_SHORT:  true,
	HELP:          true,
	HELP_SHORT:    true,
	PACKAGE:       true,
	PACKAGE_SHORT: true,
}

var usageMap = map[Flag]string{
	INPUT:         "The input file to parse",
	INPUT_SHORT:   "The input file to parse",
	OUTDIR:        "The output directory for the generated files",
	OUTDIR_SHORT:  "The output directory for the generated files",
	HELP:          "Show this help message",
	HELP_SHORT:    "Show this help message",
	PACKAGE:       "The package name for the generated files",
	PACKAGE_SHORT: "The package name for the generated files",
}

func (f Flag) String() string {
	return string(f)
}

func (f Flag) ValidateOrExit() {
	if !validationMap[f] {
		log.Printf("Unknown flag: %s\n", f.String())
		f.ShowUsage()
		os.Exit(1)
	}
}

func (f Flag) ShowUsage() {
	log.Printf("Usage: genv <command> <flags>\n")
	log.Printf("Available flags are: \n")
	for k, v := range usageMap {
		log.Printf("%s: %s\n", k.String(), v)
	}
}
