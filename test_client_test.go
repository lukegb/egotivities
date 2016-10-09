package egotivities

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	testRoot   = "http://fake/root"
	testAPIKey = "fakeAPIKey"
)

type ClosingBuffer struct {
	*strings.Reader
}

func (b *ClosingBuffer) Close() error {
	return nil
}

func NewResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       &ClosingBuffer{strings.NewReader(body)},
	}
}

type fakeBackend struct {
	Error     error
	Responses map[string]*http.Response
	Requests  []*http.Request
}

func (c *fakeBackend) Do(req *http.Request) (*http.Response, error) {
	c.Requests = append(c.Requests, req)

	if c.Error != nil {
		return nil, c.Error
	}

	key := req.URL.String()
	if strings.HasPrefix(key, testRoot) {
		key = strings.TrimPrefix(key, testRoot)
	}
	resp := c.Responses[key]
	if resp == nil {
		return nil, fmt.Errorf("unexpected request: %s", key)
	}
	return resp, nil
}

func newTestClient(responses map[string]*http.Response) (Client, *fakeBackend) {
	backend := &fakeBackend{
		Responses: responses,
	}
	client := &httpClient{
		root:    testRoot,
		apiKey:  testAPIKey,
		backend: backend,
	}
	return client, backend
}
