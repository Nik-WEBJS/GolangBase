package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	createdTime time.Time
	updatedTime time.Time
	Account
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NewAccountWithTimestamp(login, password, urlString string) (*AccountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimestamp{
		createdTime: time.Now(),
		updatedTime: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
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

func (acc *Account) OutputPass() {
	color.Cyan(acc.login)

	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

