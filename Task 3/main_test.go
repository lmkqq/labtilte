package main

import (
	"bytes"
	"fmt"
	"testing"
)

// Функція для перевірки календаря для конкретного місяця та року
func TestPrintCalendar(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		expected string
	}{
		{
			year: 2024, month: 2,
			expected: `Пн Вт Ср Чт Пт Сб Нд
 1  2  3 
 4  5  6  7  8  9 10 
11 12 13 14 15 16 17 
18 19 20 21 22 23 24 
25 26 27 28 
`,
		},
		{
			year: 2023, month: 5,
			expected: `Пн Вт Ср Чт Пт Сб Нд
 1  2  3  4  5  6  7 
 8  9 10 11 12 13 14 
15 16 17 18 19 20 21 
22 23 24 25 26 27 28 
29 30 31 
`,
		},
		{
			year: 2024, month: 12,
			expected: `Пн Вт Ср Чт Пт Сб Нд
 1  2 
 3  4  5  6  7  8  9 
10 11 12 13 14 15 16 
17 18 19 20 21 22 23 
24 25 26 27 28 29 30 
31 
`,
		},
	}

	for _, tt := range tests {
		t.Run("Testing calendar", func(t *testing.T) {
			// Перехоплюємо вивід в буфер
			var buf bytes.Buffer
			// Заміна стандартного виводу на буфер
			fmt.Printf = func(format string, args ...interface{}) (n int, err error) {
				return fmt.Fprintf(&buf, format, args...)
			}
			// Викликаємо метод для друку календаря
			printCalendar(tt.year, tt.month)

			// Перевірка, чи вивід співпадає з очікуваним
			got := buf.String()
			if got != tt.expected {
				t.Errorf("for year %d and month %d, expected:\n%s\nbut got:\n%s", tt.year, tt.month, tt.expected, got)
			}
		})
	}
}
