package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
)

// TODO: add autocomplete options for server ids
var Commands = []api.CreateCommandData{
	{
		Name:        "power_on",
		Description: "power on the server",
	},
	{
		Name:        "power_status",
		Description: "get power status",
	},
}
