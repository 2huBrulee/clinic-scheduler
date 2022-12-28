package main

import (
	"clinic-scheduler/module/doctor/application"
	"clinic-scheduler/module/doctor/ui/event_bus"
	"clinic-scheduler/module/doctor/ui/http"
	"clinic-scheduler/shared/definitions"
	"clinic-scheduler/shared/infra"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		panic(err)
	}

	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	_, err = js.StreamInfo("doctors")

	if err != nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     "doctors",
			Subjects: definitions.StringEventNames,
		})

		if err != nil {
			panic(err)
		}
	} else {
		_, err := js.UpdateStream(&nats.StreamConfig{
			Name:     "doctors",
			Subjects: definitions.StringEventNames,
		})

		if err != nil {
			panic(err)
		}
	}

	natsEventBus := infra.NewNatsEventBus(js)

	createDoctorCommandHandler := application.NewCreateDoctorCommandHandler(natsEventBus)

	createDoctorHandler := http.HandleCreateDoctor(createDoctorCommandHandler)

	displayHelloWorldCommandHandler := application.NewDisplayHelloWorldCommandHandler()

	displayHelloWorldHandler := event_bus.HandleDisplayHelloWorldOnDoctorCreated(displayHelloWorldCommandHandler)

	err = natsEventBus.Subscribe("doctor-created", "q1", displayHelloWorldHandler)

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.POST("/doctors", createDoctorHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
