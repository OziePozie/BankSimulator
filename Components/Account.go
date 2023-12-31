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

var filename = "users.json"

func (account Account) createAccount(login, password, firstName, secondName string) Account {

	account.FirstName = firstName

	account.SecondName = secondName

	account.Login = login

	account.Password = password

	saveAccountToFile(account)

	return account
}

func (Account) deleteAccount(login string) bool {
	var res string
	accounts := getAccountsFromDataBase()
	for i, acc := range accounts {
		if login == acc.Login {
			accounts = append(accounts[:i], accounts[i+1:]...)

		}
	}

	for _, account := range accounts {
		res += account.StructToJSON()
	}
	file, _ := os.Open("users.json")

	file.WriteString(res)
	return true
}

func (account Account) updateAccount(login string) {
	accounts := getAccountsFromDataBase()
	for i, acc := range accounts {
		if login == acc.Login {
			accounts[i] = acc

		}
	}

	_, err := json.Marshal(account)
	if err != nil {
		fmt.Println("Произошел ")
	}

}

func getBytesDataBase() []byte {
	file, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filename)
	}
	b, _ := os.ReadFile(filename)
	defer file.Close()
	return b
}

func createDateBase() {

}

func getAccountsFromDataBase() []Account {
	var accounts []Account

	err := json.Unmarshal(getBytesDataBase(), &accounts) //!!!!
	if err != nil {
		return nil
	}
	return accounts
}

func GetAccountByLogin(login string) Account {
	var account Account

	accounts := getAccountsFromDataBase()

	for _, acc := range accounts {
		if acc.Login == login {
			account = acc
			break
		}
	}
	return account
}

func saveAccountToFile(account Account) {

	accounts := getAccountsFromDataBase()

	flag := false

	for i, a := range accounts {
		if a.Login == account.Login {
			accounts[i] = account
			flag = true
			fmt.Println(a)
			break
		}
	}
	if !flag {
		accounts = append(accounts, account)
	}

	file, _ := os.OpenFile("users.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Println(accounts)
	encoder := json.NewEncoder(file)
	err := encoder.Encode(accounts)

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
	account := GetAccountByLogin(login)
	if account.Password == password {
		return account, nil
	} else {

		panic("Неправильные креды")

	}
}

func (account Account) JSONToStruct() Account {
	b, _ := os.ReadFile(filename)
	_ = json.Unmarshal(b, &account)
	return account
}

func (account Account) StructToJSON() string {
	b, err := json.Marshal(account)
	if err != nil {

	}
	return string(b)
}
