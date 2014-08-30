package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rkrombho/getch"
)

func NewGetCommand(client *getch.Client) cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "retrieve a value for a given key",
		Action: func(c *cli.Context) {
			if len(c.Args()) == 0 {
				panic("get command requires a key to be queried")
			}
			key := c.Args()[0]
			value := client.Get(key)
			fmt.Print(value)
		},
	}
}
