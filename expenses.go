package main

import (
	"fmt"
)

type Expense struct {
	ID          int
	Amount      int
	Category    string
	Description string
	Reimbursed  bool
}

func (e Expense) Print() {
	fmt.Println("ID: ", e.ID)
	fmt.Println("Amount: ", e.Amount)
	fmt.Println("Category: ", e.Category)
	fmt.Println("Description: ", e.Description)
	fmt.Println("Reimbursed: ", e.Reimbursed)
	fmt.Println()
}

func (e Expense) IsValid() bool {
	if e.Amount > 0 && e.Category != "" {
		return true
	}

	return false
}

func (e *Expense) UpdateAmount(newAmount int) {
	e.Amount = newAmount
}

func (e *Expense) UpdateDescription(newDescription string) {
	e.Description = newDescription
}

func (e *Expense) UpdateCategory(newCategory string) {
	e.Category = newCategory
}

func (e *Expense) MarkReimbursed() {
	e.Reimbursed = !e.Reimbursed
}

var expenses []Expense

var id int = 1

func addExpense(amount int, category, description string) {
	expense := Expense{
		ID:          id,
		Amount:      amount,
		Category:    category,
		Description: description,
		Reimbursed:  false,
	}

	if expense.IsValid() && len(expenses) < MaxExpenses {
		expenses = append(expenses, expense)
		id++
	} else {
		fmt.Println("Expenses list is full!")
	}
}

func updateAmount(id int, amount int) bool {
	isUpdated := false
	for i := range expenses {
		if expenses[i].ID == id {
			expenses[i].UpdateAmount(amount)
			isUpdated = true
			break
		}
	}

	if isUpdated {
		fmt.Println("Updated the amount successfully")
	} else {
		fmt.Println("Expense not found!")
	}

	return isUpdated
}

func updateCategory(id int, category string) bool {
	isUpdated := false
	for i := range expenses {
		if expenses[i].ID == id {
			expenses[i].UpdateCategory(category)
			isUpdated = true
			break
		}
	}

	if isUpdated {
		fmt.Println("Updated the category successfully")
	} else {
		fmt.Println("Expense not found!")
	}

	return isUpdated
}

func updateDescription(id int, description string) bool {
	isUpdated := false
	for i := range expenses {
		if expenses[i].ID == id {
			expenses[i].UpdateDescription(description)
			isUpdated = true
			break
		}
	}

	if isUpdated {
		fmt.Println("Updated the description successfully")
	} else {
		fmt.Println("Expense not found!")
	}

	return isUpdated
}

func deleteExpense(id int) {
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(
				expenses[:i],
				expenses[i+1:]...,
			)

			fmt.Println("Expense deleted")
			return
		}
	}

	fmt.Println("Expense not found")
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
