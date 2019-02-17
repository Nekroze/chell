package config

import "os"

var TerminalMultiplexingDisabled bool

func init() {
	if os.Getenv("CHELL_MUXING") == "false" {
		TerminalMultiplexingDisabled = true
	}
}
