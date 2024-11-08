package main

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	account := BankAccount{Owner: "Назар", Balance: 1000.0}

	// Внесемо кошти
	account.Deposit(500.0)

	// Перевіримо, чи баланс збільшився на 500
	if account.Balance != 1500.0 {
		t.Errorf("expected balance 1500.0, got %.2f", account.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	account := BankAccount{Owner: "Олексій", Balance: 1000.0}

	// Знімаємо 200
	account.Withdraw(200.0)
	if account.Balance != 800.0 {
		t.Errorf("expected balance 800.0, got %.2f", account.Balance)
	}

	// Спробуємо зняти більше, ніж є на рахунку
	account.Withdraw(1000.0)
	if account.Balance != 800.0 {
		t.Errorf("expected balance 800.0, got %.2f", account.Balance)
	}
}

func TestCheckBalance(t *testing.T) {
	account := BankAccount{Owner: "Назар", Balance: 1000.0}

	// Перевіримо початковий баланс
	account.CheckBalance()
	if account.Balance != 1000.0 {
		t.Errorf("expected balance 1000.0, got %.2f", account.Balance)
	}

	// Внесемо кошти та перевіримо новий баланс
	account.Deposit(500.0)
	account.CheckBalance()
	if account.Balance != 1500.0 {
		t.Errorf("expected balance 1500.0, got %.2f", account.Balance)
	}

	// Знімемо кошти та перевіримо новий баланс
	account.Withdraw(400.0)
	account.CheckBalance()
	if account.Balance != 1100.0 {
		t.Errorf("expected balance 1100.0, got %.2f", account.Balance)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	account := BankAccount{Owner: "Олексій", Balance: 100.0}

	// Спробуємо зняти більше, ніж є на рахунку
	account.Withdraw(200.0)

	// Перевіримо, чи баланс залишився незмінним
	if account.Balance != 100.0 {
		t.Errorf("expected balance 100.0, got %.2f", account.Balance)
	}
}
