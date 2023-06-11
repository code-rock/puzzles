// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.
package main

import "reflect"

func DetermineType(v interface{}) string {
	// Переключатель типов
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func DetermineGenericType(x interface{}) string {
	// Запись в переменную нового значения, инициализированного конкретным значением,
	// хранящимся в интерфейсе x
	v := reflect.ValueOf(x)
	// Получение конкретного типа и возврат соответственного значения
	switch v.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		return "uint"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "slice"
	case reflect.Map:
		return "map"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {}
