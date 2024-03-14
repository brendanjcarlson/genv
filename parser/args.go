package parser

import (
	"log"
	"os"

	"github.com/brendanjcarlson/genv/commands"
	"github.com/brendanjcarlson/genv/flags"
)

type Args struct {
	Command commands.Command
	Input   []string
	Output  string
	Package string
}

func ParseArgs(cliArgs []string) (args *Args) {
	args = &Args{
		Command: "",
		Input:   make([]string, 0),
		Output:  "",
		Package: "",
	}

	switch len(cliArgs) {
	case 1:
		// show help and exit
		args.Command = commands.HELP
		return
	case 2:
		ParseCmd(cliArgs, args)
	default:
		ParseCmd(cliArgs, args)
		ParseFlags(cliArgs, args)
	}

	return
}

func ParseCmd(cliArgs []string, args *Args) {
	cmd := commands.Command(cliArgs[1])
	cmd.ValidateOrExit()
	args.Command = cmd
}

func ParseFlags(cliArgs []string, args *Args) {
	var inpIdx, outIdx, pkgIdx int
	offset := 2

	for i := offset; i < len(cliArgs); i++ {
		switch flags.Flag(cliArgs[i]) {
		case flags.INPUT, flags.INPUT_SHORT:
			inpIdx = i
		case flags.OUTDIR, flags.OUTDIR_SHORT:
			outIdx = i
		case flags.PACKAGE, flags.PACKAGE_SHORT:
			pkgIdx = i
		}
	}

	if inpIdx == 0 || outIdx == 0 || pkgIdx == 0 {
		log.Fatalf("Usage: genv <command> -i <input> -o <output> -p <package>")
	}

	offset = 1

	args.Input = cliArgs[inpIdx+offset : outIdx]
	ValidateInputs(args)
	args.Output = cliArgs[outIdx+offset]
	args.Package = cliArgs[pkgIdx+offset]
}

func ValidateInputs(args *Args) {
	if len(args.Input) == 0 {
		log.Fatalf("No input files specified")
	}

	for _, v := range args.Input {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			log.Fatalf("Input file does not exist: %s", v)
		}
	}
}
