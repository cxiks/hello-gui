package main

import "fmt"

// Greeting возвращает приветствие для указанного имени.
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// SumRange считает сумму чисел от from до to включительно.
func SumRange(from, to int) int {
	sum := 0
	for i := from; i <= to; i++ {
		sum += i
	}
	return sum
}
