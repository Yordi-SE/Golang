Library Management System Documentation
Overview
The Library Management System is a console-based application developed in Go. It allows users to manage books and members within a library. The system demonstrates the use of Go functionalities including structs, interfaces, methods, slices, and maps.


Folder Structure
The project is organized into the following directory structure:

library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod

Components
1. main.go
Purpose: Entry point of the application.
Functionality: Initializes the application, sets up console interaction, and invokes appropriate methods from controllers/library_controller.go.
2. controllers/library_controller.go
Purpose: Handles console input and interacts with the library service.
Functionality: Receives user input for actions such as adding a book, removing a book, borrowing a book, returning a book, and listing available or borrowed books. Calls corresponding methods in services/library_service.go.
3. models/book.go
Purpose: Defines the Book struct.
Fields:
ID (int): Unique identifier for the book.
Title (string): Title of the book.
Author (string): Author of the book.
Status (string): Status of the book, which can be "Available" or "Borrowed".
4. models/member.go
Purpose: Defines the Member struct.
Fields:
ID (int): Unique identifier for the member.
Name (string): Name of the member.
BorrowedBooks ([]Book): Slice to hold books borrowed by the member.
5. services/library_service.go
Purpose: Contains business logic and data manipulation functions.
Implementation: Implements the LibraryManager interface, which includes methods to:
AddBook(book Book): Adds a new book to the library.
RemoveBook(bookID int): Removes a book from the library by its ID.
BorrowBook(bookID int, memberID int) error: Allows a member to borrow a book if it is available.
ReturnBook(bookID int, memberID int) error: Allows a member to return a borrowed book.
ListAvailableBooks() []Book: Lists all available books in the library.
ListBorrowedBooks(memberID int) []Book: Lists all books borrowed by a specific member.
6. docs/documentation.md
Purpose: Provides documentation for the library management system.



Usage:

Running the Application
Navigate to the project directory:

cd library_management


Run the application:

go run main.go




Console Commands
Add a New Book (A): Prompts the user to enter book details and adds the book to the library.
Remove an Existing Book (R): Prompts the user to enter the book ID and removes the book if it exists.
Borrow a Book (Bo): Prompts the user to enter book ID and member ID to allow a member to borrow the book if it is available.
Return a Book (Re): Prompts the user to enter book ID and member ID to allow a member to return a borrowed book.
List All Available Books (LA): Displays all books currently available in the library.
List All Borrowed Books by a Member (LB): Prompts the user to enter member ID and displays all books borrowed by the specified member.