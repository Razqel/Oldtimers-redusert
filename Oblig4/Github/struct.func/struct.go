// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.
//
//The structs used for this program

package Oblig4

import "time"

const Link = "https://api.github.com/search/issues"
//the rest of the link "?q={query}{&page,per_page,sort,order}" will be passed through the website using the SearchAndUnmarshal function.


//structs to hold the data from the github issues api
type Total struct {
	TotalCount int `json:"total_count"`
	Items      []*Information
}

type Information struct {
	Number    int
	HTMLURL   string 	`json:"html_url"`
	Title     string
	State     string
	User      *UserInfo
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfo struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

