package main

import "fmt"

func main() {
	library := NewLibrary()

	library.AddBook("7 Habits of highly effective people", "Steven R. Covey", 2023)
	library.AddBook("Государство", "Платон", 2000)
	library.AddBook("Наедине с собой. Размышления", "Марк Аврелий", 2022)

	library.ListAllBooks()

	book := library.FindBook("Государство")
	fmt.Println("\n", "Найденная книга", book)

	library.IssueBook("Государство")
	library.IssueBook("Государство")
	library.ReturnBook("Государство")
}
