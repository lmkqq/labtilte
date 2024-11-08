package main

import (
	"fmt"
)

type Product struct {
    Name     string
    Price    float64
    Quantity int
}

// Функція для обчислення вартості конкретного продукту
func (p Product) TotalValue() float64 {
    return p.Price * float64(p.Quantity)
}

func main() {
    // Створюємо масив продуктів
    products := []Product{
        {Name: "Смартфон", Price: 15000.0, Quantity: 5},
        {Name: "Ноутбук", Price: 30000.0, Quantity: 3},
        {Name: "Мишка", Price: 500.0, Quantity: 10},
        {Name: "Клавіатура", Price: 1200.0, Quantity: 7},
    }

    var totalInventoryValue float64

    fmt.Println("Перелік продуктів та їх вартість:")
    
    // Цикл для перебору всіх продуктів і виведення їх вартості
    for _, product := range products {
        fmt.Printf("Продукт: %s, Ціна: %.2f, Кількість: %d, Вартість: %.2f\n", 
            product.Name, product.Price, product.Quantity, product.TotalValue())
        totalInventoryValue += product.TotalValue()
    }

    // Виведення загальної вартості всіх товарів
    fmt.Printf("\nЗагальна вартість товарно-матеріальних цінностей: %.2f\n", totalInventoryValue)
}
