package main

import (
	"fmt"
	"strconv"
)

func Valid(id string) bool {

	numeros := []int{}
	for _, i := range id {
		if i != 32 {
			v, err := strconv.ParseInt(string(i), 0, 32)
			if err != nil {
				return false
			}
			numeros = append(numeros, int(v))
		}
	}
	if len(numeros) == 1 && numeros[0] == 0 {
		return false
	}
	var suma = 0
	//fmt.Println("antes", numeros)
	var bandera = false
	for x := len(numeros) - 1; x >= 0; x-- {
		//fmt.Println("\tsuma: ", suma, "Bandera:", bandera)
		i := numeros[x]
		if !bandera {
			suma += numeros[x]
			bandera = !bandera
			continue
		} else {
			//fmt.Println("Pos: ", x, "Entra:", i)
			if i*2 > 9 {
				numeros[x] = (i * 2) - 9
			} else {
				numeros[x] = i * 2
			}
			suma += numeros[x]
			bandera = !bandera
		}
	}
	if suma%10 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Valid("234 567 891 234"))
}
