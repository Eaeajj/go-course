package main

import "fmt"

type ExpensesMap = map[string]float64
type ExpensesTracker struct {
	expenses ExpensesMap
}

func NewExpensesTracker() *ExpensesTracker {
	return &ExpensesTracker{
		expenses: make(ExpensesMap),
	}
}

func (et *ExpensesTracker) AddExpense(name string, sum float64) {
	et.expenses[name] += sum
}

func (et *ExpensesTracker) GetTotalExpenses() float64 {
	total := 0.0
	for _, amount := range et.expenses {
		total += amount
	}
	return total
}

func (et *ExpensesTracker) PrintTotalExpenses() {
	total := et.GetTotalExpenses()
	fmt.Printf("Total expenses are: %f", total)
}
