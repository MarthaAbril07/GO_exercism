package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {

	var cont = 1
	var cadena = ""
	var siguiente = ""
	var num = ""
	for index, caracter := range input {

		if cont == 1 {
			num = ""
		} else {
			num = fmt.Sprint(cont)
		}

		if index+1 == len(input) {
			siguiente = string(caracter)
			cadena += num + string(caracter)
		} else {
			siguiente = string(input[index+1])
		}

		if siguiente != string(caracter) {
			cadena += num + string(caracter)
			cont = 1
		} else {
			cont++
		}
	}

	return cadena
}

func esNumero(ascci int32) bool {
	if ascci >= 48 && ascci <= 57 {
		return true
	}
	return false
}

func RunLengthDecode(input string) string {

	var cadena = ""
	var num = ""
	var siguiente = int32(0)
	for index, caracter := range input {

		if index+1 == len(input) {
			siguiente = caracter
		} else {
			siguiente = int32(input[index+1])

			if esNumero(caracter) == true && esNumero(siguiente) == true {
				num = num + string(caracter) + string(siguiente)
				//	fmt.Println("----Num " + num)
				continue
			} else {
				if num != "" {
					repeticiones, _ := strconv.ParseInt(num, 10, 32)
					cadena = getLetras(int(siguiente), int(repeticiones), cadena)
					num = ""
				}
			}

			if esNumero(caracter) == true && esNumero(siguiente) == false {
				repeticiones, _ := strconv.ParseInt(string(caracter), 10, 32)
				cadena = getLetras(int(siguiente), int(repeticiones), cadena)
			}

			if esNumero(caracter) == false && index == 0 {
				cadena = getLetras(int(caracter), 1, cadena)
			}
			if esNumero(caracter) == false && esNumero(siguiente) == false {
				cadena = getLetras(int(siguiente), 1, cadena)
			}

		}

	}
	return cadena
}

func getLetras(letra int, repeticiones int, cadena string) string {

	for x := 0; x < repeticiones; x++ {
		cadena += string(letra)
	}
	return cadena
}

func main() {
	/*	fmt.Println(RunLengthEncode(""))
		fmt.Println(RunLengthEncode("XYZ"))
		fmt.Println(RunLengthEncode("AABBBCCCC"))
		fmt.Println(RunLengthEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"))
		fmt.Println(RunLengthEncode("  hsqq qww  "))
		fmt.Println(RunLengthEncode("aabbbcccc"))*/
	fmt.Println(RunLengthDecode(""))
	fmt.Println(RunLengthDecode("XYZ"))
	fmt.Println(RunLengthDecode("2A3B4C"))
	fmt.Println(RunLengthDecode("12WB12W3B24WB"))
	fmt.Println(RunLengthDecode("2 hs2q q2w2 "))
	fmt.Println(RunLengthDecode("2a3b4c"))

}
