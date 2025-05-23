package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts    []Account `json:"accounts"`
	UpdatedTime time.Time `json:"updatedTime"`
}

func NewVault() *Vault{
	return &Vault{
		Accounts: []Account{}, 
		UpdatedTime: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedTime = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
} 