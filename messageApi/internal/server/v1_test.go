package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"messageApi/internal/service"
	"messageApi/internal/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupPostRouter creates a router and adds a single POST endpoint to it
func setupPostRouter(service service.Service, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(ServiceMiddleware(service))
	router.POST("/", handler)

	return router
}

// setupPostRouter creates a router and adds a single POST endpoint to it
func setupPostRouterWithId(service service.Service, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(ServiceMiddleware(service))
	router.POST("/:id", handler)

	return router
}

// setupGetRouter creates a router and adds a single GET endpoint to it
func setupGetRouter(service service.Service, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(ServiceMiddleware(service))
	router.GET("/", handler)

	return router
}

// setupGetRouterWithId creates a router and adds a single GET endpoint with id param to it
func setupGetRouterWithId(service service.Service, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(ServiceMiddleware(service))
	router.GET("/:id", handler)

	return router
}

// setupDeleteRouter creates a router and adds a single DELETE endpoint to it
func setupDeleteRouterWithId(service service.Service, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(ServiceMiddleware(service))
	router.DELETE("/:id", handler)

	return router
}

// TestCreateMessage tests successfully creating a message
func TestCreateMessage(t *testing.T) {
	responseMessage := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	mockResponse, _ := json.Marshal(responseMessage)
	service_stub := service.ServiceStub{CreateMessageResponse: responseMessage}
	w := httptest.NewRecorder()
	router := setupPostRouter(&service_stub, CreateMessageHandler)

	jsonBody := []byte(`{"message": "racecar"}`)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestCreateMessageServiceUnavailable tests that an error is returned when the service is not initialized
func TestCreateMessageServiceUnavailable(t *testing.T) {
	mockResponse := `{"error":"Service unavailable"}`
	router := gin.Default()
	w := httptest.NewRecorder()

	router.POST("/", CreateMessageHandler)

	jsonBody := []byte(`{"message": "test}`)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer((jsonBody)))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestCreateMessageInvalidBody tests an error is returned when Message is not a string
func TestCreateMessageInvalidBody(t *testing.T) {
	mockResponse := `{"error":"Invalid body"}`
	service_stub := service.ServiceStub{}
	w := httptest.NewRecorder()
	router := setupPostRouter(&service_stub, CreateMessageHandler)

	jsonBody := []byte(`{"message": 1234}`)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestCreateMessageError tests an error is returned when the call to the service fails
func TestCreateMessageError(t *testing.T) {
	errorMsg := "create message failed"
	mockResponse := fmt.Sprintf(`{"error":"Error saving message: %s"}`, errorMsg)
	service_stub := service.ServiceStub{CreateMessageError: errors.New(errorMsg)}
	w := httptest.NewRecorder()
	router := setupPostRouter(&service_stub, CreateMessageHandler)

	jsonBody := []byte(`{"message": "racecar"}`)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestListMessage tests successfully retrieving the Message list
func TestListMessage(t *testing.T) {
	responseMessages := []types.Message{{Id: 1, Message: "racecar", IsPalindrome: true}, {Id: 2, Message: "second", IsPalindrome: false}}
	mockResponse, _ := json.Marshal(responseMessages)
	service_stub := service.ServiceStub{ListMessagesResponse: responseMessages}
	w := httptest.NewRecorder()
	router := setupGetRouter(&service_stub, ListMessageHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestListMessageServiceUnavailable tests that an error is returned when the service is not initialized
func TestListMessageServiceUnavailable(t *testing.T) {
	mockResponse := `{"error":"Service unavailable"}`
	router := gin.Default()
	w := httptest.NewRecorder()

	router.GET("/", ListMessageHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestListMessageError tests an error is returned when the call to the service module fails
func TestListMessageError(t *testing.T) {
	errorMsg := "list message failed"
	mockResponse := fmt.Sprintf(`{"error":"Error retrieving messages: %s"}`, errorMsg)
	service_stub := service.ServiceStub{ListMessagesError: errors.New(errorMsg)}
	w := httptest.NewRecorder()
	router := setupGetRouter(&service_stub, ListMessageHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestGetMessage tests successfully retrieving a single Message
func TestGetMessage(t *testing.T) {
	responseMessage := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	mockResponse, _ := json.Marshal(responseMessage)
	service_stub := service.ServiceStub{GetMessageResponse: responseMessage}
	w := httptest.NewRecorder()
	router := setupGetRouterWithId(&service_stub, GetMessageHandler)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/%d", responseMessage.Id), nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetMessageServiceUnavailable tests that an error is returned when the service is not initialized
func TestGetMessageServiceUnavailable(t *testing.T) {
	mockResponse := `{"error":"Service unavailable"}`
	router := gin.Default()
	w := httptest.NewRecorder()

	router.GET("/:id", GetMessageHandler)

	req, _ := http.NewRequest("GET", "/1", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestGetMessageError tests an error is returned when the call to the service module fails
func TestGetMessageError(t *testing.T) {
	errorMsg := "get message failed"
	mockResponse := fmt.Sprintf(`{"error":"Error retrieving message: %s"}`, errorMsg)
	service_stub := service.ServiceStub{GetMessageError: errors.New(errorMsg)}
	w := httptest.NewRecorder()
	router := setupGetRouterWithId(&service_stub, GetMessageHandler)

	req, _ := http.NewRequest("GET", "/1", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestUpdateMessage tests successfully updating a message
func TestUpdateMessage(t *testing.T) {
	responseMessage := types.Message{Id: 1, Message: "racecar", IsPalindrome: true}
	mockResponse, _ := json.Marshal(responseMessage)
	service_stub := service.ServiceStub{UpdateMessageResponse: responseMessage}
	w := httptest.NewRecorder()
	router := setupPostRouterWithId(&service_stub, UpdateMessageHandler)

	jsonBody := []byte(`{"message": "racecar"}`)

	req, _ := http.NewRequest("POST", fmt.Sprintf("/%d", responseMessage.Id), bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestUpdateMessageServiceUnavailable tests that an error is returned when the service is not initialized
func TestUpdateMessageServiceUnavailable(t *testing.T) {
	mockResponse := `{"error":"Service unavailable"}`
	router := gin.Default()
	w := httptest.NewRecorder()

	router.POST("/:id", UpdateMessageHandler)

	jsonBody := []byte(`{"message": "test}`)

	req, _ := http.NewRequest("POST", "/1", bytes.NewBuffer((jsonBody)))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestUpdateMessageInvalidBody tests an error is returned when Message is not a string
func TestUpdateMessageInvalidBody(t *testing.T) {
	mockResponse := `{"error":"Invalid body"}`
	service_stub := service.ServiceStub{}
	w := httptest.NewRecorder()
	router := setupPostRouterWithId(&service_stub, UpdateMessageHandler)

	jsonBody := []byte(`{"message": 1234}`)

	req, _ := http.NewRequest("POST", "/1", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestUpdateMessageError tests an error is returned when the call to the service fails
func TestUpdateMessageError(t *testing.T) {
	errorMsg := "update message failed"
	mockResponse := fmt.Sprintf(`{"error":"Error updating message: %s"}`, errorMsg)
	service_stub := service.ServiceStub{UpdateMessageError: errors.New(errorMsg)}
	w := httptest.NewRecorder()
	router := setupPostRouterWithId(&service_stub, UpdateMessageHandler)

	jsonBody := []byte(`{"message": "racecar"}`)

	req, _ := http.NewRequest("POST", "/1", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestDeleteMessage tests successfully updating a message
func TestDeleteMessage(t *testing.T) {
	id := 1
	service_stub := service.ServiceStub{}
	w := httptest.NewRecorder()
	router := setupDeleteRouterWithId(&service_stub, DeleteMessageHandler)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/%d", id), nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestDeleteMessageServiceUnavailable tests that an error is returned when the service is not initialized
func TestDeleteMessageServiceUnavailable(t *testing.T) {
	mockResponse := `{"error":"Service unavailable"}`
	router := gin.Default()
	w := httptest.NewRecorder()

	router.DELETE("/:id", DeleteMessageHandler)

	req, _ := http.NewRequest("DELETE", "/1", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// TestDeleteMessageError tests an error is returned when the call to the service fails
func TestDeleteMessageError(t *testing.T) {
	mockResponse := `{"error":"Failed to delete message with id 1"}`
	service_stub := service.ServiceStub{DeleteMessageError: errors.New("")}
	w := httptest.NewRecorder()
	router := setupDeleteRouterWithId(&service_stub, DeleteMessageHandler)

	req, _ := http.NewRequest("DELETE", "/1", nil)

	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
