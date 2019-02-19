package tplex

import (
	"os"
	"os/exec"

	"github.com/Nekroze/chell/pkg/config"
)

func NewTmux() error {
	cmd := exec.Command("tmux", "new-session", exename, "-c")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Split(vertical bool) error {
	direction := "-h"
	if vertical {
		direction = "-v"
	}
	return tmux("split-window", direction, "-t", config.TmuxPaneID, exename, "-c")
}
