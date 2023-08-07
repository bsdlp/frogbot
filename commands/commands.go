package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var Commands = []api.CreateCommandData{
	{
		Name:        "ping",
		Description: "ping pong!",
	},
	{
		Name:        "echo",
		Description: "echo back the argument",
		Options: []discord.CommandOption{
			&discord.StringOption{
				OptionName:  "argument",
				Description: "what's echoed back",
				Required:    true,
			},
		},
	},
	{
		Name:        "ipmi",
		Description: "ipmi controls",
		Options: []discord.CommandOption{
			discord.NewSubcommandGroupOption("chassis", "chassis commands",
				discord.NewSubcommandOption("power", "power commands",
					&discord.StringOption{
						OptionName:  "command",
						Description: "power command",
						Required:    true,
						Choices: []discord.StringChoice{
							{
								Name:  "status",
								Value: "status",
							},
							{
								Name:  "off",
								Value: "off",
							},
							{
								Name:  "on",
								Value: "on",
							},
							{
								Name:  "cycle",
								Value: "cycle",
							},
							{
								Name:  "reset",
								Value: "reset",
							},
							{
								Name:  "soft",
								Value: "soft",
							},
							{
								Name:  "diagnostic interrupt",
								Value: "diagnostic interrupt",
							},
						},
					},
				),
			),
		},
	},
}
