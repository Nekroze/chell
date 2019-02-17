package tplex

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"github.com/Nekroze/chell/pkg/config"
)

func NewTmux() error {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	config.InstanceID = string(b)

	return tmux(
		"new-session", "-d", "-s", config.InstanceID, exename, "-i", config.InstanceID, "-s", "1",
	)
}

var nextID int

func getSplitID() int {
	if config.ScreenID > nextID {
		nextID = config.ScreenID
	}
	nextID += 1
	return nextID
}

func Split(vertical bool) error {
	direction := "-h"
	if vertical {
		direction = "-v"
	}
	target := fmt.Sprintf("%s:1.%d", config.InstanceID, config.ScreenID)
	newSID := fmt.Sprintf("%d", getSplitID())
	fmt.Println("splitting")
	return tmux(
		"split-window", direction, "-t", target, exename, "-i", target, "-s", newSID,
	)
}

func Attach() error {
	cmd := exec.Command("tmux", "attach", "-t", config.InstanceID)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
