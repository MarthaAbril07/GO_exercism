package main

import "fmt"

func SquareOfSum(n int) int {
	var cont = 1
	var suma = 0
	for cont <= n {
		suma += cont
		cont++
	}
	return suma * suma
}

func SumOfSquares(n int) int {
	var cont = 1
	var suma = 0
	for cont <= n {
		suma += (cont * cont)
		cont++
	}
	return suma
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	fmt.Println(SquareOfSum(10))
	fmt.Println(SumOfSquares(10))
	fmt.Println(Difference(10))

}
