package main

import "fmt"

func showReport() {
	fmt.Println("=====", AppName, "=====")

	fmt.Println("\nListing Expenses:")
	listExpenses()

	fmt.Println("\nCategory Totals:")

	for category, total := range categoryTotal() {
		fmt.Printf("%s: %d\n", category, total)
	}

	fmt.Println("\nTotal Expenses:", calculateTotal())
	fmt.Println("Highest Expense:", highestExpense())
}
