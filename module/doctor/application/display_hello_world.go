package application

import (
	"clinic-scheduler/module/doctor/domain/event"
	"github.com/labstack/gommon/log"
)

type DisplayHelloWorldCommandHandler struct {
}

func NewDisplayHelloWorldCommandHandler() DisplayHelloWorldCommandHandler {
	return DisplayHelloWorldCommandHandler{}
}

func (dhw DisplayHelloWorldCommandHandler) Handle(event event.DoctorCreated) error {
	log.Info("Hello, dr " + event.DoctorID)

	return nil
}
