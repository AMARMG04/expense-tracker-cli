package main

import (
	"fmt"
)

type Expense struct {
	ID          int
	Amount      int
	Category    string
	Description string
}

func (e Expense) Print() {
	fmt.Println("ID: ", e.ID)
	fmt.Println("Amount: ", e.Amount)
	fmt.Println("Category: ", e.Category)
	fmt.Println("Description: ", e.Description)
	fmt.Println()
}

var expenses []Expense

var id int = 1

func addExpense(amount int, category, description string) {
	if amount <= 0 {
		fmt.Println("Invalid Amount!")
		return
	}

	if len(expenses) < MaxExpenses {
		expenses = append(expenses, Expense{
			id,
			amount,
			category,
			description,
		})
		id++
	} else {
		fmt.Println("Expenses list is full!")
	}
}

func listExpenses() {
	for _, expense := range expenses {
		expense.Print()
	}
}

func calculateTotal() int {
	total := 0

	for _, expense := range expenses {
		total += expense.Amount
	}

	return total
}

func categoryTotal() map[string]int {
	var categoryTotalMap = make(map[string]int)

	for _, expense := range expenses {
		categoryTotalMap[expense.Category] += expense.Amount
	}
	return categoryTotalMap
}

func highestExpense() int {
	if len(expenses) == 0 {
		return 0
	}

	maxExpense := 0

	for _, expense := range expenses {
		if expense.Amount > maxExpense {
			maxExpense = expense.Amount
		}
	}

	return maxExpense
}
