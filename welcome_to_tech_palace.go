package main

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var cadena = ""
	for x := 0; x < numStarsPerLine; x++ {
		cadena += "*"
	}
	x := cadena + "\n" + welcomeMsg + "\n" + cadena
	return x
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	x := strings.ReplaceAll(oldMsg, "*", "")
	x = strings.TrimSpace(x)
	return x
}
