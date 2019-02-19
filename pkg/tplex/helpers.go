package tplex

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Nekroze/chell/pkg/config"
)

var exename string

func init() {
	var err error
	exename, err = os.Executable()
	if err != nil {
		panic(err)
	}
}

func tmux(args ...string) error {
	if config.TerminalMultiplexingDisabled {
		fmt.Println(args)
		return nil
	}
	return exec.Command("tmux", args...).Run()
}
