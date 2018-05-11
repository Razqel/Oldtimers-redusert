package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"html/template"
)

type Parkering struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	AntallLedigePlasser string `json:"Antall_ledige_plasser"`
}

//variablaer av structene for å holde på JSON-dataen

var park []Parkering

func main() {
	http.HandleFunc("/1", Stavanger)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Stavanger(w http.ResponseWriter, r *http.Request){
	hentOgUnmarshal("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json", &park)
	HTMLhandtering(w, "html/Oblig4.html", park)
}

/*
 getAndUnmarshal gjør et API-kall og leser JSON-dataen og legger det inn i structs-variablene.
 */
func hentOgUnmarshal(s string, v interface{}) {
	//Henter jsondataen i s
	res, err := http.Get(s)
	defer res.Body.Close()
	errorHandling(err)
	//Leser kroppen til jsondataen
	jsonBytes, err2 := ioutil.ReadAll(res.Body)
	errorHandling(err2)
	//Legger jsondaten inn i variabelen v
	errorHandling(json.Unmarshal(jsonBytes, &v))
}

/*
	readAndExecute leser en string som er en filepath til en html-fil og kjører
	html-filen sammen med struct-variabelen
*/
func HTMLhandtering(w http.ResponseWriter,s string, v interface{}) {
	//Lager template av HTML-filene
	tmpl, err := template.ParseFiles(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Kjører HTML-templaten til writeren og mater den med interfacet
	if err := tmpl.Execute(w, v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
	Enkel errorhandler
 */
func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}