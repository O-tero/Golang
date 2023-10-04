package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create new app
	app := cli.NewApp()

	// add flags with three arguements
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "name",
			Value: "stranger",
			Usage: "your wonderful name",
		},
		cli.IntFlag{
			Name: "age",
			Value: 0,
			Usage: "your graceful age",
		},
	}
	// This functions parses and brings data in cli.context struct
	app.Action = func(c *cli.Context) error {
		// c.String, c.Int looks for value of given flag
		log.Printf("Hello %s (%d years), Welcome to the command line world",
		c.String("name"), c.Int("age"))
		return nil
	}
	// Pass os.Args to the cli app to parse content
	app.Run(os.Args)
}
