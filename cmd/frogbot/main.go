package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/bsdlp/frogbot/clients"
	"github.com/bsdlp/frogbot/commands"
	"github.com/bsdlp/frogbot/configuration"
	"github.com/bsdlp/frogbot/handlers"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

func main() {
	// read bot token
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	// get config file path env var
	cfgFilePath := os.Getenv("CONFIG_FILE_PATH")
	if cfgFilePath == "" {
		cfgFilePath = "/etc/frogbot/config"
	}

	// consume config file
	var cfg configuration.Object
	if configFh, err := os.Open(cfgFilePath); err == nil {
		err = json.NewDecoder(configFh).Decode(&cfg)
		if err != nil {
			log.Fatalf("error decoding config: %s", err.Error())
		}
		configFh.Close()
	} else {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	s := state.New("Bot " + token)

	clients, err := clients.New(&cfg)
	if err != nil {
		log.Fatalf("error instantiating clients: %s", err.Error())
	}

	// set up handlers
	h := handlers.New(s, clients)
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
