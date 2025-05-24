package account

import (
	"demo/app-1/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts    []Account `json:"accounts"`
	UpdatedTime time.Time `json:"updatedTime"`
}

func NewVault() *Vault{
	file, err := files.ReadFile("account.json")
	if err != nil {
		return &Vault{
			Accounts: []Account{}, 
			UpdatedTime: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
} 

func (vault *Vault) FindAccountUrl(url string) []Account {
	var accounts []Account

	for _, account := range vault.Accounts{
		isMatch := strings.Contains(account.Url, url)
		if isMatch {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *Vault) DeleteAccountUrl(url string) bool {
	var accounts []Account
	isDeleted := false 
	
	for _, account := range vault.Accounts{
		isMatch := strings.Contains(account.Url, url)
		if !isMatch {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}

	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdatedTime = time.Now()
	data, err := vault.ToBytes()

	if err != nil {
		color.Red("Не удалось преобразовать файл json")
	}

	files.WriteFile(data, "account.json")
}