package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		color.Red("При чтении произошла ошибка")
		return nil, err
	}

	return data, nil
}

func DeleteFile(name string) {
	err := os.Remove(name)

	if err != nil {
		color.Red("Произошла ошибка удаления")
	}

	color.Green("Файл удален")
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return
	}

	color.Green("Запись успешна")
}
