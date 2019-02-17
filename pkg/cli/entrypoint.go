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

		cli.StringFlag{
			Name:        "instance, i",
			Value:       "",
			Usage:       "start as a child of the given instance",
			Destination: &config.InstanceID,
		},

		cli.IntFlag{
			Name:        "screen, s",
			Value:       0,
			Usage:       "screen ID of the child instance",
			Destination: &config.ScreenID,
		},
	}

	if e := app.Run(os.Args); e != nil {
		panic(e)
	}
}

func eMain(c *cli.Context) error {
	if config.InstanceID == "" && !config.TerminalMultiplexingDisabled {
		if e := tplex.NewTmux(); e != nil {
			panic(e)
		}
		if e := tplex.Attach(); e != nil {
			panic(e)
		}
	} else {
		repl.Run()
	}
	return nil
}
