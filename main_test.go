package main

import (
	"net/http"
	"net/http/httptest"
	"takeHome/filehandler"
	"testing"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(filehandler.UploadFileHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
