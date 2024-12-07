package main

func main() {

	et := NewExpensesTracker()

	et.AddExpense("entertainments", 60)
	et.AddExpense("entertainments", 60)
	et.AddExpense("goods", 60)

	et.PrintTotalExpenses()

}
