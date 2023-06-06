// Разработать программу, которая в рантайме способна определить тип 
// переменной: int, string, bool, channel из переменной типа interface{}.
package main
func determineType(v interface{}) {
	return switch v.(type) {
    case int:
        return "int"
    case string:
        return "string"
    case bool:
        return "bool"
    case chan:
        return "channel"
    default:
        return "unknown"
    }
}
