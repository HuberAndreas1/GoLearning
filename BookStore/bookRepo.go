package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handleGetBooks(getCmd *flag.FlagSet, all *bool, id *string) {
	err := getCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if !*all && *id != "" {
		fmt.Println("You must use either --id or --all")
		os.Exit(1)
	}

	if *all {
		books := getAllBooks()

		fmt.Println(books)
	}

	if *id != "" {
		book, err := getBookById(id)
		checkErr(err)

		fmt.Println(book.Id, book.Title, book.Author, book.Price, book.ImageUrl)
	}

}

func getBookById(id *string) (book Book, err error) {
	books := getAllBooks()
	for _, book := range books {
		if book.Id == *id {
			return book, nil
		}
	}

	return book, errors.New("Book not found")
}

func handleAddBook(cmd *flag.FlagSet, id *string, title *string, author *string, price *string, url *string) {
	err := cmd.Parse(os.Args[2:])
	checkErr(err)
}

func handleDeleteBook(cmd *flag.FlagSet, id *string) {

}

func handleUpdateBook(cmd *flag.FlagSet, id *string, title *string, author *string, price *string, url *string) {

}

func getAllBooks() (books []Book) {
	data, err := os.ReadFile("./books.json")
	checkErr(err)
	err = json.Unmarshal(data, &books)
	checkErr(err)

	return books
}
