package main

func main() {
	addExpense(100, "Food", "Spent in KFC")
	addExpense(100, "Travel", "Went to Majestic")
	addExpense(100, "Food", "Spent in McDonalds")
	addExpense(100, "Subscriptions", "")

	updateAmount(1, 250)

	updateDescription(2, "Bus to Majestic")

	updateCategory(4, "Entertainment")

	expenses[0].MarkReimbursed()

	deleteExpense(3)

	showReport()
}
