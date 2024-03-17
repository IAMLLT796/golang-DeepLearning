package main

import (
	"io"
	"net/http"
	"testing"
)

func Test_client(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/ping", "", nil)
	if err != nil {
		panic(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Logf("%s", body)
	t.Error("test")
}
