package Components

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

type Bill struct {
	Number       uuid.UUID `json:"number"`
	Balance      Balance   `json:"balance"`
	Cards        []Card    `json:"cards"`
	History      []History `json:"history"`
	Limit        int       `json:"limit"`
	IsBillActive bool      `json:"isBillActive"`
}

func (account Account) CreateBill() Account {
	bills := account.Bill

	bill := Bill{
		Number: uuid.UUID{},
		Balance: Balance{
			Dollars: 0,
			Euros:   0,
			Rubles:  0,
		},
		Cards:        nil,
		History:      nil,
		Limit:        0,
		IsBillActive: true,
	}
	//card := bill.CreateCard()
	//bill.Cards = append(bill.Cards, card)

	bills = append(bills, bill)

	account.Bill = bills

	saveAccountToFile(account)

	return account
}

func (bill Bill) CreateCard() Card {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var number string
	for i := 0; i < 16; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	var cvv string
	for i := 0; i < 3; i++ {
		cvv += strconv.Itoa(r.Intn(10))
	}
	card := Card{
		Number:         number,
		Cvv:            cvv,
		ExpirationDate: time.Now().AddDate(4, 0, 0),
		Balance: Balance{
			Dollars: 0,
			Euros:   0,
			Rubles:  0,
		},
		History:      nil,
		IsCardActive: true,
	}

	return card
}

func (bill Bill) getCards() []Card {
	cards := bill.Cards

	return cards
}

func (bill Bill) closeBill() Bill {
	bill.IsBillActive = false

	return bill
}

func (bill Bill) setLimit(limit int) Bill {
	bill.Limit = limit
	return bill
}

func (bill Bill) addHistory(history History) Bill {

	bill.History = append(bill.History, history)

	return bill

}

func (bill Bill) getBalance() Balance {

	return bill.Balance

}
