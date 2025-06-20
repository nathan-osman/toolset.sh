package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/toolset.sh/manager"
	"github.com/nathan-osman/toolset.sh/server"
	"github.com/nathan-osman/toolset.sh/tools/ip"
	"github.com/nathan-osman/toolset.sh/tools/rand"
	"github.com/nathan-osman/toolset.sh/tools/time"
	"github.com/nathan-osman/toolset.sh/tools/uuid"
	"github.com/urfave/cli/v2"
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
		},
		Action: func(c *cli.Context) error {

			// Create the manager and register all of the tools
			m := manager.New()
			m.Register(ip.New())
			m.Register(rand.New())
			m.Register(time.New())
			m.Register(uuid.New())

			// Create the server
			s, err := server.New(c.String("server-addr"), m)
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
