package main

import (
	"fmt"
	"testing"
)

func TestCountDivisors(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},  // 1 має тільки один дільник: 1
		{input: 6, expected: 4},  // 6 має дільники: 1, 2, 3, 6
		{input: 10, expected: 4}, // 10 має дільники: 1, 2, 5, 10
		{input: 15, expected: 4}, // 15 має дільники: 1, 3, 5, 15
		{input: 28, expected: 6}, // 28 має дільники: 1, 2, 4, 7, 14, 28
		{input: 13, expected: 2}, // 13 є простим числом, має тільки дільники: 1, 13
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%d", test.input), func(t *testing.T) {
			result := countDivisors(test.input)
			if result != test.expected {
				t.Errorf("expected %d divisors for %d, got %d", test.expected, test.input, result)
			}
		})
	}
}
