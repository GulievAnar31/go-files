package main

import (
	"demo/password/account"
	"demo/password/errors"
	"demo/password/files"
	"demo/password/output"
	"fmt"
)

func main() {
	getMenu()
}

func getMenu() {
	variant := promptData([]string{
		"Выберите действие: ",
		"1. Создать аккаунт",
		"2. Найти аккаунт",
		"3. Удалить аккаунт",
		"4. Выйти",
	})
MenuLoop: // Метка цикла
	for {
		// fmt.Scan(&variant)
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
			fmt.Println("Введите имя пользователя:")
			fmt.Scan(&fieldName)
			deleteAccount(fieldName)
		case "4":
			fmt.Println("Выход из программы.")
			break MenuLoop
		default:
			output.PrintError("Неверный выбор. Попробуйте снова.")
			continue MenuLoop
		}
	}
}

func findAccount(name string) {
	jsonDb := files.NewJsonDb("data.json")
	_, err := jsonDb.ReadFile()

	if err != nil {
		output.PrintError("Не удалось считать файл data.json")
	}

	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("data.json"))
	// пример того как можно использовать di и заменять что то одно другим

	for _, value := range vault.Accounts {
		if value.Login == name {
			fmt.Println(value)
			return
		}
	}

	output.PrintError("Такого аккаунта нет")
}

func deleteAccount(name string) {
	jsonDb := files.NewJsonDb("data.json")
	vault := account.NewVault(files.NewJsonDb("data.json"))

	newAccounts := []account.Account{}

	for _, value := range vault.Accounts {
		if value.Login == name {
			continue
		}
		newAccounts = append(newAccounts, value)
	}

	vault.Accounts = newAccounts
	newData, err := vault.ToBytes()

	if err != nil {
		output.PrintError("Не удалось преобразовать в байты")
	}

	jsonDb.WriteFile(newData)
}

func createAccount() {
	login := promptData([]string{"Введите логин: "})
	password := promptData([]string{"Введите пароль: "})
	url := promptData([]string{"Введите url: "})

	account1, err := account.NewAccount(login, password, url)

	if err != nil {
		errors.GetErrorString(err)
		return
	}

	vault := account.NewVault(files.NewJsonDb("data.json"))
	vault.AddAccount(*account1)

	if err != nil {
		output.PrintError("С созданием файла что то не так!")
		return
	}

	account1.OutputPassword()
}

func promptData[T any](prompt []T) string {
	for i, val := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", val)
		} else {
			fmt.Println(val)
		}
	}
	var res string
	for {
		fmt.Println(prompt)
		fmt.Scan(&res)
		if res != "" {
			break
		}

		output.PrintError("Поле не может быть пустым. Попробуйте снова.")
	}
	return res
}
