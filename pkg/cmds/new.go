package cmds

import (
	"github.com/areller/dnmk/pkg/env"
)

type NewCommand struct {

}

func (nc *NewCommand) GetName() string {
	return "new"
}

func (nc *NewCommand) GetDescription() string {
	return "Creates new project from given template"
}

func (nc *NewCommand) Run(args []string, flags map[string]interface{}, envi *env.Env) {
	envi.IO.Print(env.Minimal, flags["name"].(string))
}

func NewNewCommand() *NewCommand {
	return &NewCommand{
		
	}
}