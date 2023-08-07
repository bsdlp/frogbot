package handlers

import (
	"github.com/bsdlp/frogbot/clients"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

type Handler struct {
	*cmdroute.Router
	s       *state.State
	clients *clients.Clients
}

func New(s *state.State, clients *clients.Clients) *Handler {
	h := &Handler{
		s:       s,
		clients: clients,
	}

	h.Router = cmdroute.NewRouter()
	// Automatically defer handles if they're slow.
	h.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{}))
	h.AddFunc("power_on", h.powerOn)
	h.AddFunc("power_status", h.powerStatus)

	return h
}

func (h *Handler) OnCommand(e *gateway.InteractionCreateEvent) {
	switch data := e.Data.(type) {
	case *discord.CommandInteraction:
		switch data.Name {
		case "ipmi":
		}
	}
}
