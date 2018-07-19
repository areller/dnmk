package cmds

import (
	"github.com/areller/dnmk/pkg/env"
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

func (hc *HelpCommand) Run(args []string, flags map[string]interface{}, envi *env.Env) {
	envi.IO.Print(env.Minimal, "dnmk version 0.1")
	for name, cmd := range hc.cmds {
		envi.IO.Print(env.Minimal, name + " - " + cmd.GetDescription())
	}
}

func NewHelpCommand(commands map[string]Command) *HelpCommand {
	return &HelpCommand{
		cmds: commands,
	}
}