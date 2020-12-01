package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/diazharizky/scheduler/internal/utils"
	"github.com/diazharizky/scheduler/pkg/httphandler"
	"github.com/diazharizky/scheduler/pkg/postgresql"
	"github.com/diazharizky/scheduler/pkg/server"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli"
)

var app *cli.App
var httpHandler *httphandler.HTTPHandler
var serverIP string
var serverPort string

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

	httpHandler = &httphandler.HTTPHandler{
		Cron: cron.New(cron.WithSeconds()),
		DB: &postgresql.PGInstance{
			User:     server.Config.GetString("POSTGRES_USER"),
			Password: server.Config.GetString("POSTGRES_PASSWORD"),
			Host:     server.Config.GetString("POSTGRES_HOST"),
			Port:     server.Config.GetInt("POSTGRES_PORT"),
			Database: server.Config.GetString("POSTGRES_DATABASE"),
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
		serverIP := server.Config.GetString("SERVER_IP")
		serverPort := server.Config.GetString("SERVER_PORT")
		serverName := server.Config.GetString("SERVER_NAME")

		log.Println("{\"label\":\"server-http\",\"level\":\"info\",\"msg\":\"server worker started at pid " + strconv.Itoa(os.Getpid()) + " listening on " + net.JoinHostPort(serverIP, serverPort) + "\",\"service\":\"" + serverName + "\",\"time\":" + fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")

		http.ListenAndServe(net.JoinHostPort(serverIP, serverPort), httpHandler.Handler())
	}()

	return <-errC
}

func migrateUp() (err error) {
	dsn := utils.GetDSN("postgres", server.Config.GetString("POSTGRES_USER"), server.Config.GetString("POSTGRES_PASSWORD"), server.Config.GetString("POSTGRES_HOST"), server.Config.GetInt("POSTGRES_PORT"), server.Config.GetString("POSTGRES_DATABASE"), false)
	m, err := migrate.New("file://internal/migrations/postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := m.Up(); err != nil {
		log.Fatal(err.Error())
	}

	return
}

func migrateDown() (err error) {
	dsn := utils.GetDSN("postgres", server.Config.GetString("POSTGRES_USER"), server.Config.GetString("POSTGRES_PASSWORD"), server.Config.GetString("POSTGRES_HOST"), server.Config.GetInt("POSTGRES_PORT"), server.Config.GetString("POSTGRES_DATABASE"), false)
	m, err := migrate.New("file://internal/migrations/postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := m.Down(); err != nil {
		log.Fatal(err.Error())
	}

	return
}
