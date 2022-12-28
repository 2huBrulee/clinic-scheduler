package event_bus

import (
	"clinic-scheduler/module/doctor/application"
	"clinic-scheduler/module/doctor/domain/event"
	"github.com/labstack/gommon/log"
)

func HandleDisplayHelloWorldOnDoctorCreated(cmdHandler application.DisplayHelloWorldCommandHandler) func(message string) {
	return func(message string) {
		byteMessage := []byte(message)

		doctorCreatedEvent, err := event.DoctorCreatedEventFromJSON(byteMessage)

		if err != nil {
			log.Fatal(err)
		}

		err = cmdHandler.Handle(doctorCreatedEvent)

		if err != nil {
			log.Fatal(err)
		}
	}
}
