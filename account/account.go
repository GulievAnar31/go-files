package account

import (
	"demo/password/errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (account *Account) OutputPassword() {
	color.Green(account.Password)
}

func (account *Account) generatePassword(n int) {
	res := make([]rune, n)

	makeIndex := len(account.Password)

	if makeIndex == 0 {
		color.Red("Пароль пуст")
		return
	}

	for i := range res {
		res[i] = rune(account.Password[rand.IntN(makeIndex)])
	}

	account.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)

	if len(login) < 3 {
		return nil, errors.ErrInvalidLogin
	}

	if err != nil {
		return nil, errors.ErrInvalidURL
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(newAcc.Password) == 0 {
		newAcc.generatePassword(9)
	}

	return newAcc, nil
}
