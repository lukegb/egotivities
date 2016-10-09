package egotivities

import (
	"encoding/json"
	"net/http"
)

const defaultRoot = "https://eactivities.union.ic.ac.uk/API"

// APIError is an error returned by the eActivities API itself.
type APIError struct {
	HTTPCode int
	Message  string `json:"message"`
}

// Error returns the message contained in the error.
func (err APIError) Error() string {
	return err.Message
}

// Client wraps access to the eActivities API.
type Client interface {
	Get(path string, output interface{}) error
}

// backend is the backend used to retrieve data from eActivities.
type backend interface {
	Do(req *http.Request) (*http.Response, error)
}

// httpClient is an implementation of Client which uses the HTTP interface.
type httpClient struct {
	root    string
	apiKey  string
	backend backend
}

// Ensure that httpClient implements Client.
var _ Client = new(httpClient)

// Get retrieves a given path and deserializes the result into the provided interface{} using json.Unmarshal.
func (c *httpClient) Get(path string, output interface{}) error {
	req, err := http.NewRequest("GET", c.root+path, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", c.apiKey)
	resp, err := c.backend.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err := APIError{}
		d.Decode(&err)
		err.HTTPCode = resp.StatusCode
		return err
	}

	return d.Decode(output)
}

// NewClient creates a new instance of Client using the default implementation.
func NewClient(apiKey string) Client {
	return &httpClient{
		root:    defaultRoot,
		apiKey:  apiKey,
		backend: http.DefaultClient,
	}
}
