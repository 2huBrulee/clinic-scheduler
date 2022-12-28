package domain

type DomainEvent interface {
	EventName() string
	ToJSON() ([]byte, error)
}

type DomainEvents []DomainEvent

type BaseDomainEvent struct {
	EventID     string                 `json:"event_id"`
	AggregateID string                 `json:"aggregate_id"`
	Timestamp   int64                  `json:"timestamp"`
	Metadata    map[string]interface{} `json:"metadata"`
}
