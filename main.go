package main

import (
	"fmt"
	"demo/app-1/account"
	"demo/app-1/files"
)

func main() { 
	fmt.Println("___Менеджер паролей___")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1: 
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1 Создать аккаунт:")
	fmt.Println("2 Найти аккаунт")
	fmt.Println("3 Удалить аккаунт")
	fmt.Println("4 Выход")

	fmt.Scan(&variant)
	var discard string
	fmt.Scanln(&discard)
	return variant

}

func findAccount(){

}

func deleteAccount(){

}

func createAccount() {
	login:=promtData("Введите логин: ")
	password:=promtData("Введите пароль: ")
	url:= promtData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	
	if err != nil {
		fmt.Println("Ошибка преобразования JSON:", err)
		return
	}

	files.WriteFile(data, "account.json")
}

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

