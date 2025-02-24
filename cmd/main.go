package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Text string `json:"text"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages []Message

func GetHendler(c echo.Context) error {
	return c.JSON(http.StatusOK, &messages)
}

func PostHendler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Can not added message",
		})
	}
	messages = append(messages, message)
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message have been added succsesfuly",
	})
}

func main() {
	e := echo.New()
	e.GET("/messages", GetHendler)
	e.POST("/messages", PostHendler)
	e.Start(":8080")
}
