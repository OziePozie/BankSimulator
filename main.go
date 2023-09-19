package main

import (
	"awesomeProject4/Components"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//Components.RegisterAccount(scanner)

	fmt.Println("Введите логин")
	scanner.Scan()
	login := scanner.Text()
	fmt.Println("Введите пароль")
	scanner.Scan()
	password := scanner.Text()
	account, _ := Components.AuthAccount(login, password)

	account.CreateBill()
}
