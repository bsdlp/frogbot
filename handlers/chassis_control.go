package handlers

import (
	"context"
	"fmt"
	"strings"

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

	return &api.InteractionResponseData{
		Content:         option.NewNullableString(formatPowerStatusResponse(resp)),
		AllowedMentions: &api.AllowedMentions{}, // don't mention anyone
	}
}

func formatPowerStatusResponse(resp *chassis.GetChassisStatusResponse) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "target: %s\n", resp.Target)
	fmt.Fprintf(&builder, "powered on: %t", resp.PoweredOn)
	return builder.String()
}
