// Package weather can forecast the current weather condition of various //cities in Goblinocus.
package main

//CurrentCondition represents the actual condition.
var CurrentCondition string

//CurrentLocation represents the actual position.
var CurrentLocation string

//Forecast returns the  weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
