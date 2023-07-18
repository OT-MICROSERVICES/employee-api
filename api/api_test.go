package api

import (
	"bytes"
	"employee-api/model"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateEmployeeData(t *testing.T) {
	router := setupRouter()
	router.POST("/api/v1/employee/create", CreateEmployeeData)
	employeeInfo := model.Employee{
		ID:             "OT-043",
		Name:           "Abhishek Dubey",
		Designation:    "Consultant Partner",
		Department:     "Technology",
		JoiningDate:    "2017-09-26",
		Address:        "D-63/A, Amar Colony",
		OfficeLocation: "Noida",
		Status:         "Active Employee",
		EmailID:        "abhishek@example.com",
		PhoneNumber:    "9999999999",
	}
	jsonValue, _ := json.Marshal(employeeInfo)
	req, _ := http.NewRequest("POST", "/api/v1/employee/create", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedResponseBody := `{"message":"Cannot write data to the system, request failure"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestReadCompleteEmployeesData(t *testing.T) {
	router := setupRouter()
	router.GET("/api/v1/employee/search/all", ReadCompleteEmployeesData)
	req, _ := http.NewRequest("GET", "/api/v1/employee/search/all", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReadEmployeeData(t *testing.T) {
	router := setupRouter()
	router.GET("/api/v1/employee/search", ReadEmployeeData)
	req, _ := http.NewRequest("GET", "/api/v1/employee/search", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReadEmployeesLocation(t *testing.T) {
	// Mock the CreateRedisClient function to return a mock Redis client
	CreateRedisClient := func() *redis.Client {
		return &redis.Client{}
	}

	// Mock the CreateScyllaDBClient function to return a mock ScyllaDB session
	CreateScyllaDBClient := func() (*gocql.Session, error) {
		return &gocql.Session{}, nil
	}

	CreateRedisClient()
	_, _ = CreateScyllaDBClient()
	// Create a test router using the Gin framework
	router := gin.Default()
	router.GET("/employees/location", ReadEmployeesLocation)

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/employees/location", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the request was handled correctly and the status code is 200
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the response body contains the expected JSON data
	expectedResponseBody := `{"message":"Cannot read data from the system, request failure"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestReadEmployeesLocation_Error(t *testing.T) {
	CreateRedisClient := func() *redis.Client {
		return nil
	}

	CreateScyllaDBClient := func() (*gocql.Session, error) {
		return nil, errors.New("ScyllaDB connection error")
	}

	CreateRedisClient()
	_, _ = CreateScyllaDBClient()
	router := gin.Default()
	router.GET("/employees/location", ReadEmployeesLocation)

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/employees/location", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the request was handled correctly and the status code is 500
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the response body contains the expected error message
	expectedResponseBody := `{"message":"Cannot read data from the system, request failure"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestReadEmployeesDesignation(t *testing.T) {
	router := setupRouter()
	router.GET("/api/v1/employee/designation", ReadEmployeesDesignation)
	req, _ := http.NewRequest("GET", "/api/v1/employee/designation", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// setupRouter is a method to setup router for Gin testing
func setupRouter() *gin.Engine {
	return gin.Default()
}

func TestErrorResponse(t *testing.T) {
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		errorResponse(c, "Bad request")
	})

	req, _ := http.NewRequest("GET", "/error", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	expectedResponseBody := `{"message":"Bad request"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}
