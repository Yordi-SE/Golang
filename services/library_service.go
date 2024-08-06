package service

import (
	"bufio"
	"fmt"
	manager "myproject/controllers"
	"myproject/models"
	"os"
	"strconv"
	"strings"
)

var LibraryObject = manager.Library{
	Books:   make(map[int]models.Book),
	Members: make(map[int]models.Member),
}

func Console() {
	// code for console app
	reader := bufio.NewReader(os.Stdin)
	var err error
	var memberId int
	var Id string
	var BookId int
	fmt.Println("Enter Command: A | R | Bo | Re | LA | LB")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)
	if command == "A" {
		fmt.Print("Enter Title: ")
		Title, _ := reader.ReadString('\n')
		Title = strings.TrimSpace(Title)
		fmt.Print("Enter Author: ")
		Author, _ := reader.ReadString('\n')
		Author = strings.TrimSpace(Author)
		fmt.Print("Enter Status: ")
		Status, _ := reader.ReadString('\n')
		Status = strings.TrimSpace(Status)
		LibraryObject.AddBook(models.Book{
			ID:     len(LibraryObject.Books),
			Title:  Title,
			Author: Author,
			Status: Status,
		})
		fmt.Print(LibraryObject.Books)

	} else if command == "R" {
		fmt.Print("Enter Book Id: ")
		Id, _ = reader.ReadString('\n')
		BookId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || BookId < 0 {
				fmt.Print("Enter Valid Book Id: ")
				Id, _ = reader.ReadString('\n')
				BookId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}
		LibraryObject.RemoveBook(BookId)
	} else if command == "Bo" {
		fmt.Print("Enter Book Id: ")
		Id, _ = reader.ReadString('\n')
		BookId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || BookId < 0 {
				fmt.Print("Enter Valid Book Id: ")
				Id, _ = reader.ReadString('\n')
				BookId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}

		fmt.Print("Enter Member Id: ")
		Id, _ = reader.ReadString('\n')

		memberId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || memberId < 0 {
				fmt.Print("Enter Valid Member Id: ")
				Id, _ = reader.ReadString('\n')
				memberId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}
		errs := LibraryObject.BorrowBook(BookId, memberId)
		if errs != nil {
			fmt.Println(errs)
		}

	} else if command == "Re" {
		fmt.Print("Enter Book Id: ")
		Id, _ = reader.ReadString('\n')
		BookId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || BookId < 0 {
				fmt.Print("Enter Valid Book Id: ")
				Id, _ = reader.ReadString('\n')
				BookId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}

		fmt.Print("Enter Member Id: ")
		Id, _ = reader.ReadString('\n')

		memberId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || memberId < 0 {
				fmt.Print("Enter Valid Member Id: ")
				Id, _ = reader.ReadString('\n')
				memberId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}
		errs := LibraryObject.ReturnBook(BookId, memberId)
		if errs != nil {
			fmt.Println(errs)
		}
	} else if command == "LA" {
		fmt.Println(LibraryObject.ListAvailableBooks())
	} else if command == "LB" {
		fmt.Print("Enter Member Id: ")
		Id, _ = reader.ReadString('\n')

		memberId, err = strconv.Atoi(strings.TrimSpace(Id))

		for {
			if err != nil || memberId < 0 {
				fmt.Print("Enter Valid Member Id: ")
				Id, _ = reader.ReadString('\n')
				memberId, err = strconv.Atoi(strings.TrimSpace(Id))
			} else {
				break
			}
		}
		fmt.Println(LibraryObject.ListBorrowedBooks(memberId))
	}
}
