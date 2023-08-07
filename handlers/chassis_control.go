package handlers

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (h *Handler) ipmi(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
	log.Println("test")
	return &api.InteractionResponseData{
		Content: option.NewNullableString(spew.Sdump(data.CommandInteractionOption)),
	}
}
