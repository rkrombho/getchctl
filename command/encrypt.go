package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rkrombho/getch"
)

func NewEncryptCommand(client *getch.Client) cli.Command {
	return cli.Command{
		Name:  "encrypt",
		Usage: "encrypts a given value so that it can be used in Getch server-side configurations",
		Action: func(c *cli.Context) {
			if len(c.Args()) == 0 {
				panic("encrypt command requires a value to be encrypted")
			}
			value := c.Args()[0]
			encvalue := client.Encrypt(value)
			fmt.Println(encvalue)
		},
	}
}
