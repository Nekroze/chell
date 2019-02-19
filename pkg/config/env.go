package config

import "os"

var TmuxPaneID string
var TerminalMultiplexingDisabled bool

func init() {
	TmuxPaneID = os.Getenv("TMUX_PANE")

	if os.Getenv("CHELL_MUXING") == "false" {
		TerminalMultiplexingDisabled = true
	}
}
