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
	Login    string `json:"login"`
	Password string	`json:"password"`
	Url      string `json:"url"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		Login:    login,
		Password: password,
		Url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

func (acc *Account) OutputPass() {
	color.Cyan(acc.Login)

	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.Password = string(res)
}