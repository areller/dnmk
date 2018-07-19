package cli

import (
	"os"
	"fmt"
	"errors"
	"github.com/areller/dnmk/pkg/cmds"
	"github.com/areller/dnmk/pkg/env"
)

type CLI struct {
	cmds map[string]cmds.Command
}

func (cli *CLI) Run(args []string, flags map[string]interface{}) {
	verbose := false
	if v, ok := flags["v"]; ok {
		verbose = v.(bool)
	}

	quiet := false
	if q, ok := flags["q"]; ok {
		quiet = q.(bool)
	}

	var level int
	if verbose && quiet {
		fmt.Fprintln(os.Stderr, "Can't run verbosly and quietly at the same time")
		return
	} else if verbose && !quiet {
		level = env.Verbose
	} else if !verbose && quiet {
		level = env.Minimal
	} else {
		level = env.Normal
	}

	envi := env.New(level)

	if len(args) == 0 {
		cli.cmds["help"].Run(args, flags, envi)
	} else {
		cmd, ok := cli.cmds[args[0]]
		if !ok {
			envi.IO.Error(env.Minimal, errors.New("Invalid command '" + args[0] + "'"))
			return
		}

		cmd.Run(args, flags, envi)
	}
}

func addCommand(commands map[string]cmds.Command, cmd cmds.Command) {
	commands[cmd.GetName()] = cmd
}

func New() *CLI {
	commands := make(map[string]cmds.Command)
	addCommand(commands, cmds.NewNewCommand())
	addCommand(commands, cmds.NewHelpCommand(commands))

	return &CLI{
		cmds: commands,
	}
}