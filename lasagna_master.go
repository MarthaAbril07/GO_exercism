package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(capas []string, x int) int {
	if x == 0 {
		return len(capas) * 2
	} else {
		return len(capas) * x
	}
}

// TODO: define the 'Quantities()' function
func Quantities(capas []string) (int, float64) {
	x := 0
	s := 0.0
	for _, value := range capas {
		if value == "noodles" {
			x++
		}
		if value == "sauce" {
			s++
		}
	}
	return (x * 50), (s * 0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	for index, value := range myList {
		if value == "?" {
			myList[index] = friendsList[len(friendsList)-1]
		}
	}
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(cantidades []float64, porciones int) []float64 {
	nuevas := []float64{}
	for _, x := range cantidades {
		nuevas = append(nuevas, (x/2)*float64(porciones))
	}
	return nuevas
}

func main() {
	m := []string{"sauce", "noodles", "béchamel", "meat", "mozzarella", "noodles", "ricotta", "eggplant", "béchamel", "noodles", "sauce", "mozzarella"}
	x := PreparationTime(m, 1)

	fmt.Println(x)

	n := []float64{0.5, 250, 150, 3, 0.5}
	fmt.Println(ScaleRecipe(n, 6))
}
