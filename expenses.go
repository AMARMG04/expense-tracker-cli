package main
import "fmt"

var expenses []int

func addExpense(amount int) {
	if amount <= 0{
		fmt.Println("Invalid Amount!")
	}else if len(expenses) <= MaxExpenses{
		expenses = append(expenses, amount)
	}else{
		fmt.Println("Max Expense Limit Reached!")
	}
}

func removeExpense(index int) {
	if index > 0 && index < len(expenses){
		fmt.Println("Removed", expenses[index])
		expenses = append(expenses[:index], expenses[index+1:]...,)
	}else{
		fmt.Println("Index out of bound")
	}
}

func listExpenses() {
	for i, expense := range expenses{
		fmt.Println("Expense ", i+1, " : ", expense)
	}
}

func calculateTotal() int{
	var total int
	for _, expense := range expenses{
		total += expense
	}

	return total
}

func highestExpense() int{
	var maxExpense int = -1
	for _, expense := range expenses{
		if maxExpense < expense{
			maxExpense = expense
		}
	}

	return maxExpense
}
