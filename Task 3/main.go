package main

import (
	"fmt"
	"time"
)

func printCalendar(year, month int) {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(firstDay.Weekday())
	daysInMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	fmt.Println("Пн Вт Ср Чт Пт Сб Нд")
	for i := 0; i < dayOfWeek; i++ {
		fmt.Print("   ")
	}
	for day := 1; day <= daysInMonth; day++ {
		fmt.Printf("%2d ", day)
		if (day+dayOfWeek)%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func main() {
	var year, month int
	fmt.Print("Введіть рік: ")
	fmt.Scan(&year)
	fmt.Print("Введіть місяць (1-12): ")
	fmt.Scan(&month)
	printCalendar(year, month)
}
