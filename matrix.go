package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	filas    [][]int
	columnas [][]int
}

func New(s string) (Matrix, error) {
	s = strings.Trim(s, " ")
	matriz := Matrix{}
	var err error
	if len(s) > 0 {
		filtro1 := strings.SplitAfterN(s, "\n", strings.Count(s, "\n")+1)
		var filas = [][]int{}

		for _, f := range filtro1 {
			sub := strings.Split(f, " ")
			subFila := []int{}
			for _, v := range sub {
				if v != "\n" && v != " " && v != "" {
					v = strings.Trim(v, " ")
					v = strings.Trim(v, "\n")
					numero, err1 := strconv.Atoi(v)
					if err1 != nil {
						err = err1
						return matriz, err
					}
					subFila = append(subFila, numero)
				}

			}
			filas = append(filas, subFila)
		}

		numColumas := len(filas[0]) - 1
		var numFila = 0
		var totalfilas = len(filas) - 1
		var bandera = false
		var columnas = [][]int{}
		var cont = 0
		for bandera == false {
			for cont <= numColumas {
				var columna = []int{}
				for numFila <= totalfilas {
					fil := filas[numFila]                //[1,2]
					columna = append(columna, fil[cont]) //1
					numFila++
				}
				columnas = append(columnas, columna)
				cont++
				numFila = 0
			}
			bandera = true
		}

		matriz.filas = filas
		matriz.columnas = columnas
	}

	return matriz, err
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	return m.columnas
}

func (m Matrix) Rows() [][]int {
	return m.filas
}

func (m Matrix) Set(row, col, val int) bool {
	if len(m.filas) >= row && len(m.columnas) >= col {
		m.filas[row-1][col-1] = val
		m.columnas[col-1][row-1] = val
		return true
	}
	return false
}

func main() {
	//New("   ")
	m, _ := New("1 2\n10 20 30")
	fmt.Println(m)
	//fmt.Println(m.Set(1, 1, 50))
	//New("1 2 3\n4 5 6\n7 8 9\n 8 7 6")
	//New("")

}
