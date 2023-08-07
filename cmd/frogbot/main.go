package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/bsdlp/frogbot/commands"
	"github.com/bsdlp/frogbot/handlers"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	s := state.New("Bot " + token)

	// set up handlers
	h := handlers.New(s)
	s.AddInteractionHandler(h)
	s.AddIntents(gateway.IntentGuilds)
	s.AddHandler(func(*gateway.ReadyEvent) {
		me, _ := s.Me()
		log.Println("connected to the gateway as", me.Tag())
	})

	// initialize commands
	if err := overwriteCommands(s); err != nil {
		log.Fatalln("cannot update commands:", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// connect to discord gateway
	if err := s.Connect(ctx); err != nil {
		log.Fatalln("cannot connect:", err)
	}
}

func overwriteCommands(s *state.State) error {
	return cmdroute.OverwriteCommands(s, commands.Commands)
}
