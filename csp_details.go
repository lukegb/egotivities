package eactivities

import "fmt"

// CSPInfo encapsulates information about a single CSP.
type CSPInfo struct {
	Code    string
	Name    string
	WebName string
	Acronym string
}

// ListCSPs gets the list of Clubs, Societies or Projects that this API key has permission to view.
func ListCSPs(c Client) ([]CSPInfo, error) {
	var cspInfo []CSPInfo
	if err := c.Get("/CSP", &cspInfo); err != nil {
		return nil, err
	}
	return cspInfo, nil
}

// CSPDetails returns basic details about a specified CSP.
func CSPDetails(c Client, centre string) (CSPInfo, error) {
	var cspInfo CSPInfo
	if err := c.Get(fmt.Sprintf("/CSP/%s", centre), &cspInfo); err != nil {
		return CSPInfo{}, err
	}
	return cspInfo, nil
}
