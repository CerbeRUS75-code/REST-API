package main

import (
	hendels "API/handels"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/messages", hendels.GetHandler)
	e.POST("/messages", hendels.PostHandler)
	e.PATCH("/messages/:id", hendels.PatchHandler)
	e.DELETE("/messages/:id", hendels.DeleteHandler)
	e.Start(":8080")
}
