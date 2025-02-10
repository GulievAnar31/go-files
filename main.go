package main

import (
	"demo/password/account"
	"demo/password/errors"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"
)

func main() {
	getMenu()
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

func findAccountByUrl(url string) {
	jsonDb := files.NewJsonDb("data.json")
	_, err := jsonDb.ReadFile()

	if err != nil {
		output.PrintError("Не удалось считать файл data.json")
	}

	vault := account.NewVault(jsonDb)

	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			fmt.Println(value)
		}
	}

	output.PrintError("Это все аккаунт который нашлись")
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
	getMenu()
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
		fmt.Scan(&res)
		if res != "" {
			break
		}

		output.PrintError("Поле не может быть пустым. Попробуйте снова.")
	}
	return res
}

func getMenu() {
	var menu = map[string]interface{}{
		"1": createAccount,
		"2": findAccount,
		"3": findAccountByUrl,
		"4": deleteAccount,
	}

	variant := promptData([]string{
		"Выберите действие: ",
		"1. Создать аккаунт",
		"2. Найти аккаунт по логину",
		"3. Найти аккаунт по URL",
		"4. Удалить аккаунт",
		"5. Выйти",
	})

	for variant != "5" {
		if action, exists := menu[variant]; exists {
			switch fn := action.(type) {
			case func(): // Функция без аргументов
				fn()
			case func(string): // Функция с аргументом
				var input string
				fmt.Print("Введите значение: ")
				fmt.Scan(&input)
				fn(input)
			default:
				fmt.Println("Ошибка: неподдерживаемый тип функции")
			}
		} else {
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}

		// Запрос нового варианта после выполнения команды
		variant = promptData([]string{
			"Выберите действие: ",
			"1. Создать аккаунт",
			"2. Найти аккаунт по логину",
			"3. Найти аккаунт по URL",
			"4. Удалить аккаунт",
			"5. Выйти",
		})
	}
}
