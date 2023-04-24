package main

import (
	"errors"
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	var error error
	var x = uint64(0)
	if number <= 0 || number > 64 {
		error = errors.New("error")
	} else {
		x = uint64(math.Pow(2, float64(number-1)))
	}
	return x, error
}

func Total() uint64 {
	var x = uint64(1)
	var cont = 0
	var suma = uint64(0)
	for cont < 64 {
		x = uint64(math.Pow(2, float64(cont)))
		suma += x
		cont++
	}
	return suma
}

func main() {
	fmt.Println(Total())
}
