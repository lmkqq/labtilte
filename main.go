package main

import (
	"fmt"
)

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	var number int
	fmt.Print("Введіть число: ")
	fmt.Scan(&number)
	fmt.Printf("Кількість дільників числа %d: %d\n", number, countDivisors(number))
}
