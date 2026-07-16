package main

import "fmt"

func showReport() {
	fmt.Println("=====", AppName, "=====")

	fmt.Println("Listing Expenses:")
	listExpenses()

	fmt.Println("Category Totals:")
	fmt.Println(categoryTotal())

	fmt.Println("\nTotal Expenses:", calculateTotal())
	fmt.Println("Highest Category Expense:", highestExpense())
}