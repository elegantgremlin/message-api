package server

import (
	"fmt"
	"messageApi/internal/service"
	"messageApi/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// addV1Routes adds the service middleware and routes to the RouterGroup
func addV1Routes(group *gin.RouterGroup, service service.Service) {
	group.Use(ServiceMiddleware(service))

	group.POST("/messages", CreateMessageHandler)
	group.GET("/messages", ListMessageHandler)
	group.GET("/messages/:id", GetMessageHandler)
	group.POST("/messages/:id", UpdateMessageHandler)
	group.DELETE("/messages/:id", DeleteMessageHandler)
}

// getService gets the service module from the request context
func getService(c *gin.Context) (service.Service, bool) {
	ctxService, ok := c.Get("service")
	if !ok {
		return nil, ok
	}

	return ctxService.(service.Service), true
}

// CreateMessageHandler handles requests to create messages
func CreateMessageHandler(c *gin.Context) {
	service, ok := getService(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorAsJSON("Service unavailable"))
		return
	}

	var msg types.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, errorAsJSON("Invalid body"))
		return
	}

	msg, err := service.CreateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorAsJSON(fmt.Sprintf("Error saving message: %v", err)))
		return
	}

	c.JSON(http.StatusCreated, msg)
}

// ListMessageHandler handles requests to list Messages
func ListMessageHandler(c *gin.Context) {
	service, ok := getService(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorAsJSON("Service unavailable"))
		return
	}

	msgs, err := service.ListMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorAsJSON(fmt.Sprintf("Error retrieving messages: %v", err)))
		return
	}

	c.JSON(http.StatusOK, msgs)
}

// GetMessageHandler handles requests to get a single Message
func GetMessageHandler(c *gin.Context) {
	service, ok := getService(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorAsJSON("Service unavailable"))
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}

	msg, err := service.GetMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorAsJSON(fmt.Sprintf("Error retrieving message: %v", err)))
		return
	}

	if msg.Id == 0 {
		c.JSON(http.StatusNotFound, errorAsJSON(fmt.Sprintf("Message not found for id %d", id)))
		return
	}

	c.JSON(http.StatusOK, msg)
}

// UpdateMessageHandler handles requests to update an existing Message
func UpdateMessageHandler(c *gin.Context) {
	service, ok := getService(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorAsJSON("Service unavailable"))
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorAsJSON("Invalid Id"))
		return
	}

	var msg types.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, errorAsJSON("Invalid body"))
		return
	}

	msg.Id = id

	msg, err = service.UpdateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorAsJSON(fmt.Sprintf("Error updating message: %v", err)))
		return
	}

	if msg.Id == 0 {
		c.JSON(http.StatusNotFound, errorAsJSON(fmt.Sprintf("Message not found for id %d", id)))
		return
	}

	c.JSON(http.StatusOK, msg)
}

// DeleteMessageHandler handles requests to delete an existing Message
func DeleteMessageHandler(c *gin.Context) {
	service, ok := getService(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorAsJSON("Service unavailable"))
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}

	if err := service.DeleteMessage(id); err != nil {
		c.JSON(http.StatusInternalServerError, errorAsJSON(fmt.Sprintf("Failed to delete message with id %d", id)))
		return
	}

	c.Status(http.StatusOK)
}
