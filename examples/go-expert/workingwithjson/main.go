package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BankAccount struct {
	AccountNumber string  `json:"accountNumber"`
	Value         float64 `json:"value"`
}

func NewBankAccount() BankAccount {
	return BankAccount{
		AccountNumber: "985462",
		Value:         1000.0,
	}
}

func main() {
	ba := NewBankAccount()
	// marshal
	res, err := json.Marshal(ba)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// encode
	// create a new encoder and print to console
	err = json.NewEncoder(os.Stdout).Encode(ba)
	if err != nil {
		panic(err)
	}

	// unmarshal
	raw := []byte(`{"accountNumber": "546960", "value": 1000.69}`)
	var bankAccount BankAccount

	err = json.Unmarshal(raw, &bankAccount)
	if err != nil {
		panic(err)
	}
	fmt.Println(bankAccount)
}
