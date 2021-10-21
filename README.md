# Pay as you Go
This is a small and simple library for Golang to scheduler deduct user's wallet balance every X seconds.

- [What is **P**ay **A**s **Y**ou **G**o ?](#What-is-Pay-As-You-Go-)
- [Concept](#Concept)
- [Example](#Example)

## What is **P**ay **A**s **Y**ou **G**o ?
One way to get money from a user is to deduct some money from the user's wallet based on a specific time, and if the user's wallet does not have money, we will no longer provide the service.

## Concept
- **DoFunc** is a function for each X seconds when it comes time to pay them. why you need yo pass it? Because you want to reduce the amount of RAM consumed for each item.
- **UUID** is a unique id for a new pay as you go item. You can set it any things but it must be unique like: user transaction id in your database.  
- **Wallet** is user wallet id and why you need to pass it? Because by passing it to DoFunc you can get your user's wallet balance and update it.  

## Example
In this example we use wallets that are in variable. We deduct wallet balance and print that every 10s.

```go
package main

import (
	"fmt"

	payg "github.com/iamnonroot/pay-as-you-go"
)

type Wallets map[string]int

func main() {
    // [wallet id] => wallet balance balance
	var wallets Wallets = Wallets{
		"root": 1000,
		"user": 10000,
	}

	item := payg.New(&payg.Item{
		Every: 10, // every 10 seconds
		DoFunc: func(option payg.ItemOption) bool {
            // get wallet balance
			balance := wallets[option.Wallet]

            // if the wallet balance is low, delete scheduler by returning false
			if balance-option.Price < 0 {
				return false
			}

            // set wallet balance
			wallets[option.Wallet] = balance - option.Price

			fmt.Println(wallets[option.Wallet])

            // continue
			return true
		},
	})

	item.Add(payg.ItemOption{
		UUID:   "root-going", // a uuid for pay as you go
		Wallet: "root", // user's wallet id
		Price:  400, // product price
	})

	item.Add(payg.ItemOption{
		UUID:   "user-going",
		Wallet: "user",
		Price:  6000,
	})

	item.Start()
}

```