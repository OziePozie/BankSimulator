package Components

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Bill struct {
	Number       uuid.UUID `json:"number"`
	Balance      float64   `json:"balance"`
	Cards        []Card    `json:"cards"`
	History      History   `json:"history"`
	IsBillActive bool      `json:"isBillActive"`
}

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	History        History   `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}

type History struct {
	Date          time.Time `json:"date"`
	Destination   string    `json:"destination"`
	OperationType string    `json:"operationType"`
	Sum           float64   `json:"sum"`
}

func (account Account) CreateBill() Account {
	bills := account.Bill

	bill := Bill{
		Number:       uuid.UUID{},
		Balance:      0,
		Cards:        nil,
		History:      History{},
		IsBillActive: true,
	}
	card := bill.CreateCard()
	bill.Cards = append(bill.Cards, card)

	bills = append(bills, bill)

	saveAccountToFile(account)

	return account
}

func (bill Bill) CreateCard() Card {

	var number string
	for i := 0; i < 16; i++ {
		number += string(rand.Intn(10))
	}
	var cvv string
	for i := 0; i < 3; i++ {
		cvv += string(rand.Intn(10))
	}
	card := Card{
		Number:         number,
		Cvv:            cvv,
		ExpirationDate: time.Now().AddDate(4, 0, 0),
		Balance:        0,
		History:        History{},
		IsCardActive:   false,
	}

	return card
}

func (bill Bill) getCards() []Card {
	cards := bill.Cards

	return cards
}
