package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ysergeyru/go-task-foxg/config"
	"github.com/ysergeyru/go-task-foxg/server"
)

func TestHandleUserLogDuplicatesCheck(t *testing.T) {
	// Read config
	config := config.Get()

	// Create new server instance
	s := server.New(config)

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Create handler
	handler := s.HTTPHandler()

	// Create a request
	req, err := http.NewRequest("GET", "/"+strconv.Itoa(rand.Intn(100))+"/"+strconv.Itoa(rand.Intn(100)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call handler's ServeHTTP method directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func BenchmarkHandleUserLogDuplicatesCheck(b *testing.B) {
	var t *testing.T
	for i := 0; i < b.N; i++ {
		TestHandleUserLogDuplicatesCheck(t)
	}
}
