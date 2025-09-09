package routes

import (
	"net/http"
	"strconv"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event created", "event": event})

}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})

}
