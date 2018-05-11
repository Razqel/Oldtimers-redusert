// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.

package Oblig4

import (
	"net/http"
	"fmt"
	"encoding/json"
	"net/url"
	"strings"
)


// SearchAndUnmarshal is a function that queries the GitHub issue tracker.
func SearchAndUnmarshal(terms []string) (*Total, error) {
	//using the given query string containing data so it can be passed to our web application
	q := url.QueryEscape(strings.Join(terms, " "))

	//getting the link and adding "?q=" + the search terms given from the website to fulfill the search terms for an api
	request, err := http.NewRequest("GET", Link+"?q="+q, nil)
	if err != nil {
		return nil, err  //error handling
	}
	request.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(request)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}


	var result Total //variable to hold the info of the api
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { //unmarshalling the json data
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}





