package main

import (
	"awesomeProject4/Components"
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var account Components.Account

func main() {
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Регистрация" +
		"2. Вход")
	scanner.Scan()
	switcher := scanner.Text()
	entryListener(switcher)

	mainLogicListener()
}

func entryListener(switcher string) {

	switch switcher {
	case "1":
		Components.RegisterAccount(scanner)
		entryListener("2")
	case "2":
		fmt.Println("Введите логин")
		scanner.Scan()
		login := scanner.Text()
		fmt.Println("Введите пароль")
		scanner.Scan()
		password := scanner.Text()
		account, _ = Components.AuthAccount(login, password)

	}
}
func mainLogicListener() {
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Карты" +
		"2. Счета" +
		"3. История")
	scanner.Scan()
	switcher := scanner.Text()
	switch switcher {
	case "1":
		cardListener()
	case "2":
		billListener()
	case "3":
		historyListener()
	default:

	}

}

func historyListener() {
	scanner.Scan()
	switcher := scanner.Text()
	switch switcher {
	case "1":
		historyCardListener()
	case "2":
		historyBillListener()
	default:
		mainLogicListener()
	}
}

func historyBillListener() {
	bills := account.Bill
	for _, bill := range bills {
		fmt.Println(bill.History)
	}
	mainLogicListener()
}

func historyCardListener() {
	cards := account.Bill[0].Cards
	for _, card := range cards {
		fmt.Println(card.History)
	}
	mainLogicListener()
}

func billListener() {

	if account.Bill == nil {
		fmt.Println("У вас нет счетов, создать ?" +
			"1. Da" +
			"2. Net")
		scanner.Scan()
		switcher := scanner.Text()
		switch switcher {
		case "1":
			account.CreateBill()
		case "2":
			mainLogicListener()
		}

	}
	fmt.Println("Выберите что вы хотите сделать:" +
		"1. Посмотреть список карт" +
		"2. Создать новый счет " +
		"3. Установить лимит для счета" +
		"4. Посмотреть баланс на счете" +
		"5. Закрыть счет")
	scanner.Scan()
	switcher := scanner.Text()
	switch switcher {
	case "1":
		bills := account.Bill
		for _, bill := range bills {

			cards := bill.Cards
			if cards == nil {
				fmt.Println("У вас нет карт, создать ?" +
					"1. Da" +
					"2. Net")
				scanner.Scan()
				switcherok := scanner.Text()
				switch switcherok {
				case "1":
					account.Bill[0].CreateCard()
					fallthrough
				case "2":
					mainLogicListener()
				}
			}
			for _, card := range cards {
				fmt.Printf("Номер карты %s, CVV: %s, Дата окончания: %s \n EURO: %.2f \n RUB: %.2f \n DOL: %.2f \n",
					card.Number, card.Cvv,
					card.ExpirationDate, card.Balance.Euros,
					card.Balance.Rubles, card.Balance.Dollars)
			}
		}

	case "2":
		account.CreateBill()
	case "3":
	default:
		mainLogicListener()
	}

}

func cardListener() {

	scanner.Scan()
	switcher := scanner.Text()
	switch switcher {
	case "1":
		account.CreateBill()
	case "2":
		mainLogicListener()
	}

}
