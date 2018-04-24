package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"html/template"
)

type Apihunter struct {
	Name           	string
	Id             	int
	Uri 			string
	Urls			[]string
	Resource_url   	string
	Releases_url   	string
	Realname       	string
	Profile        	string
	Data_quality   	string
	Namevariations 	[]string

}


func postMalone(w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("https://api.discogs.com/artists/4016434")

	temp, _ := ioutil.ReadAll(res.Body)

	var apihunter Apihunter
	err := json.Unmarshal(temp, &apihunter)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	Pagevariables := Apihunter {
		Name: apihunter.Name,
		Id: apihunter.Id,
		Uri: apihunter.Uri,
		Urls: apihunter.Urls,
		Resource_url: apihunter.Resource_url,
		Releases_url: apihunter.Releases_url,
		Realname: apihunter.Realname,
		Profile: apihunter.Profile,
		Data_quality: apihunter.Data_quality,
		Namevariations: apihunter.Namevariations,
	}

	/*       //Bruker man kun disse koder for man vanlig rein tekst uten template. Det er lik for de andre funksjoner
	fmt.Fprintln(w, "Artist Name: ", apihunter.Name)
	fmt.Fprintln(w, "Id: ", apihunter.Id,)
	fmt.Fprintln(w, "The official link:", apihunter.Uri)
	info := strings.Join(apihunter.Urls, ", ")
	information := strings.Split(info, "[]")
	fmt.Fprintln(w, "Other links with information:", strings.Join(information, "[]"))
	fmt.Fprintln(w, "Resource URL: ", apihunter.ResourceUrl)
	fmt.Fprintln(w, "Releases URL : ", apihunter.ReleasesUrl)
	fmt.Fprintln(w, "Real Name: ", apihunter.Realname)
	fmt.Fprintln(w, "Profile: ", apihunter.Profile)
	fmt.Fprintln(w, "Data Quality : ", apihunter.DataQuality)
	NameV := strings.Join(apihunter.Namevariations, ", ")
	Namevar := strings.Split(NameV, "[]")
	fmt.Fprintln(w,"Name variations : ", strings.Join(Namevar, "[]"))
	*/

	t, err := template.ParseFiles("Oblig3_2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, Pagevariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


func theBeatles(w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("https://api.discogs.com/artists/82730")

	temp, _ := ioutil.ReadAll(res.Body)

	var apihunter Apihunter
	err := json.Unmarshal(temp, &apihunter)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	Pagevariables := Apihunter {
		Name: apihunter.Name,
		Id: apihunter.Id,
		Uri: apihunter.Uri,
		Urls: apihunter.Urls,
		Resource_url: apihunter.Resource_url,
		Releases_url: apihunter.Releases_url,
		Realname: apihunter.Realname,
		Profile: apihunter.Profile,
		Data_quality: apihunter.Data_quality,
		Namevariations: apihunter.Namevariations,
	}


	t, err := template.ParseFiles("Oblig3_2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, Pagevariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}

}


func ACDC(w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("https://api.discogs.com/artists/84752")

	temp, _ := ioutil.ReadAll(res.Body)

	var apihunter Apihunter
	err := json.Unmarshal(temp, &apihunter)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	Pagevariables := Apihunter {
		Name: apihunter.Name,
		Id: apihunter.Id,
		Uri: apihunter.Uri,
		Urls: apihunter.Urls,
		Resource_url: apihunter.Resource_url,
		Releases_url: apihunter.Releases_url,
		Realname: apihunter.Realname,
		Profile: apihunter.Profile,
		Data_quality: apihunter.Data_quality,
		Namevariations: apihunter.Namevariations,
	}


	t, err := template.ParseFiles("Oblig3_2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, Pagevariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


func Drexciya(w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("https://api.discogs.com/artists/1172")

	temp, _ := ioutil.ReadAll(res.Body)

	var apihunter Apihunter
	err := json.Unmarshal(temp, &apihunter)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	Pagevariables := Apihunter {
		Name: apihunter.Name,
		Id: apihunter.Id,
		Uri: apihunter.Uri,
		Urls: apihunter.Urls,
		Resource_url: apihunter.Resource_url,
		Releases_url: apihunter.Releases_url,
		Realname: apihunter.Realname,
		Profile: apihunter.Profile,
		Data_quality: apihunter.Data_quality,
		Namevariations: apihunter.Namevariations,
	}


	t, err := template.ParseFiles("Oblig3_2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, Pagevariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


func bobMarley(w http.ResponseWriter, r *http.Request){
	res, _ := http.Get("https://api.discogs.com/artists/41441")

	temp, _ := ioutil.ReadAll(res.Body)

	var apihunter Apihunter
	err := json.Unmarshal(temp, &apihunter)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	Pagevariables := Apihunter {
		Name: apihunter.Name,
		Id: apihunter.Id,
		Uri: apihunter.Uri,
		Urls: apihunter.Urls,
		Resource_url: apihunter.Resource_url,
		Releases_url: apihunter.Releases_url,
		Realname: apihunter.Realname,
		Profile: apihunter.Profile,
		Data_quality: apihunter.Data_quality,
		Namevariations: apihunter.Namevariations,
	}


	t, err := template.ParseFiles("Oblig3_2.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, Pagevariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


func main() {
	//funksjoner til de forskellige stier.
	http.HandleFunc("/1", postMalone)
	http.HandleFunc("/2", theBeatles)
	http.HandleFunc("/3", ACDC)
	http.HandleFunc("/4", Drexciya)
	http.HandleFunc("/5", bobMarley)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}