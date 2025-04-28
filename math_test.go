package main

import "testing"

func TestAdd(t *testing.T) {
	// Тестовые случаи
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"положительные числа", 2, 3, 5},
		{"отрицательные числа", -1, -1, -2},
		{"ноль", 0, 5, 5},
	}

	// Проверяем каждый тестовый случай
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; ожидалось %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
