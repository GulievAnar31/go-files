package output

import "github.com/fatih/color"

// interface{} = any = any alias
func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
		break
	case int:
		color.Red("Код ошибки %d", t)
		break
	default:
		color.Red("Неизвестный код ошибки")
	}
}
