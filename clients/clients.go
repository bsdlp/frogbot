package clients

import (
	"net/http"

	"github.com/bsdlp/chassiscontrol/rpc/chassis"
	"github.com/bsdlp/frogbot/configuration"
)

type Clients struct {
	ChassisControl chassis.ChassisControl
}

func New(cfg *configuration.Object) (*Clients, error) {
	c := &Clients{
		ChassisControl: chassis.NewChassisControlProtobufClient(cfg.ChassisControl.ServerAddress, http.DefaultClient),
	}
	return c, nil
}
