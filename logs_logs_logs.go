package main

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var code = ""
	for _, char := range log {
		x := fmt.Sprintf("%U", char)
		if x == "U+2757" || x == "U+1F50D" || x == "U+2600" {
			code = x
			break
		}
	}
	switch code {
	case "U+2757":
		return "recommendation"
	case "U+1F50D":
		return "search"
	case "U+2600":
		return "weather"
	default:
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var modificada = ""
	for _, value := range log {
		if value == oldRune {
			modificada += string(newRune)
		} else {
			modificada += string(value)
		}
	}
	return modificada
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	fmt.Println(utf8.RuneCountInString(log))
	if utf8.RuneCountInString(log) <= limit {
		return true
	} else {
		return false
	}
}

func main() {
	//Application("h")
	//Replace("Hola", 72, 192)
	fmt.Println(WithinLimit("exercism❗", 9))
	fmt.Println(WithinLimit("exercism❗", 10))
	fmt.Println(WithinLimit("exercism❗", 8))
	fmt.Println(WithinLimit("exercism❗", 9))

}
