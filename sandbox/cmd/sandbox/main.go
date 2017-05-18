package main

import (
	"github.com/psmiraglia/go-sandbox/sandbox"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Authors = []cli.Author{
		cli.Author{"Paolo Smiraglia", "paolo.smiraglia@gmail.com"},
		cli.Author{"Alice", "alice@example.com"},
	}
	app.Description = "My Sandbox application"
	app.Usage = "My Sandbox application"
	app.Version = sandbox.Version()

	app.Commands = []cli.Command {
		{
			Name:   "doit",
			Usage:  "Let's do it!",
			Action: func(c *cli.Context) {
				sandbox.DoIt()
			},
		},
	}
	app.Run(os.Args)
}
