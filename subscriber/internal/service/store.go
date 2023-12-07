package service

import "github.com/nats-io/stan.go"

type InterfaceNats interface {
	Connect() (stan.Conn, string, error)
}

type Nats struct {
	ClusterId,
	ClientId,
	Channel string
}
