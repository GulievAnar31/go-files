package account

import (
	"demo/password/files"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json: "accounts"`
	UpdatedAt time.Time `json: "updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())
		color.Red("Не удалось прочитать data.json")

		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()

	if err != nil {
		color.Red(err.Error())
		color.Red("Не удалось преобразовать файл")
	}

	files.WriteFile(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return data, nil
}
