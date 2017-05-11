package cli

import (
	"github.com/psmiraglia/go-sandbox/sandbox/common"
	"github.com/urfave/cli"
	"os"
)

func Run() {
	app := cli.NewApp()
	app.Author = "Paolo Smiraglia"
	app.Description = "My Sandbox application"
	app.Email = "paolo.smiraglia@gmail.com"
	app.Usage = "My Sandbox application"
	app.Version = common.Version()

	app.Commands = []cli.Command {
		{
			Name:   "doit",
			Usage:  "Let's do it!",
			Action: func(c *cli.Context) {
				doit()
			},
		},
	}
	app.Run(os.Args)
}
