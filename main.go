package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)
type account struct{
	login string
	password string
	url string
}

type accountWithTimestamp struct{
	createdTime time.Time
	updatedTime time.Time
	account
}

func newAccountWithTimestamp(login, password, urlString string) (*accountWithTimestamp, error){
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimestamp{
		createdTime: time.Now(),
		updatedTime: time.Now(),
		account: account{
			login: login,
			password: password,
			url: urlString,
		},
		
	}

	if password == "" {
		newAcc.generatePassword(10)
	}
	return newAcc, nil 
}

// func newAccount(login, password, urlString string) (*account, error){
// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}
// 	newAcc := &account{
// 		login: login,
// 		password: password,
// 		url: urlString,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(10)
// 	}
// 	return newAcc, nil 
// }

func (acc *account) outputPass(){
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	} 

	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() { 
	login:=promtData("Введите логин: ")
	password:=promtData("Введите пароль: ")
	url:= promtData("Введите url: ")

	myAccount, err := newAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	myAccount.outputPass()
}

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

