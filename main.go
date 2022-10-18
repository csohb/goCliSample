package main

import (
	"github.com/urfave/cli/v2"
	"goCliSample/broker"
	"log"
	"os"
)

func main() {
	var app = &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "broker",
				Aliases: []string{"t"},
				Usage:   "broker client pub or sub",
				Subcommands: []*cli.Command{
					{
						Name:  "pub",
						Usage: "publish message to topic",
						Action: func(context *cli.Context) error {
							topic := context.Args().First()
							msg := context.Args().Get(1)
							broker.Publish(topic, msg)
							return nil
						},
					}, {
						Name:  "sub",
						Usage: "subscribe topic",
						Action: func(context *cli.Context) error {
							topic := context.Args().First()
							broker.Subscribe(topic)
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
