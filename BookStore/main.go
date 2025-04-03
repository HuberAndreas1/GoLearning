package main

import (
	"flag"
	"fmt"
	"os"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

//# get books
//./bookstore get --all
//./bookstore get --id 5
//
//# add a book with id ,title, author, price, image_url
//./bookstore add --id 6 --title test-book --author akilan --price 200 --image_url http://akilan.com/test.png
//
//# update a book with id ,title, author, price, image_url
//./bookstore update --id 6 --title test-book-1 --author akilan1 --price 2001 --image_url http://akilan.com/test.png1
//
//# delete a book by --id
//./bookstore delete --id 6

func main() {
	// get books
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all books")
	getById := getCmd.String("id", "", "Get Book by id")

	// add a book
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addId := addCmd.String("id", "", "Add Book by id")
	addTitle := addCmd.String("title", "", "Add Book by title")
	addAuthor := addCmd.String("author", "", "Add Book by author")
	addPrice := addCmd.String("price", "", "Add Book by price")
	addImageUrl := addCmd.String("image", "", "Add Book by image")

	// update a book
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.String("id", "", "Update Book by id")
	updateTitle := updateCmd.String("title", "", "Update Book by title")
	updateAuthor := updateCmd.String("author", "", "Update Book by author")
	updatePrice := updateCmd.String("price", "", "Update Book by price")
	updateImageUrl := updateCmd.String("image", "", "Update Book by image")

	// delete a book by its id
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "Delete Book by id")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get', 'add', 'update' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGetBooks(getCmd, getAll, getById)
	case "add":
		handleAddBook(addCmd, addId, addTitle, addAuthor, addPrice, addImageUrl)
	case "update":
		handleUpdateBook(updateCmd, updateId, updateTitle, updateAuthor, updatePrice, updateImageUrl)
	case "delete":
		handleDeleteBook(deleteCmd, deleteId)

	}
}


