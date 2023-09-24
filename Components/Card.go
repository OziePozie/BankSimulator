package Components

import (
	"time"
)

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        Balance   `json:"balance"`
	History        History   `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}

func (card Card) blockCard() bool {
	card.IsCardActive = false

	return true
}

func (card Card) activateCard() bool {
	card.IsCardActive = true
	return true
}

func (bill Bill) getCard(id string) *Card {
	cards := bill.getCards()

	for _, card := range cards {
		if card.Number == id {

			return &card
		}
	}

	return nil
}
