package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/toolset.sh/server"
	"github.com/nathan-osman/toolset.sh/templates"
	"github.com/urfave/cli/v2"

	_ "github.com/nathan-osman/toolset.sh/tools/ip"
	_ "github.com/nathan-osman/toolset.sh/tools/lorem"
	_ "github.com/nathan-osman/toolset.sh/tools/pi"
	_ "github.com/nathan-osman/toolset.sh/tools/rand"
	_ "github.com/nathan-osman/toolset.sh/tools/time"
	_ "github.com/nathan-osman/toolset.sh/tools/useragent"
	_ "github.com/nathan-osman/toolset.sh/tools/uuid"
)

func main() {
	app := &cli.App{
		Name:  "toolset.sh",
		Usage: "Free online utility functions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "server-addr",
				Value:   ":80",
				EnvVars: []string{"SERVER_ADDR"},
				Usage:   "HTTP address to listen on",
			},
			&cli.StringFlag{
				Name:    "server-name",
				Value:   "toolset.sh",
				EnvVars: []string{"SERVER_NAME"},
				Usage:   "domain name used for URLs",
			},
		},
		Action: func(c *cli.Context) error {

			// Set the server name (for the templates)
			templates.Name = c.String("server-name")

			// Create the server
			s, err := server.New(c.String("server-addr"))
			if err != nil {
				return err
			}
			defer s.Close()

			// Wait for SIGINT or SIGTERM
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
		os.Exit(1)
	}
}
