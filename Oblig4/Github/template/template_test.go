// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.
//
// Test function to test whether formatting of the templates are right.


package template

import (
	"bytes"
	"fmt"
	"testing"
	"Oldtimers-redusert/Github/Oblig4/struct.func"
)

func Test_SearchAndUnmarshal(t *testing.T) {

	result := Oblig4.Total{
		TotalCount: 1,
		Items: []*Oblig4.Information{
			{
				Title:   "getting started",
				Number:  3,
				HTMLURL: "http://test/3",
				State:   "open",      //might be closed by the time you read this
				User: &Oblig4.UserInfo{
					Login:   "figaro",
					HTMLURL: "http://test/user",
				},
			},
		},
	}

	buffer := bytes.NewBufferString("")
	HTMLMain(buffer)
	fmt.Printf(buffer.String())

	buffer2 := bytes.NewBufferString("")
	HTMLSearch(buffer2, &result)
	fmt.Printf(buffer2.String())
}
