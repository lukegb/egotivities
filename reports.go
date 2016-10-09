package eactivities

import "fmt"

// CommitteeMember represents a single member of a CSP committee.
type CommitteeMember struct {
	FirstName string
	Surname   string
	CID       string
	Email     string
	Login     string
	PostName  string
	PhoneNo   string
	StartDate Time
	EndDate   Time
}

// Committee returns the list of committee members for a given year.
func Committee(c Client, centre string, year string) ([]CommitteeMember, error) {
	var committee []CommitteeMember
	if err := c.Get(fmt.Sprintf("/CSP/%s/reports/committee?year=%s", centre, year), &committee); err != nil {
		return nil, err
	}
	return committee, nil
}

// Member represents a single member of a CSP.
type Member struct {
	FirstName  string
	Surname    string
	CID        string
	Email      string
	Login      string
	OrderNo    uint64
	MemberType string
}

// Members returns the list of CSP members for a given year.
func Members(c Client, centre string, year string) ([]Member, error) {
	var members []Member
	if err := c.Get(fmt.Sprintf("/CSP/%s/reports/members?year=%s", centre, year), &members); err != nil {
		return nil, err
	}
	return members, nil
}

// Customer represents the customer information for a CSP's sale.
type Customer struct {
	FirstName string
	Surname   string
	CID       string
	Email     string
	Login     string
}

// VAT represents VAT information for a product line.
type VAT struct {
	Code string
	Name string
	Rate float64
}

// OnlineSale represents a single online sale of a product.
type OnlineSale struct {
	OrderNumber       string
	SaleDateTime      Time
	ProductID         uint64
	ProductLineID     uint64
	Price             float64
	Quantity          uint
	QuantityCollected uint
	Customer          Customer
	VAT               VAT
}

// OnlineSales returns the list of online sales for a CSP's products for a given year.
func OnlineSales(c Client, centre string, year string) ([]OnlineSale, error) {
	var sales []OnlineSale
	if err := c.Get(fmt.Sprintf("/CSP/%s/reports/onlinesales?year=%s", centre, year), &sales); err != nil {
		return nil, err
	}
	return sales, nil
}

// Funding represents a funding code.
type Funding struct {
	Code string
	Name string
}

// Activity represents an activity code.
type Activity struct {
	Code string
	Name string
}

// Account represents an account.
type Account struct {
	Code string
	Name string
	Type string
}

// TransactionLine represents a single transaction listed in eActivities.
type TransactionLine struct {
	TransID     uint64
	TransDate   Time
	Document    string
	Description string
	Amount      float64
	Funding     Funding
	Activity    Activity
	Account     Account
	Pending     bool
	Outstanding bool
}

// TransactionLines returns a CSP's transaction lines for a given year.
func TransactionLines(c Client, centre string, year string) ([]TransactionLine, error) {
	var transactions []TransactionLine
	if err := c.Get(fmt.Sprintf("/CSP/%s/reports/transactionlines?year=%s", centre, year), &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
