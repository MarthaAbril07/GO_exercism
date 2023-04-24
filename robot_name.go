package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	Nombre string
}

func (r *Robot) Name() (string, error) {
	letra1 := string(rune(getLetra()))
	letra2 := string(rune(getLetra()))
	numero := getNumero()
	s := letra1 + letra2 + fmt.Sprint(numero)

	if r.Nombre == s {
		err := errors.New("error")
		return "", err
	}
	r.Nombre = s
	return s, nil
}

func getLetra() int {
	aleatorio := rand.Intn(90-65) + 65
	return aleatorio
}

func getNumero() int {
	aleatorio := rand.Intn(999-100) + 100
	return aleatorio
}

func (r *Robot) Reset() {
	r.Nombre = ""
}

func main() {
	r := Robot{}

	r.Name()
	fmt.Println("antes:", r.Nombre)
	r.Reset()
	fmt.Println("Despues", r.Nombre)
}
