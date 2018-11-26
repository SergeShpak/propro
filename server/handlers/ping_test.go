package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	r := gin.New()
	r.GET("/ping", Ping())

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	expectedResponse := "pong"
	if resp.Body.String() != expectedResponse {
		t.Fatalf("expected response: %s, actual response: %s", expectedResponse, resp.Body.String())
	}
}
