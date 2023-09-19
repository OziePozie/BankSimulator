package Components

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Account struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Bill       []Bill `json:"bills"`
}

func (account Account) createAccount(login, password, firstName, secondName string) Account {

	account.FirstName = firstName

	account.SecondName = secondName

	account.Login = login

	account.Password = password

	saveAccountToFile(account)

	return account
}

func (Account) deleteAccount(login string) bool {

	return true
}

func (account Account) updateAccount() {
	account.CreateBill()
}

func openDataBase() {

}

func createDateBase() {

}

func getAccountsFromDataBase() ([]Account, *os.File, error) {
	filename := "users.json"

	var accounts []Account

	file, err := os.Open(filename)

	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filename)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&accounts)

	return accounts, file, err
}

func getAccountByLogin(login string) Account {
	var account Account

	accounts, _, _ := getAccountsFromDataBase()

	for _, acc := range accounts {
		if acc.Login == login {
			account = acc
			break
		}
	}
	return account
}

func saveAccountToFile(account Account) {

	accounts, file, err := getAccountsFromDataBase()

	accounts = append(accounts, account)

	file, _ = os.OpenFile("users.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(accounts)

	if err != nil {
		fmt.Println("Ошибка при сохранении данных в файл:", err)
		return
	}

}
func RegisterAccount(scanner *bufio.Scanner) {
	fmt.Println("Регистрация пользователя")
	var account Account
	fmt.Print("Введите Имя: ")
	scanner.Scan()
	account.FirstName = scanner.Text()

	fmt.Print("Введите Фамилию: ")
	scanner.Scan()
	account.SecondName = scanner.Text()

	fmt.Print("Введите логин: ")
	scanner.Scan()
	account.Login = scanner.Text()

	fmt.Print("Введите пароль: ")
	scanner.Scan()
	account.Password = scanner.Text()

	account.createAccount(account.Login, account.Password,
		account.FirstName, account.SecondName)

	fmt.Println("Регистрация успешно завершена!")
}

func AuthAccount(login, password string) (Account, error) {
	account := getAccountByLogin(login)
	if account.Password == password {
		return account, nil
	} else {

		panic("Неправильные креды")

	}
}
