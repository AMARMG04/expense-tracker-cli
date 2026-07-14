package main

import "fmt"

func showReport() {
	fmt.Println("===== ", AppName, "" =====")
	fmt.Println("Expenses:")
	listExpenses()

	fmt.Println("Total: ", calculateTotal())
	fmt.Println("MaxExpense: ", highestExpense())
}