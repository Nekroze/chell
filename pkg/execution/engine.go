package execution

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Nekroze/chell/pkg/tplex"
)

var builtins = map[string]func([]string){
	"cd": func(args []string) {
		dest := os.Getenv("HOME")
		if len(args) != 0 {
			dest = args[0]
		}
		if e := os.Chdir(dest); e != nil {
			panic(e)
		}
	},

	"exit": func(_ []string) {
		os.Exit(0)
	},

	"help": func(_ []string) {
		fmt.Println("Usage: chell")
	},
}

func Execute(head string, tail ...string) error {
	if len(strings.TrimSpace(head)) == 0 {
		return nil
	}

	if f, ok := builtins[head]; ok {
		f(tail)
		return nil
	}

	cmd := exec.Command(head, tail...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Muxecute(head string, tail ...string) error {
	errChan := make(chan error)

	go func() {
		errChan <- Execute(head, tail...)
	}()

	select {
	case <-time.After(5 * time.Second):
		if e := tplex.Split(false); e != nil {
			os.Stderr.WriteString(e.Error())
		}
	case e := <-errChan:
		return e
	}

	err := <-errChan

	if err != nil {
		fmt.Println("command exit: ", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	os.Exit(0)
	return nil
}
