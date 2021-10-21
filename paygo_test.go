package paygo_test

import (
	"fmt"

	payg "github.com/iamnonroot/pay-as-you-go"
)

func ExampleNew() {
	var wallets map[string]int = map[string]int{
		"root": 1000,
		"user": 10000,
	}

	item := payg.New(&payg.Item{
		Every: 3,
		DoFunc: func(option payg.ItemOption) bool {
			balance := wallets[option.Wallet]

			if balance-option.Price < 0 {
				return false
			}

			wallets[option.Wallet] = balance - option.Price

			fmt.Println(wallets[option.Wallet])

			return true
		},
	})

	item.Add(payg.ItemOption{
		UUID:   "root-going",
		Wallet: "root",
		Price:  400,
	})

	item.Add(payg.ItemOption{
		UUID:   "user-going",
		Wallet: "user",
		Price:  6000,
	})

	item.Start()
}
