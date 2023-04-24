package main

import (
	"fmt"
	"sync"
)

// Define the Account type here.
var wg sync.WaitGroup

type Account struct {
	balance int64
	activa  bool
	//lectura y escritura
	acceso sync.RWMutex
}

func Open(amount int64) *Account {
	var cuenta Account
	if amount < 0 {
		return nil
	}
	cuenta = Account{
		balance: amount,
		activa:  true,
	}

	return &cuenta
}

func (a *Account) Balance() (int64, bool) {
	a.acceso.RLock()
	if a.activa == false {
		return a.balance, false
	}
	wg.Done()
	a.acceso.RUnlock()
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.acceso.Lock()
	defer wg.Done()
	if a.activa == false {
		return a.balance, false
	}
	saldo := a.balance + amount
	if amount < 0 {
		return saldo, false
	} else {
		a.balance = saldo
	}
	a.acceso.Unlock()
	return a.balance, true

}

func (a *Account) Close() (int64, bool) {
	a.acceso.Lock()
	if a.activa == false {
		return 0, false
	}
	a.activa = false
	a.balance = 0
	wg.Done()
	a.acceso.Unlock()
	return a.balance, true
}

func main() {
	const amt = 10
	cuenta := Open(amt)
	wg.Add(3)

	//recuperar return de una rutina
	/*var x int64
	go func() {
		x, _ = cuenta.Balance()
	*/

	go cuenta.Deposit(20)
	go cuenta.Close()
	go cuenta.Deposit(40) //ya no deja porque está cerrada
	wg.Wait()
	fmt.Println("Saldo", cuenta.balance)
}
