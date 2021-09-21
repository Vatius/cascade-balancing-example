package main

import (
	"bytes"
	"encoding/json"
	model "github.com/Vatius/cascade-balancing-example"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPostPayload(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handlerPostPayload))
	defer testServer.Close()
	testObject := []model.Payload{{1, 2, 3, 4, 5}}
	testBody, _ := json.Marshal(testObject)
	res, err := http.Post(testServer.URL, "application/json", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("bad response status code")
	}
}
