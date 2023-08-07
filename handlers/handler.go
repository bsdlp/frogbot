package handlers

import (
	"github.com/bsdlp/frogbot/clients"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
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
	h.AddFunc("ping", h.cmdPing)
	h.AddFunc("echo", h.cmdEcho)
	h.AddFunc("ipmi", h.ipmi)

	return h
}
