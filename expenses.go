package main

import "fmt"

var categoryTotals = make(map[string]int)

func addExpense(amount int, category string) {
	if amount <= 0 {
		fmt.Println("Invalid Amount!")
		return
	}

	if len(categoryTotals) >= MaxExpenses {
		fmt.Println("Max Expense Limit Reached!")
		return
	}

	categoryTotals[category] += amount
}

func listExpenses() {
	for category, amount := range categoryTotals {
		fmt.Println("Category:", category, "| Amount:", amount)
	}
}

func calculateTotal() int {
	total := 0

	for _, amount := range categoryTotals {
		total += amount
	}

	return total
}

func highestExpense() int {
	if len(categoryTotals) == 0 {
		return 0
	}

	maxExpense := 0

	for _, amount := range categoryTotals {
		if amount > maxExpense {
			maxExpense = amount
		}
	}

	return maxExpense
}