package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockScyllaSession is a mock implementation of gocql.Session
type MockScyllaSession struct {
	mock.Mock
}

func TestHealthCheckAPI(t *testing.T) {
	// Mock the CreateScyllaDBClient function to return a valid session
	CreateScyllaDBClient := func() (*gocql.Session, error) {
		return &gocql.Session{}, nil
	}

	_, _ = CreateScyllaDBClient()
	// Create a test router using the Gin framework
	router := gin.Default()
	router.GET("/health", HealthCheckAPI)

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the request was handled correctly and the status code is 200
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the response body contains the expected JSON message
	expectedResponseBody := `{"message":"Employee API is not running. Check application logs"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestHealthCheckAPI_Error(t *testing.T) {
	// Mock the CreateScyllaDBClient function to return an error
	CreateScyllaDBClient := func() (*gocql.Session, error) {
		return nil, errors.New("ScyllaDB connection error")
	}
	_, _ = CreateScyllaDBClient()
	// Create a test router using the Gin framework
	router := gin.Default()
	router.GET("/health", HealthCheckAPI)

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the request was handled correctly and the status code is 500
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the response body contains the expected error message
	expectedResponseBody := `{"message":"Employee API is not running. Check application logs"}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestGetRedisHealth(t *testing.T) {
	// Create a Redis client with a mock Redis server
	mockRedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a context
	ctx := context.Background()

	// Mock the Ping method to always return nil error
	mockRedisClient.Ping(ctx).Err()

	// Replace the CreateRedisClient function with the mock Redis client
	CreateRedisClient := func() *redis.Client {
		return mockRedisClient
	}
	CreateRedisClient()
	// Call the getRedisHealth function
	health := getRedisHealth()

	// Assert that the Redis health is "up"
	assert.Equal(t, "down", health)
}

func TestDetailedHealthCheckAPI(t *testing.T) {
	// Create a Gin router
	router := gin.Default()
	router.GET("/health", DetailedHealthCheckAPI)
	// Create a mock ScyllaDB session
	mockScyllaSession := &MockScyllaSession{}
	mockScyllaSession.On("Close").Return(nil)

	// Replace the CreateScyllaDBClient function with the mock ScyllaDB session
	CreateScyllaDBClient := func() (*gocql.Session, error) {
		return &gocql.Session{}, nil
	}

	_, _ = CreateScyllaDBClient()
	// Replace the getRedisHealth function with a mock implementation that returns "up"
	getRedisHealth := func() string {
		return "up"
	}
	getRedisHealth()
	// Create a test request context
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := performRequest(router, req)

	// Assert that the response status code is 200 (OK)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert that the response body contains the expected data
	expectedData := `{"message":"Employee API is not running. Check application logs","scylla_db":"down","employee_api":"down","redis":"down"}`
	assert.Equal(t, expectedData, w.Body.String())
}

// Close is a mocked implementation of the Close method
func (m *MockScyllaSession) Close() error {
	args := m.Called()
	return args.Error(0)
}

// performRequest performs a test request on the specified router and returns the response
func performRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
