package hendels

import (
	"API/models"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

var nextID = 1
var messages = make(map[int]models.Message)

func GetHandler(c echo.Context) error {
	var msgSlice []models.Message

	for _, msg := range messages {
		msgSlice = append(msgSlice, msg)
	}

	sort.Slice(msgSlice, func(i, j int) bool {
		return msgSlice[i].ID < msgSlice[j].ID
	})

	return c.JSON(http.StatusOK, &msgSlice)
}

func PostHandler(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Can not added message",
		})
	}
	message.ID = nextID
	nextID++
	messages[message.ID] = message
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message have been added succsesfuly",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updatedMessage models.Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Can not update the message",
		})
	}
	if _, exits := messages[id]; !exits {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}
	updatedMessage.ID = id
	messages[id] = updatedMessage
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message was updated",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if _, exits := messages[id]; !exits {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}

	delete(messages, id)
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message was deleted",
	})
}
