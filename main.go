// https://hackmd.io/@enrinal/ByHXRMEMkg#Example-Unit-Test-in-bank_testgo

package main

import (
    "fmt"
    "bank-account-management/bank" // FS9966251 ~ Mengimpor paket bank dari modul bank-account-management
)

func main() {
    account := bank.Account{} // FS9966251 ~ Membuat objek account dari tipe struct Account

    // Menyimpan uang
    if err := account.Deposit(100.50); err != nil {
        fmt.Println("Error:", err) // FS9966251 ~ Menangani error jika deposit gagal
    }

    // Mencoba menarik uang
    if err := account.Withdraw(50); err != nil {
        fmt.Println("Error:", err) // FS9966251 ~ Menangani error jika penarikan gagal
    }

    // Mencoba menarik uang lebih dari saldo
    if err := account.Withdraw(100); err != nil {
        fmt.Println("Error:", err) // FS9966251 ~ Menangani error jika penarikan melebihi saldo
    }

    fmt.Println("Saldo Akhir:", account.GetBalance()) // FS9966251 ~ Menampilkan saldo akhir
}