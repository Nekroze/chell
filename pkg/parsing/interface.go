package parsing

import (
	"os"

	"github.com/Nekroze/chell/pkg/execution"
)

func exec_command(command string, args []string) {
	err := execution.Muxecute(command, args...)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
