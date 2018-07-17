package main

import (
	"flag"
)

type NewCommand struct {

}

func (nc *NewCommand) GetName() string {
	return "new"
}

func (nc *NewCommand) GetDescription() string {
	return "Creates new project from given template"
}

func (nc *NewCommand) Run(io *IO, args []string, flags *flag.FlagSet) {
	io.Print(Minimal, *flags.String("name", "noName", "sdf"))
}

func NewNewCommand() *NewCommand {
	return &NewCommand{
		
	}
}