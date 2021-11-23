package main

import (
	"khidr/todo/api"
	"khidr/todo/persistence"
	"khidr/todo/service"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Config struct {
	ListenAddress string
	Timeout       int
	DSN           string
}

func main() {
	conf := Config{}
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "Todo App"
	app.Usage = "Create and View Todos"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Usage:       "`IP:PORT` for the service to listen on",
			Value:       ":8080",
			Destination: &conf.ListenAddress,
		},
		&cli.IntFlag{
			Name:        "timeout",
			Usage:       "Requests `TIMEOUT` in seconds",
			Value:       120,
			Destination: &conf.Timeout,
		},
		&cli.StringFlag{
			Name:        "dsn",
			Usage:       "Postgres dsn",
			Value:       "host=localhost user=khidr password=khidr dbname=onboarding port=5432 sslmode=disable",
			Destination: &conf.DSN,
		},
	}

	app.Action = func(context *cli.Context) error {
		timeout := time.Duration(conf.Timeout) * time.Second
		postgresConfig, err := persistence.Init(conf.DSN)
		if err != nil {
			log.Fatal(err)
		}
		todoSvc := service.New(postgresConfig)
		server := api.New(conf.ListenAddress, timeout, todoSvc)
		log.Println("Service A Running on Port " + conf.ListenAddress)
		return server.Start()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
