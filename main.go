package main

func main() {
	addExpense(500, "Food")
	addExpense(1000, "Travel")
	addExpense(250, "Food")
	addExpense(-100, "Shopping")

	showReport()
}