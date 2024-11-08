package main

import (
	"testing"
)

func TestTotalValue(t *testing.T) {
	tests := []struct {
		product  Product
		expected float64
	}{
		{Product{Name: "Смартфон", Price: 15000.0, Quantity: 5}, 75000.0},
		{Product{Name: "Ноутбук", Price: 30000.0, Quantity: 3}, 90000.0},
		{Product{Name: "Мишка", Price: 500.0, Quantity: 10}, 5000.0},
		{Product{Name: "Клавіатура", Price: 1200.0, Quantity: 7}, 8400.0},
	}

	for _, test := range tests {
		t.Run(test.product.Name, func(t *testing.T) {
			result := test.product.TotalValue()
			if result != test.expected {
				t.Errorf("Для продукту %s очікувана вартість %.2f, але отримано %.2f", test.product.Name, test.expected, result)
			}
		})
	}
}

func TestTotalInventoryValue(t *testing.T) {
	// Створюємо масив продуктів
	products := []Product{
		{Name: "Смартфон", Price: 15000.0, Quantity: 5},
		{Name: "Ноутбук", Price: 30000.0, Quantity: 3},
		{Name: "Мишка", Price: 500.0, Quantity: 10},
		{Name: "Клавіатура", Price: 1200.0, Quantity: 7},
	}

	var totalInventoryValue float64
	// Цикл для підрахунку загальної вартості
	for _, product := range products {
		totalInventoryValue += product.TotalValue()
	}

	expectedTotalValue := 75000.0 + 90000.0 + 5000.0 + 8400.0

	// Перевіряємо, чи загальна вартість правильна
	if totalInventoryValue != expectedTotalValue {
		t.Errorf("Очікувана загальна вартість %.2f, але отримано %.2f", expectedTotalValue, totalInventoryValue)
	}
}
