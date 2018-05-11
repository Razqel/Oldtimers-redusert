//
// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.
//
// This program starts a server which can be used to find the most related issues, depending on the terms given by the user.
// The user writes in terms/keywords in the search box and this program will find the 30 issues that relate the most.
// This is useful for either developers who want to help others with their issues, or for developers who have come across
// an issue they can't solve. By searching for closed issues with related names they might be able to find a solution to their
// problem.
//
//
//
//
// !! Either run this main.go file straight from the terminal and go to localhost:8080/1 in your browser
// !! or find the program ready for use in the bin directory.
// !! Once you've opened your browser at localhost:8080/1 follow the given instructions.







package main

import (
	"log"
	"Oldtimers-redusert/Github/Oblig4/template"
	"net/http"
	"os"
	"strings"
	"Oldtimers-redusert/Github/Oblig4/struct.func"
)

//variable to store data of what goes wrong.
var logger *log.Logger

//function to print what issues/problems have caused the program to stop.
func Init() {
	logger = log.New(os.Stderr, "The following problem has been found: ", log.LstdFlags)
}

//function to write the template on the main page.
func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template.HTMLMain(w)
}

//function to get the keywords and search and unmarshal the api's that have been found.
func Search(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	v := r.FormValue("keywords") // to get the keywords from the text box
	keywords := strings.Split(v, " ")
	result, err := Oblig4.SearchAndUnmarshal(keywords)
	if err != nil {
		logger.Print(err) //writes to the terminal what went wrong
	} else {
		w.Header().Set("Content-Type", "text/html")
		template.HTMLSearch(w, result) //if everything goes as expected write the template + data
	}
}

//start the server at port 8080 and handle the given functions.
func main() {
	http.HandleFunc("/1", MainPage)
	http.HandleFunc("/search", Search)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
