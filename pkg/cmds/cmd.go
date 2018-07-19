package cmds

import (
	"github.com/areller/dnmk/pkg/env"
)

type Command interface {
	GetName() string
	GetDescription() string
	Run(args []string, flags map[string]interface{}, envi *env.Env)
}