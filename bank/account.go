package bank

import (
    "errors"
)

// Account struct memiliki field Balance yang menyimpan saldo.
type Account struct {
    Balance float64
}

// Deposit menambahkan sejumlah uang ke saldo akun. Jika jumlahnya tidak valid, akan mengembalikan error.
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    a.Balance += amount
    return nil
}

// Withdraw mengurangi saldo akun dengan jumlah yang ditentukan. Jika jumlahnya lebih besar dari saldo, mengembalikan error.
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdraw amount must be positive")
    }
    if amount > a.Balance {
        return errors.New("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

// GetBalance mengembalikan saldo akun saat ini.
func (a *Account) GetBalance() float64 {
    return a.Balance
}
