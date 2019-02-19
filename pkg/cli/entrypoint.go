package cli

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/Nekroze/chell/pkg/config"
	"github.com/Nekroze/chell/pkg/repl"
	"github.com/Nekroze/chell/pkg/tplex"
)

func Entrypoint() {
	app := cli.NewApp()
	app.Name = "chell"
	app.Usage = "Now you're executing with portals!"
	app.Action = eMain
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "child, c",
			Usage:       "start as a child, splits tmux but does not start it",
			Destination: &config.Child,
		},
	}

	if e := app.Run(os.Args); e != nil {
		panic(e)
	}
}

func eMain(c *cli.Context) (err error) {
	if config.Child || config.TerminalMultiplexingDisabled {
		repl.Run()
	} else {
		err = tplex.NewTmux()
	}
	return err
}
