package main

import (
	"github.com/urfave/cli"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "scheduler"
	app.Usage = "An amazing service"
	app.Action = func(c *cli.Context) error { return serve() }
	app.Commands = []cli.Command{
		{
			Name:   "migrate",
			Usage:  "Required when service initially running",
			Action: func(c *cli.Context) error { return doMigrate() },
		},
	}

}

func serve() (err error) {
	return
}

func doMigrate() (err error) {
	return
}
