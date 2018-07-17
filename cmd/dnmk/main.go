package main

import (
	"fmt"
	"flag"
	"os"
)

type Command interface {
	GetName() string
	GetDescription() string
	Run(io *IO, args []string, flags *flag.FlagSet)
}

func addCommand(commands map[string]Command, cmd Command) {
	commands[cmd.GetName()] = cmd
}

func main() {
	cmds := make(map[string]Command)

	addCommand(cmds, NewHelpCommand(cmds))
	addCommand(cmds, NewNewCommand())

	verbose := flag.Bool("v", false, "Verbose output")
	quiet := flag.Bool("q", false, "Quiet output")
	flag.Parse()

	args := flag.Args()

	var level int
	if *verbose && *quiet {
		fmt.Fprintln(os.Stderr, "Can't run verbosly and quietly at the same time")
		return
	} else if *verbose && !*quiet {
		level = Verbose
	} else if !*verbose && *quiet {
		level = Minimal
	} else {
		level = Normal
	}

	if len(args) == 0 {
		cmds["help"].Run(NewIO(level), []string{}, flag.CommandLine)
	} else {
		cmd, ok := cmds[args[0]]
		if !ok {
			fmt.Fprintln(os.Stderr, "Invalid command '" + args[0] + "'")
			return
		}

		cmd.Run(NewIO(level), args[1:], flag.CommandLine)
	}
}