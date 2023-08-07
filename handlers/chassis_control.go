package handlers

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
)

func (h *Handler) ipmi(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
	spew.Dump(data.CommandInteractionOption)
	return &api.InteractionResponseData{}
}
