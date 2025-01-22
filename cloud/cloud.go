package cloud

import (
	"fmt"
)

type CloudDb struct {
	FilePath string
}

func NewCloudDb(filePath string) *CloudDb {
	return &CloudDb{FilePath: filePath}
}

func (c *CloudDb) ReadFile() ([]byte, error) {
	fmt.Printf("Чтение из облачной базы данных: %s\n", c.FilePath)
	return []byte{}, nil
}

func (c *CloudDb) WriteFile(content []byte) error {
	fmt.Printf("Запись в облачную базу данных: %s\n", c.FilePath)
	return nil
}
