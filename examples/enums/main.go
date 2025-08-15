package main

import "fmt"

type BillingStatus int

const (
	OPEN BillingStatus = iota
	PAID
	CANCELLED
)

var statusName = map[BillingStatus]string{
	OPEN:      "OPEN",
	PAID:      "PAID",
	CANCELLED: "CANCELLED",
}

func (bs BillingStatus) String() string {
	return statusName[bs]
}

type Billing struct {
	Status BillingStatus
	Value  float64
}

func NewBilling() *Billing {
	return &Billing{
		Status: OPEN,
		Value:  0,
	}
}

func main() {
	billing := NewBilling()
	fmt.Println(billing)
}
