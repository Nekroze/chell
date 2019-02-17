package tplex

import (
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/Nekroze/chell/pkg/config"
)

var exename string

func init() {
	rand.Seed(time.Now().UnixNano())
	var err error
	exename, err = os.Executable()
	if err != nil {
		panic(err)
	}
}

func tmux(args ...string) error {
	if config.TerminalMultiplexingDisabled {
		return nil
	}
	return exec.Command("tmux", args...).Run()
}
