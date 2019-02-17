package parsing

import (
	"fmt"
	"os"

	"github.com/Nekroze/chell/pkg/execution"
)

func Parse(in string) {
	lex := newLexer([]byte(in))
	e := yyParse(lex)
	if e != 0 {
		os.Stderr.WriteString(fmt.Sprintf("Parser exit code: %d\n", e))
	}
}

func exec_command(command string, args []string) {
	err := execution.Muxecute(command, args...)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
