package main

import (
	"flag"
)

type HelpCommand struct {
	cmds map[string]Command
}

func (hc *HelpCommand) GetName() string {
	return "help"
}

func (hc *HelpCommand) GetDescription() string {
	return "Prints version and all other commands"
}

func (hc *HelpCommand) Run(io *IO, args []string, flags *flag.FlagSet) {
	io.Print(Minimal, "dnmk version 0.1")
	for name, cmd := range hc.cmds {
		io.Print(Minimal, name + " - " + cmd.GetDescription())
	}
}

func NewHelpCommand(commands map[string]Command) *HelpCommand {
	return &HelpCommand{
		cmds: commands,
	}
}