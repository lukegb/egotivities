package eactivities

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestAPIError(t *testing.T) {
	const msg = "slartibartfast"
	err := APIError{
		HTTPCode: http.StatusForbidden,
		Message:  msg,
	}
	if got := err.Error(); got != msg {
		t.Errorf("err.Error() = %q, want %q", got, msg)
	}
}

func TestAPIErrorParse(t *testing.T) {
	var input = []byte(`{
	"message": "IP address xxx.xxx.xxx.xxx has been banned until 20\/06\/2015 18:05:00"
}`)
	e := APIError{}
	if err := json.Unmarshal(input, &e); err != nil {
		t.Errorf("json.Unmarshal(%q, APIError{}): %v", string(input), err)
		return
	}
	want := "IP address xxx.xxx.xxx.xxx has been banned until 20/06/2015 18:05:00"
	if got := e.Error(); got != want {
		t.Errorf("e.Error() = %q, want %q", got, want)
	}
}

func TestHTTPClientGet(t *testing.T) {
	client, backend := newTestClient(map[string]*http.Response{
		"/path": NewResponse(http.StatusOK, `{"hello_there": "why hello there"}`),
	})
	type output struct {
		HelloThere string `json:"hello_there"`
	}
	out := &output{}
	if err := client.Get("/path", out); err != nil {
		t.Errorf("client.Get: %v", err)
	}
	if out.HelloThere != "why hello there" {
		t.Errorf("out.HelloThere = %q; want %q", out.HelloThere, "why hello there")
	}
	if len(backend.Requests) != 1 {
		t.Errorf("backend.Requests = %d; want 1", len(backend.Requests))
	}
	if got := backend.Requests[0].Header.Get("X-API-Key"); got != testAPIKey {
		t.Errorf("backend.Requests[0].Header.Get(%q) = %q; want %q", "X-API-Key", got, testAPIKey)
	}
}

func TestHTTPClientGetAPIError(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/path": NewResponse(http.StatusForbidden, `{"message": "IP address xxx.xxx.xxx.xxx has been banned until 20/06/2015 19:00:00"}`),
	})
	out := make(map[string]interface{})
	err := client.Get("/path", out)
	if err == nil {
		t.Fatalf("client.Get: %v; want error", err)
	}
	if len(out) != 0 {
		t.Errorf("len(out) = %d; want 0", len(out))
	}
	want := "IP address xxx.xxx.xxx.xxx has been banned until 20/06/2015 19:00:00"
	if err.Error() != want {
		t.Errorf("err.Error() = %q; want %q", err.Error(), want)
	}
	eaerr, ok := err.(APIError)
	if !ok {
		t.Fatalf("err is not an APIError")
	}
	if eaerr.HTTPCode != http.StatusForbidden {
		t.Errorf("eaerr.HTTPCode = %d; want %d", eaerr.HTTPCode, http.StatusForbidden)
	}
}
