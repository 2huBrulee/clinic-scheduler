package domain

type AggregateRoot struct {
	events DomainEvents
}

func InitAggregateRoot() AggregateRoot {
	return AggregateRoot{
		events: make(DomainEvents, 0),
	}
}

func (ar AggregateRoot) RecordEvent(event DomainEvent) {
	ar.events = append(ar.events, event)
}

func (ar AggregateRoot) FlushEvents() {
	ar.events = make(DomainEvents, 0)
}

func (ar AggregateRoot) PullEvents() DomainEvents {
	events := ar.events
	ar.FlushEvents()

	return events
}
