package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/diazharizky/scheduler/pkg/httphandler"
	"github.com/diazharizky/scheduler/pkg/postgresql"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	app         *cli.App
	httpHandler *httphandler.HTTPHandler
	serverIP    string
	serverPort  string
)

const (
	appName = "scheduler"
)

func init() {
	app = cli.NewApp()
	app.Name = "scheduler"
	app.Usage = "Core service."
	app.Action = func(c *cli.Context) error { return serve() }
	app.Commands = []cli.Command{
		{
			Name:   "migrate-up",
			Usage:  "Required when service initially running.",
			Action: func(c *cli.Context) error { return migrateUp() },
		},
		{
			Name:   "migrate-down",
			Usage:  "Undo all actions on last migrate up.",
			Action: func(c *cli.Context) error { return migrateDown() },
		},
	}

	dbPort, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	httpHandler = &httphandler.HTTPHandler{
		Cron: cron.New(cron.WithSeconds()),
		DB: &postgresql.PGInstance{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     dbPort,
			Database: os.Getenv("POSTGRES_DATABASE"),
		},
	}
	httpHandler.DB.Open()
	httpHandler.Cron.Start()
}

func run() (err error) {
	err = app.Run(os.Args)

	return
}

func serve() error {
	errC := make(chan error)

	go func() {
		serverIP := os.Getenv("SERVER_IP")
		serverPort := os.Getenv("SERVER_PORT")

		log.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server worker started at pid " + strconv.Itoa(os.Getpid()) + " listening on " + net.JoinHostPort(serverIP, serverPort) + "\",\"service\":\"" + appName + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")

		httpHandler.Middlewares = append(httpHandler.Middlewares, getAllowAllCORS())
		sh, err := swaggerHandler()
		if err != nil {
			log.Fatal(err.Error())
		}

		h := httpHandler.Handler()
		h.Mount("/swagger.json", sh)
		http.ListenAndServe(net.JoinHostPort(serverIP, serverPort), h)
	}()

	return <-errC
}

func migrateUp() (err error) {
	err = httpHandler.DB.MigrateUp()

	return
}

func migrateDown() (err error) {
	err = httpHandler.DB.MigrateDown()

	return
}
