package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rkrombho/getch"
)

func NewListCommand(client *getch.Client) cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "List all values stored in Getch for the host where this command is executed from",
		Action: func(c *cli.Context) {
			//in Getch a List is a a normal Get request with the keyword 'list'
			key := "list"
			value := client.Get(key)
			fmt.Println(value)
		},
	}
}
