package main

import (
	"fmt"
	"strings"
)

func Encode(pt string) string {
	pt = strings.ToLower(pt)
	for _, x := range pt {
		if x < 48 || x > 57 && x < 97 || x > 122 {
			pt = strings.ReplaceAll(pt, string(x), "")
		}
	}
	if len(pt) > 0 {
		fil, col := encontrarMultiplicacion(len(pt))
		var cadena = ""
		var cadenaCompleta = ""
		var matriz = []string{}
		for index, value := range pt {
			if (index+1)%col == 0 {
				cadena += string(value)
				cadenaCompleta += cadena + " "
				//cadena += string(value) + " "
				matriz = append(matriz, cadena)
				cadena = ""
			} else {
				cadena += string(value)
				//cadenaCompleta += string(value)
			}
		}
		//fmt.Println(cadenaCompleta)
		total := len(pt) + col
		if len(cadenaCompleta) <= total {
			for len(cadenaCompleta) < total {
				cadena += " "
				cadenaCompleta += " "
			}
			matriz = append(matriz, cadena)
		}
		return voltear(matriz, fil, col)
	}
	return ""
}

func voltear(matriz []string, totalFilas, totalColumnas int) string {
	//7  8
	var bandera = false
	var nuevaPalabra = ""
	var numFila = 0
	var numColumna = 0
	var volteado = ""
	for bandera == false {
		for numFila < totalFilas {
			palabra := matriz[numFila]
			nuevaPalabra += string(palabra[numColumna])
			numFila++
		}

		numColumna++
		if numFila == totalFilas && numColumna == totalColumnas {
			volteado += nuevaPalabra
		} else {
			volteado += nuevaPalabra + " "
		}
		numColumna--

		nuevaPalabra = ""
		numFila = 0
		numColumna++

		if numColumna == totalColumnas {
			bandera = true
		}

	}
	return volteado
}

func encontrarMultiplicacion(longitud int) (fila int, col int) {
	var c = 1
	var r = 1
	for r <= longitud {
		//fmt.Println("Entra: ", r, " * ", c)
		if r*c < longitud {
			if c == longitud {
				c = 1
				r++
			} else {
				c++
			}
		} else {
			if c >= r {
				if (c - r) <= 1 {
					return r, c
				} else {
					r++
					c = 1
				}
			} else {
				if c == longitud {
					c = 1
				} else {
					c++
				}
			}
		}
	}
	return r, c
}

func main() {
	/*x := Encode("If man was meant to stay on the ground, god would have given us roots.")
	fmt.Println(len(x), " vs ", len("imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau "))*/
	//x := Encode("Time is an illusion. Lunchtime doubly so.")
	b := Encode("")
	fmt.Println(len(b), " vs ", len(""))

}
