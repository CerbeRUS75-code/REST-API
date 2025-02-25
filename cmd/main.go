package main

import (
	"API/hendels"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/messages", hendels.GetHendler)
	e.POST("/messages", hendels.PostHendler)
	e.Start(":8080")
}
