package account

import (
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	ReadFile() ([]byte, error)
	WriteFile(content []byte) error
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
	db        Db
}

func NewVault(db Db) *Vault {
	file, err := db.ReadFile()
	if err != nil {
		color.Yellow("Ошибка чтения базы данных: %v", err)
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Ошибка парсинга базы данных: %v", err)
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
			db:        db,
		}
	}

	vault.db = db
	return &vault
}

func (v *Vault) AddAccount(acc Account) error {
	v.Accounts = append(v.Accounts, acc)
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		return err
	}
	v.db.WriteFile(data)
	color.Green("Файл записан!")
	return nil
}

func (v *Vault) ToBytes() ([]byte, error) {
	return json.Marshal(v)
}
