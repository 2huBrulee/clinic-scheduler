package http

import (
	"clinic-scheduler/module/doctor/application"
	"github.com/labstack/echo/v4"
)

func HandleCreateDoctor(cmdHandler application.CreateDoctorCommandHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		createdID, err := cmdHandler.Handle()

		if err != nil {
			return err
		}

		return c.JSON(201, map[string]interface{}{
			"id": createdID,
		})
	}
}
