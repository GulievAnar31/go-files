package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		color.Red("При чтении произошла ошибка")
		return nil, err
	}

	return data, nil
}

func (db *JsonDb) DeleteFile() {
	err := os.Remove(db.fileName)

	if err != nil {
		color.Red("Произошла ошибка удаления")
	}

	color.Green("Файл удален")
}

func (db *JsonDb) WriteFile(content []byte) error {
	file, err := os.Create(db.fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return err
	}

	color.Green("Запись успешна")
	return nil
}
