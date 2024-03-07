package daytwo

import (
	"errors"
	"fmt"
	"sync"
)

func Bank() {
	var wa sync.WaitGroup
	var bank IBankActions
	bank = &BankAccount{balance: 500}
	wa.Add(3)
	go func() {
		remainingBalance, err := bank.deposit(20.0)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
		}
		fmt.Println("Deposit Success, remaining balance:", remainingBalance)
		wa.Done()
	}()
	go func() {

		remainingBalance, err := bank.withdraw(700.3)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
		}
		fmt.Println("Withdraw Success, remaining balance:", remainingBalance)
		wa.Done()
	}()
	go func() {
		remainingBalance, err := bank.withdraw(700.3)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
		}
		fmt.Println("Withdraw Success, remaining balance:", remainingBalance)
		wa.Done()
	}()
	wa.Wait()

}

type IBankActions interface {
	withdraw(amt float64) (float64, error)
	deposit(amt float64) (float64, error)
}

type BankAccount struct {
	balance float64
	mux     sync.Mutex
}

func (b *BankAccount) withdraw(amt float64) (float64, error) {
	if amt <= 0 {
		return b.balance, errors.New("withdrawl amount needs to be greater than 0")
	}
	result := b.balance - amt
	if amt < 0 {
		return b.balance, errors.New("insufficient balance")
	} else {
		b.mux.Lock()
		defer b.mux.Unlock()
		b.balance -= amt
	}
	return result, nil
}
func (b *BankAccount) deposit(amt float64) (float64, error) {
	if amt <= 0 {
		return b.balance, errors.New("deposit amount needs to be greater than 0")
	}
	b.mux.Lock()
	defer b.mux.Unlock()
	b.balance += amt
	return b.balance, nil
}
