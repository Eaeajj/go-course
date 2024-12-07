package main

import (
	"fmt"
	"math/rand"
)

func NewLibrary() *Library {
	return &Library{}
}

func (l *Library) AddBook(name, author string, year int) {
	id := rand.Intn(1000) + 1
	l.Books = append(l.Books, Book{
		Id:     id,
		Name:   name,
		Author: author,
		Year:   year,
		Status: BookStatusAvailable,
	})
}

func (l *Library) ListAllBooks() {
	fmt.Println("Список всех книг:")
	for _, book := range l.Books {
		fmt.Printf("%+v\n", book)
	}
}

func (l *Library) FindBook(name string) *Book {
	for _, book := range l.Books {
		if book.Name == name {
			return &book
		}
	}
	return nil
}

func (l *Library) IssueBook(name string) {
	for i, book := range l.Books {
		if book.Name == name && book.Status == BookStatusAvailable {
			l.Books[i].Status = BookStatusIssued
			fmt.Printf("Книга '%s' выдана\n", name)
			return
		}
	}
	fmt.Printf("Книга '%s' не найдена или уже выдана\n", name)
}

func (l *Library) ReturnBook(name string) {
	for i, book := range l.Books {
		if book.Name == name {
			l.Books[i].Status = BookStatusAvailable
			fmt.Printf("\nКнига '%s' возвращена\n", name)
			return
		}
	}
	fmt.Printf("\nКнига '%s' не найдена\n", name)
}
