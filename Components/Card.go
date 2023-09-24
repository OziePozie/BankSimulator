package Components

import (
	"time"
)

type Card struct {
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        Balance   `json:"balance"`
	History        []History `json:"history"`
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

func (card Card) transferToCard(currency string, sum float64, login string) bool {
	userToTransfer := GetAccountByLogin(login)
	billToTransfer := userToTransfer.Bill
	cardToTransfer := billToTransfer[0].Cards[0]
	switch currency {
	case "RUB":
		cardToTransfer.Balance.Rubles += sum
		card.Balance.Rubles -= sum
	case "DOL":
		cardToTransfer.Balance.Dollars += sum
		card.Balance.Dollars -= sum
	case "EU":
		cardToTransfer.Balance.Euros += sum
		card.Balance.Euros -= sum
	}
	card.History = append(card.History, createHistoryField(cardToTransfer.Number, "transfer", sum))

	return true
}
