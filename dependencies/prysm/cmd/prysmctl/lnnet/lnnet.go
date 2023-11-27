package lnnet

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	{
		Name:  "lnnet",
		Usage: "commands for dealing with Ethereum beacon chain testnets",
		Subcommands: []*cli.Command{
			generateGenesisStateCmd,
		},
	},
}
