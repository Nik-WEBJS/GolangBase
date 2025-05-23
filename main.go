package main

import (
	"fmt"
	"demo/app-1/account"
	"demo/app-1/files"
)

func main() { 
	login:=promtData("Введите логин: ")
	password:=promtData("Введите пароль: ")
	url:= promtData("Введите url: ")

	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	myAccount.OutputPass()
	files.WriteFile()
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

