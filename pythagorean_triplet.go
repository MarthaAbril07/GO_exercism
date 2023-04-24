package main

import (
	"math"
)

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	tripleta := []Triplet{}
	minimo := float64(min)
	maximo := float64(max)
	a := minimo
	b := minimo
	//var bandera = false
	for b < maximo {
		suma := (a * a) + (b * b)
		c := math.Ceil(math.Sqrt(suma)*100) / 100
		cCuadrada := c * c
		//fmt.Println(a, " + ", b, " = ", suma, " -> ", cCuadrada, " - ", suma == cCuadrada)
		if suma == cCuadrada {
			if a < b && b < c {
				tripleta = append(tripleta, Triplet{int(a), int(b), int(c)})
			}
		}
		if a < maximo {
			a++
		} else {
			a = minimo
			b++
		}

	}

	return tripleta
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	min := int(math.Sqrt(float64(p)))
	max := p
	tripleta := Range(min, max)
	h1 := []Triplet{}
	resultados := []Triplet{}
	for _, item := range tripleta {
		a := item[0]
		b := item[1]
		c := item[2]
		if a+b+c == p {
			h1 = append(h1, Triplet{a, b, c})
		}
	}

	//
	for x := len(h1) - 1; x >= 0; x-- {
		resultados = append(resultados, h1[x])
	}
	//fmt.Println("Resultados", resultados)
	return resultados
}

func main() {
	//Range(11, 20)
	Sum(180)
}
