package output

import "github.com/fatih/color"

// interface{} = any = any alias
func PrintError(value any) {
	stringVal, ok := value.(string) // вот так можно попробовать преобразовать в определенный тип
	if ok {
		color.Red(stringVal)
	}
	intVal, ok := value.(int)
	if ok {
		color.Red("Код ошибки %d", intVal)
	}
	errVal, ok := value.(error)
	if ok {
		color.Red(errVal.Error())
		return
	}
}

func sum[T int | float64 | float32](a, b T) T {  // так используются дженерики
	return a + b
}
