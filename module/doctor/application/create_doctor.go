package application

import (
	"clinic-scheduler/module/doctor/domain/entity"
	"clinic-scheduler/module/doctor/domain/event"
	"clinic-scheduler/shared/domain"
)

type CreateDoctorCommand struct {
	Name string
}

func (c CreateDoctorCommand) GetCommandName() string {
	return "create-doctor"
}

type CreateDoctorCommandHandler struct {
	eventBus domain.EventBus
}

func NewCreateDoctorCommandHandler(eventBus domain.EventBus) CreateDoctorCommandHandler {
	return CreateDoctorCommandHandler{
		eventBus: eventBus,
	}
}

func (cd CreateDoctorCommandHandler) Handle() (string, error) {
	newDoctor := entity.NewDoctor("MELOCOTON")

	doctorCreatedEvent := event.NewDoctorCreatedEvent(newDoctor.ID())

	err := cd.eventBus.Publish(doctorCreatedEvent)

	if err != nil {
		return "", err
	}

	return newDoctor.ID(), nil
}
