package handlers

import (
	"context"
	"fmt"

	"github.com/bsdlp/chassiscontrol/rpc/chassis"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (h *Handler) powerOn(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
	resp, err := h.clients.ChassisControl.IssueChassisControlCommand(ctx, &chassis.ChassisControlRequest{
		Target:                "1RYJN22",
		ChassisControlCommand: chassis.ChassisControlCommand_ON,
	})
	if err != nil {
		return errorResponse(err)
	}
	if resp.Error != "" {
		return errorResponse(err)
	}

	return &api.InteractionResponseData{
		Content:         option.NewNullableString(fmt.Sprintf("power on command sent to target %s", resp.Target)),
		AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
	}
}

func (h *Handler) powerStatus(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
	resp, err := h.clients.ChassisControl.GetChassisStatus(ctx, &chassis.GetChassisStatusRequest{
		Target: "1RYJN22",
	})
	if err != nil {
		return errorResponse(err)
	}

	content := fmt.Sprintf("%s: :green_circle:", resp.Target)
	return &api.InteractionResponseData{
		Content:         option.NewNullableString(content),
		AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
	}
}
