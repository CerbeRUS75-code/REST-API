package hendels

import (
	"API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var messages []models.Message

func GetHendler(c echo.Context) error {
	return c.JSON(http.StatusOK, &messages)
}

func PostHendler(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Can not added message",
		})
	}
	messages = append(messages, message)
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message have been added succsesfuly",
	})
}
