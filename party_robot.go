package main

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	n := fmt.Sprintf("%v", age)
	return "Happy birthday " + name + "! You are now " + n + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var numero = "0"
	if table < 10 {
		numero = "00" + fmt.Sprintf("%v", table)
	} else if table < 100 {
		numero = "0" + fmt.Sprintf("%v", table)
	} else {
		numero = fmt.Sprintf("%v", table)
	}
	n := fmt.Sprintf("%.1f", distance)
	x := Welcome(name) + "\n" + "You have been assigned to table " + numero + ". Your table is " + direction + ", exactly " + n + " meters from here." + "\n" +
		"You will be sitting next to " + neighbor + "."
	return x
}
