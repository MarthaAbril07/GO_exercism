package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	} else if balance >= 0 && balance < 1000 {
		return 0.5
	} else if balance >= 1000 && balance < 5000 {
		return 1.621
	} else {
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return (balance * float64(InterestRate(balance))) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var cont = 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		cont++
	}
	return cont
}

func main() {
	fmt.Println(InterestRate(200.75))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))
	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))

}
