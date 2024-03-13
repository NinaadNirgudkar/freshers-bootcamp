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
		defer wa.Done()
		remainingBalance, err := bank.deposit(20.0)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
			return
		}
		fmt.Println("Deposit Success, remaining balance:", remainingBalance)
	}()
	go func() {
		defer wa.Done()
		remainingBalance, err := bank.withdraw(700.3)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
			return
		}
		fmt.Println("Withdraw Success, remaining balance:", remainingBalance)
	}()
	go func() {
		defer wa.Done()
		remainingBalance, err := bank.withdraw(70.3)
		fmt.Println(remainingBalance, err)
		if err != nil {
			fmt.Println(err.Error(), "Remaining balance:", remainingBalance)
			return
		}
		fmt.Println("Withdraw Success, remaining balance:", remainingBalance)
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
	b.mux.Lock()
	defer b.mux.Unlock()
	if b.balance-amt <= 0 {
		return b.balance, errors.New("Withdrawl failed.")
	}
	b.balance -= amt
	return b.balance, nil
}
func (b *BankAccount) deposit(amt float64) (float64, error) {
	b.mux.Lock()
	defer b.mux.Unlock()
	if amt <= 0 {
		return b.balance, errors.New("deposit amount needs to be greater than 0")
	}
	b.balance += amt
	return b.balance, nil
}
