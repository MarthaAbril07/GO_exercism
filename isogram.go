package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	palabra := strings.ToUpper(word)
	var valores = map[string]int32{}
	var bandera = true
	for _, j := range palabra {
		if j == 45 || j == 32 {
			continue
		}
		if _, ok := valores[string(j)]; ok {
			bandera = false
			break
		} else {
			valores[string(j)] = j
		}
	}
	return bandera
}

func main() {
	fmt.Println(IsIsogram("hola"))
}
