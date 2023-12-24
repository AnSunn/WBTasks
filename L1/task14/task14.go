package task14

import "fmt"

func Launch() {

	var a interface{} = 123.1
	fmt.Printf("First option: The type is %s\n", VariableType(a))
	fmt.Printf("The second option: The type is %T\n", a)

}

func VariableType(value interface{}) string {
	switch value.(type) {
	case int, int64, int32:
		return "int"
	case float64, float32:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan bool, chan int, chan float64:
		return "chan"
	default:
		return "unknown"
	}
}
