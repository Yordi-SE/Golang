package manager

import (
	"errors"
	"myproject/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
	Manager LibraryManager
}

func (L *Library) AddBook(book models.Book) {
	L.Books[book.ID] = book
}
func (L *Library) RemoveBook(bookID int) {
	delete(L.Books, bookID)
}
func (L *Library) BorrowBook(bookID int, memberID int) error {
	var book, ok = L.Books[bookID]
	if !ok {
		return errors.New("book not foud")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is borrowed already")
	}
	member, exits := L.Members[memberID]
	if !exits {
		return errors.New("user not found")
	}
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	L.Books[bookID] = book
	L.Members[memberID] = member
	return nil
}

func (L *Library) ReturnBook(bookID int, memberID int) error {
	var book, ok = L.Books[bookID]
	if !ok {
		return errors.New("book not foud")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}
	member, exits := L.Members[memberID]
	if !exits {
		return errors.New("user not found")
	}
	book.Status = "Available"
	for i, v := range member.BorrowedBooks {
		if v.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	L.Books[bookID] = book
	L.Members[memberID] = member
	return nil
}

func (L *Library) ListAvailableBooks() []models.Book {
	var books []models.Book
	for _, v := range L.Books {
		if v.Status == "Available" {
			books = append(books, v)
		}
	}
	return books
}

func (L *Library) ListBorrowedBooks(memberID int) []models.Book {
	var member, ok = L.Members[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}
