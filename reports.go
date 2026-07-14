package main

import "fmt"

func showReport() {
	fmt.Println("=====", AppName, "=====")

	fmt.Println("\nCategory Totals:")
	listExpenses()

	fmt.Println("\nTotal Expenses:", calculateTotal())
	fmt.Println("Highest Category Expense:", highestExpense())
}