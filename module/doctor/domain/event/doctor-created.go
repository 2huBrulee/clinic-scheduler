package event

import (
	"clinic-scheduler/shared/definitions"
	"clinic-scheduler/shared/domain"
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"time"
)

type DoctorCreated struct {
	DoctorID  string                 `json:"doctor_id"`
	EventBase domain.BaseDomainEvent `json:"event_base"`
}

func NewDoctorCreatedEvent(doctorId string) DoctorCreated {
	return DoctorCreated{
		DoctorID: doctorId,
		EventBase: domain.BaseDomainEvent{
			EventID:     ulid.Make().String(),
			AggregateID: doctorId,
			Metadata:    map[string]interface{}{},
			Timestamp:   time.Now().UnixMilli(),
		},
	}
}

func DoctorCreatedEventFromJSON(raw []byte) (DoctorCreated, error) {
	var doctorCreatedEvent = DoctorCreated{}

	err := json.Unmarshal(raw, &doctorCreatedEvent)

	return doctorCreatedEvent, err

}

func (e DoctorCreated) EventName() string {
	return string(definitions.DoctorCreated)
}

func (e DoctorCreated) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
