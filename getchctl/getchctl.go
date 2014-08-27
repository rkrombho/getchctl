package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rkrombho/getch"
	"github.com/rkrombho/command"
	"os"
)

const (
	GETCH_SERVER_ENV_VAR = "GETCH_SERVER"
)

func main() {
	//prepare the Getch client
	url := os.Getenv(GETCH_SERVER_ENV_VAR)
	if len(url) == 0 {
		panic(fmt.Sprintf("Environment variable %v not set", GETCH_SERVER_ENV_VAR))
	}
	client := getch.NewClient(url)
	//define the CLI interface
	app := cli.NewApp()
	app.Name = "getchctl"
	app.Usage = "A command line client for Getch"
	app.Version = VERSION
	app.Commands = []cli.Command{
		command.NewGetCommand(client),
		command.NewListCommand(client),
		command.NewGetFileCommand(client),
		command.NewEncryptCommand(client),
	}
	app.Run(os.Args)
}
