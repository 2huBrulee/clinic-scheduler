package entity

import (
	"clinic-scheduler/module/doctor/domain/event"
	"clinic-scheduler/shared/domain"
	"github.com/oklog/ulid/v2"
)

type Doctor struct {
	id   string
	name string
	domain.AggregateRoot
}

func NewDoctor(name string) Doctor {
	ar := domain.InitAggregateRoot()

	doctor := Doctor{
		id:   ulid.Make().String(),
		name: name,
	}

	doctor.AggregateRoot = ar

	doctorCreatedEvent := event.NewDoctorCreatedEvent(doctor.id)

	doctor.RecordEvent(doctorCreatedEvent)

	return doctor
}

func (d Doctor) ID() string {
	return d.id
}
