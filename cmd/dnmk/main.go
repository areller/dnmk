package main

import (
	"fmt"
	"strings"
	"os"
	"github.com/areller/dnmk/pkg/cli"
)

func buildArgsAndFlags() ([]string, map[string]interface{}) {
	var args []string
    flags := make(map[string]interface{})

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") {
			withoutDash := arg[2:]
			eqIdx := strings.Index(withoutDash, "=")

			if eqIdx == -1 {
				flags[withoutDash] = true
			} else {
				key := withoutDash[0:eqIdx]
				eqIdx++
				value := withoutDash[eqIdx:]
				flags[key] = value
			}
		} else if strings.HasPrefix(arg, "-") {
			withoutDash := arg[1:]
			eqIdx := strings.Index(withoutDash, "=")

			if eqIdx == -1 {
				for _, flag := range withoutDash {
					flags[fmt.Sprintf("%c", flag)] = true
				}
			} else {
				keys := withoutDash[0:eqIdx]
				eqIdx++
				value := withoutDash[eqIdx:]

				for _, flag := range keys {
					flags[fmt.Sprintf("%c", flag)] = value
				}
			}
		} else {
			args = append(args, arg)
		}
	}

	return args, flags
}

func main() {
	app := cli.New()
	app.Run(buildArgsAndFlags())
}