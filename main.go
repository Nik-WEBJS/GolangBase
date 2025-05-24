package main

import (
	"demo/app-1/account"
	"fmt"

	"github.com/fatih/color"
)

func main() { 
	fmt.Println("___Менеджер паролей___")
	vault := account.NewVault()

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1: 
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func findAccount(vault *account.Vault) {
	url:= promtData("Введите url: ")
	accounts := vault.FindAccountUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунты не найдены")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault){
	url:= promtData("Введите url: ")
	isDeleted := vault.DeleteAccountUrl(url)
	if isDeleted {
		color.Green("Аккаунт успешно удален")
	} else {
		color.Red("Аккаунт не найден")
	}
}

func createAccount(vault *account.Vault) {
	login:=promtData("Введите логин: ")
	password:=promtData("Введите пароль: ")
	url:= promtData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
		vault.AddAccount(*myAccount)
}

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

