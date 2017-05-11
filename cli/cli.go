package cli

import (
	"github.com/psmiraglia/sandbox/common"
	"github.com/urfave/cli"
	"os"
)

func Run() {
	app := cli.NewApp()
	app.Author = "Paolo Smiraglia"
	app.Description = "My Sandbox application"
	app.Email = "paolo.smiraglia@gmail.com"
	app.Name = common.Name
	app.Usage = "My Sandbox application"
	app.Version = common.Version()
	app.Run(os.Args)
}
