package egotivities

import (
	"net/http"
	"testing"
	"time"

	"github.com/d4l3k/messagediff"
)

func TestCommittee(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP/170/reports/committee?year=14-15": NewResponse(http.StatusOK, `
[
    {
	"FirstName":"Joe",
	"Surname":"Bloggs",
	"CID":"00000000",
	"Email":"joe.bloggs50@imperial.ac.uk",
	"Login":"jbloggs50",
	"PostName":"Chief Ferret Fancier",
	"PhoneNo":"02075948060",
	"StartDate":"2014-08-01 00:00:00",
	"EndDate":"2015-07-31 23:59:59"
    }
]
				`),
	})
	got, err := Committee(client, FerretFanciersCentre, "14-15")
	want := []CommitteeMember{{
		FirstName: "Joe",
		Surname:   "Bloggs",
		CID:       "00000000",
		Email:     "joe.bloggs50@imperial.ac.uk",
		Login:     "jbloggs50",
		PostName:  "Chief Ferret Fancier",
		PhoneNo:   "02075948060",
		StartDate: Time(time.Date(2014, 8, 1, 0, 0, 0, 0, eActivitiesLocation)),
		EndDate:   Time(time.Date(2015, 7, 31, 23, 59, 59, 0, eActivitiesLocation)),
	}}
	if err != nil {
		t.Errorf("Committee(...): %v", err)
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("Committee(...) = %#v\n%s", want, diff)
	}
}

func TestMembers(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP/170/reports/members?year=14-15": NewResponse(http.StatusOK, `
[
    {
	"FirstName":"Joe",
	"Surname":"Bloggs",
	"CID":"00000000",
	"Email":"joe.bloggs50@imperial.ac.uk",
	"Login":"jbloggs50",
	"OrderNo":1000,
	"MemberType":"Full"
    }
]
				`),
	})
	got, err := Members(client, FerretFanciersCentre, "14-15")
	want := []Member{{
		FirstName:  "Joe",
		Surname:    "Bloggs",
		CID:        "00000000",
		Email:      "joe.bloggs50@imperial.ac.uk",
		Login:      "jbloggs50",
		OrderNo:    1000,
		MemberType: "Full",
	}}
	if err != nil {
		t.Errorf("Members(...): %v", err)
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("Members(...) = %#v\n%s", got, diff)
	}
}

func TestOnlineSales(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP/170/reports/onlinesales?year=14-15": NewResponse(http.StatusOK, `
[
    {
	"OrderNumber":"1000",
	"SaleDateTime":"2015-06-20 19:00:00",
	"ProductID":1234,
	"ProductLineID":4567,
	"Price":30,
	"Quantity":1,
	"QuantityCollected":0,
	"Customer":
	{
		"FirstName":"Joe",
		"Surname":"Bloggs",
		"CID":"00000000",
		"Email":" joe.bloggs50@imperial.ac.uk ",
		"Login":"jbloggs50"
	},
	"VAT":
	{
		"Code":"S1",
		"Name":"S1 – Sales Standard Rated",
		"Rate":20
	}
    }
]
				`),
	})
	got, err := OnlineSales(client, FerretFanciersCentre, "14-15")
	want := []OnlineSale{{
		OrderNumber:       "1000",
		SaleDateTime:      Time(time.Date(2015, 6, 20, 19, 00, 00, 0, eActivitiesLocation)),
		ProductID:         1234,
		ProductLineID:     4567,
		Price:             newMoney(30),
		Quantity:          1,
		QuantityCollected: 0,
		Customer: Customer{
			FirstName: "Joe",
			Surname:   "Bloggs",
			CID:       "00000000",
			Email:     " joe.bloggs50@imperial.ac.uk ",
			Login:     "jbloggs50",
		},
		VAT: VAT{
			Code: "S1",
			Name: "S1 – Sales Standard Rated",
			Rate: newMoney(20),
		},
	}}
	if err != nil {
		t.Errorf("OnlineSales(...): %v", err)
	}
	for n, w := range want {
		g := got[n]
		if g.Price.Cmp(&w.Price.Float) != 0 {
			t.Errorf("OnlineSales(...)[%d].Price = %s; want %s", n, g.Price.String(), w.Price.String())
		}
		want[n].Price = &Money{g.Price.Float}
		if g.VAT.Rate.Cmp(&w.VAT.Rate.Float) != 0 {
			t.Errorf("OnlineSales(...)[%d].VAT.Rate = %s; want %s", n, g.VAT.Rate.String(), w.VAT.Rate.String())
		}
		want[n].VAT.Rate = &Money{g.VAT.Rate.Float}
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("OnlineSales(...) = %#v\n%s", got, diff)
	}
}

func TestTransactionLines(t *testing.T) {
	client, _ := newTestClient(map[string]*http.Response{
		"/CSP/170/reports/transactionlines?year=14-15": NewResponse(http.StatusOK, `
[
    {
	"TransID": 234567,
	"TransDate": "2015-06-20",
	"Document": "CF 12345 (234567)",
	"Description": "Pens and card for making signs",
	"Amount": -234,
	"Funding":
	{
		"Code": "0",
		"Name": "Grant (0)"
	},
	"Activity":
	{
		"Code": "00",
		"Name": "General (0)"
	},
	"Account":
	{
		"Code": "860",
		"Name": "Stationery (860)",
		"Type": "Expenditure"
	},
	"Pending": true,
	"Outstanding": false
    }
]
				`),
	})
	got, err := TransactionLines(client, FerretFanciersCentre, "14-15")
	want := []TransactionLine{{
		TransID:     234567,
		TransDate:   Time(time.Date(2015, 6, 20, 0, 0, 0, 0, eActivitiesLocation)),
		Document:    "CF 12345 (234567)",
		Description: "Pens and card for making signs",
		Amount:      newMoney(-234),
		Funding: Funding{
			Code: "0",
			Name: "Grant (0)",
		},
		Activity: Activity{
			Code: "00",
			Name: "General (0)",
		},
		Account: Account{
			Code: "860",
			Name: "Stationery (860)",
			Type: "Expenditure",
		},
		Pending:     true,
		Outstanding: false,
	}}
	if err != nil {
		t.Errorf("TransactionLines(...): %v", err)
	}
	for n, w := range want {
		g := got[n]
		if g.Amount.Cmp(&w.Amount.Float) != 0 {
			t.Errorf("TransactionLines(...)[%d].Amount = %s; want %s", n, g.Amount.String(), w.Amount.String())
		}
		want[n].Amount = &Money{g.Amount.Float}
	}
	if diff, equal := messagediff.PrettyDiff(want, got); !equal {
		t.Errorf("TransactionLines(...) = %#v\n%s", got, diff)
	}
}
