package main

import "fmt"

func Convert(number int) string {
	var cadena = ""
	if number%3 == 0 {
		cadena = cadena + "Pling"
	}
	if number%5 == 0 {
		cadena = cadena + "Plang"
	}
	if number%7 == 0 {
		cadena = cadena + "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		cadena = fmt.Sprint(number)
	}
	return cadena
}

func main() {
	fmt.Println(Convert(1))
}
