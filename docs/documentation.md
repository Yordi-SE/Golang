# Getting Started

To run the applicaiton, run the following command in the root directory of the project:
```bash
go run .
```

## Commands

| Operation | Command | Description |
| - | - | - |
| Add book  |   `A`   | Adds a book in the in-memory db |
| Remove book | `R` | Removes a book from the library after prompting for the bookID |
| Borrow book | `Bo` | Lends a book specified by the bookID to the member specified by the memberID |
| Return book | `Re` | Returns a previously borrowed book specified by the bookID by the member specified by the memberID |
| List available books | `LA` | List of all the available (non-borrowed) books from the list of books |
| List borrowed books | `LB` | List of all the borrowed books from the list of books |

	"Remove book":          "R",
	"Borrow book":          "Bo",
	"Return book":          "Re",
	"List available books": "LA",
	"List borrowed books":  "LB",
