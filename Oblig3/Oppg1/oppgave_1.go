package main

import (
	"net/http"
	"log"
	"html/template"
)

type PageText struct {
	Text string
}

func helloClient(w http.ResponseWriter, r *http.Request) {
	TextPage := PageText{
		Text: "Hello Client",
	}

	t, err := template.ParseFiles("Oblig3_1.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, TextPage)
	if err != nil {
		log.Print("template executing error: ", err)
	}

	//bruker man kun den koden for man teksten Hello Client uten template.
	//fmt.Fprintf(w, "Hello Client")
}



func main() {
	http.HandleFunc("/", helloClient)

	//Sett opp en server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}