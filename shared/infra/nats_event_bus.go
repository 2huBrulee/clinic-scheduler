package infra

import (
	"clinic-scheduler/shared/domain"
	"github.com/nats-io/nats.go"
)

type NatsEventBus struct {
	jetStream nats.JetStream
}

func NewNatsEventBus(jetStream nats.JetStream) NatsEventBus {
	return NatsEventBus{
		jetStream: jetStream,
	}
}

func (nev NatsEventBus) Publish(event domain.DomainEvent) error {
	payload, err := event.ToJSON()

	if err != nil {
		return err
	}

	_, err = nev.jetStream.Publish(event.EventName(), payload)

	if err != nil {
		return err
	}

	return nil
}

func (nev NatsEventBus) Subscribe(eventName string, queue string, handler func(string)) error {
	_, err := nev.jetStream.QueueSubscribe(eventName, queue, func(msg *nats.Msg) {
		strMessage := string(msg.Data)
		handler(strMessage)
	})

	if err != nil {
		return err
	}

	return nil
}
