package main

type BookStatus string

const (
	BookStatusAvailable BookStatus = "доступна"
	BookStatusIssued    BookStatus = "на руках у читателя"
)

type Book struct {
	Id     int
	Name   string
	Author string
	Year   int
	Status BookStatus
}

type Library struct {
	Books []Book
}
