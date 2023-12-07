package service

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func New(cluster, client, channel string) InterfaceNats {
	return &Nats{
		ClusterId: cluster,
		ClientId:  client,
		Channel:   channel,
	}
}

func (n *Nats) Connect() (stan.Conn, string, error) {
	sc, err := stan.Connect(n.ClusterId, n.ClientId)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка при подключении %s", err)
	}

	return sc, n.Channel, nil
}
