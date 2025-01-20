package main

import (
	"demo/password/account"
	"demo/password/errors"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	getMenu()
}

func getMenu() {
	var variant string

MenuLoop: // Метка цикла
	for {
		fmt.Println("Выберите действие: ")
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выйти")

		fmt.Scan(&variant)

		switch variant {
		case "1":
			createAccount()
		case "2":
			fileName := ""
			fmt.Println("Введите имя аккаунта:")
			fmt.Scan(&fileName)
			findAccount(fileName)
		case "3":
			fieldName := ""
			fmt.Println("Введите название файла:")
			fmt.Scan(&fieldName)
			files.DeleteFile(fieldName)
		case "4":
			fmt.Println("Выход из программы.")
			break MenuLoop
		default:
			color.Red("Неверный выбор. Попробуйте снова.")
			continue MenuLoop
		}
	}
}

func findAccount(name string) {
	_, err := files.ReadFile("data.json")

	if err != nil {
		color.Red("Не удалось считать файл data.json")
	}

	vault := account.NewVault()

	for _, value := range vault.Accounts {
		if value.Login == name {
			fmt.Println(value)
			return
		}
	}

	color.Red("Такого аккаунта нет")
}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	account1, err := account.NewAccount(login, password, url)

	if err != nil {
		errors.GetErrorString(err)
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*account1)

	if err != nil {
		color.Red("С созданием файла что то не так!")
		return
	}

	account1.OutputPassword()
}

func promptData(prompt string) string {
	var res string
	for {
		fmt.Println(prompt)
		fmt.Scan(&res)
		if res != "" {
			break
		}

		color.Red("Поле не может быть пустым. Попробуйте снова.")
	}
	return res
}
