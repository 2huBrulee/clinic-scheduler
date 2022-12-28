package domain

type EventBus interface {
	Publish(event DomainEvent) error
	Subscribe(eventName string, queue string, handler func(strMessage string)) error
}
