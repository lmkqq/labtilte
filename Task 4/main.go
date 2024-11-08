package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
	fmt.Printf("Внесено %.2f. Новий баланс: %.2f\n", amount, acc.Balance)
}

func (acc *BankAccount) Withdraw(amount float64) {
	if amount <= acc.Balance {
		acc.Balance -= amount
		fmt.Printf("Знято %.2f. Новий баланс: %.2f\n", amount, acc.Balance)
	} else {
		fmt.Println("Недостатньо коштів на рахунку.")
	}
}

func (acc *BankAccount) CheckBalance() {
	fmt.Printf("Поточний баланс рахунку: %.2f\n", acc.Balance)
}

func main() {
	account := BankAccount{Owner: "Назар", Balance: 1000.0}

	var choice int
	var amount float64

	for {
		fmt.Println("\nМеню операцій:")
		fmt.Println("1. Внести кошти")
		fmt.Println("2. Зняти кошти")
		fmt.Println("3. Перевірити баланс")
		fmt.Println("4. Вийти")
		fmt.Print("Оберіть операцію (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Введіть суму для внесення: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Введіть суму для зняття: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.CheckBalance()
		case 4:
			fmt.Println("Вихід з програми.")
			return // вихід з функції main, завершення програми
		default:
			fmt.Println("Невірний вибір. Спробуйте ще раз.")
		}
	}
}
