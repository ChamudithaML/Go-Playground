package main

import (
	"fmt"
	"sync"
	"time"
)

// BankAccount represents a shared account
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// Deposit allows adding money to the account
func (acc *BankAccount) Deposit(amount int, name string) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.balance < 1000 {
		acc.balance += amount
		fmt.Printf("%s deposited %d. Balance: %d\n", name, amount, acc.balance)
	} else {
		fmt.Printf("%s tried to deposit but balance is >= 1000. Balance: %d\n", name, acc.balance)
	}
}

// Withdraw allows removing money from the account
func (acc *BankAccount) Withdraw(amount int, name string) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.balance > 500 {
		acc.balance -= amount
		fmt.Printf("%s withdrew %d. Balance: %d\n", name, amount, acc.balance)
	} else {
		fmt.Printf("%s tried to withdraw but balance <= 500. Balance: %d\n", name, acc.balance)
	}
}

func wife(account *BankAccount, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		account.Deposit(100, "Wife")
		time.Sleep(100 * time.Millisecond) // Simulate time taken to deposit
	}
}

func husband(account *BankAccount, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		account.Withdraw(100, "Husband")
		time.Sleep(150 * time.Millisecond) // Simulate time taken to withdraw
	}
}

func example2_main() {
	account := &BankAccount{balance: 500} // Initial balance
	var wg sync.WaitGroup

	wg.Add(2) // Two goroutines: wife and husband

	go wife(account, 10, &wg)
	go husband(account, 10, &wg)

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("Final balance:", account.balance)
}
