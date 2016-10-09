package eactivities

import (
	"net/http"
	"testing"

	"github.com/d4l3k/messagediff"
)

const FerretFanciersCentre = "170"

func TestListCSPs(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP": NewResponse(http.StatusOK, `
[
    {
	"Code":"170",
	"Name":"RCC Ferret Fanciers (TEST CLUB)",
	"WebName":"Ferrets",
	"Acronym":"RFF"
    }
]
		`),
	})
	got, err := ListCSPs(client)
	want := []CSPInfo{{
		Code:    "170",
		Name:    "RCC Ferret Fanciers (TEST CLUB)",
		WebName: "Ferrets",
		Acronym: "RFF",
	}}
	if err != nil {
		t.Errorf("ListCSPs(client): %v", err)
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("ListCSPs(client) = %#v\n%s", got, diff)
	}
}

func TestCSPDetails(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP/170": NewResponse(http.StatusOK, `
{
    "Code":"170",
    "Name":"RCC Ferret Fanciers (TEST CLUB)",
    "WebName":"Ferrets",
    "Acronym":"RFF"
}
		`),
	})
	got, err := CSPDetails(client, FerretFanciersCentre)
	want := CSPInfo{
		Code:    "170",
		Name:    "RCC Ferret Fanciers (TEST CLUB)",
		WebName: "Ferrets",
		Acronym: "RFF",
	}
	if err != nil {
		t.Errorf("CSPDetails(client, %q): %v", FerretFanciersCentre, err)
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("CSPDetails(client, %q) = %#v\n%s", FerretFanciersCentre, got, diff)
	}
}
