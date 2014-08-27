package command

import (
	"github.com/codegangsta/cli"
	"github.com/rkrombho/getch"
)

func NewGetFileCommand(client *getch.Client) cli.Command {
	return cli.Command{
		Name:  "getfile",
		Usage: "Downloads the file with the given name from Getch",
		Flags: []cli.Flag{
			cli.StringFlag{"outputfile, o", "", "write the output file with the given name", ""},
		},
		Action: func(c *cli.Context) {
			file := c.Args()[0]
			outputname := c.String("outputfile")
			client.GetFile(file, outputname)
		},
	}
}
