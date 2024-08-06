/*Problem:

Create a Go program that defines a struct to represent a Book. Each Book should have a title, author, and number of pages. Your task is to initialize a few Book instances and print their details.

Details:

Define a struct named Book with the following fields:
Title (string)
Author (string)
Pages (int)
Initialize at least three instances of the Book struct with different values.
Write a function printBookDetails that takes a Book as an argument and prints its details.
Call printBookDetails for each Book instance to print their details.
Example Output:

If the struct instances are initialized with the following values:

Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Pages: 380
Title: "Go in Action", Author: "William Kennedy", Pages: 300
Title: "Introducing Go", Author: "Caleb Doxsey", Pages: 200'

The Output should be:

Title: The Go Programming Language, Author: Alan A. A. Donovan, Pages: 380
Title: Go in Action, Author: William Kennedy, Pages: 300
Title: Introducing Go, Author: Caleb Doxsey, Pages: 200

*/

package main

import "fmt"

type book struct{
	Title  string
	Author string
	Pages  int
}

func printBookDetails(details book) {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", details.Title, details.Author, details.pages)
}

func main() {
	book1 := ("A Caged Bird Sings", "Maya Angelou", 205)
	book2 := ("The Color Purple", "Alice Walker", 245)
	book3 := ("The Price of a Ticket", "James Baldwin", 705)

	printBookDetails(book1, book2, book3)
}

