package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(LoggingMiddleware())

	req, _ := http.NewRequest("GET", "/example", nil)
	req.Header.Set("X-Real-IP", "192.168.0.1")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

	expectedLogMessage := "404 page not found"
	assert.Contains(t, w.Body.String(), expectedLogMessage)
	assert.Contains(t, req.Method, "GET")
	assert.Contains(t, req.URL.Path, "/example")
	assert.Contains(t, "192.168.0.1", "192.168.0.1")
}
