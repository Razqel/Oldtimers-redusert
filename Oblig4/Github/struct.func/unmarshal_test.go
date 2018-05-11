// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.
//
// Test function to test if the main function of the program is working and testing whether the outcome is as expected.

package Oblig4

import (
	"testing"
)

func Test_SearchAndUnmarshal(t *testing.T) {
	//https://github.com/datadewins/MeetUp/issues/3
	r, _ := SearchAndUnmarshal([]string{"repo:golang/go", "is:open", "getting started", "3"}) //this issue might be closed by the time you read this

	//test if this is the only case with the given specs.
	if r.TotalCount != 1 {
		t.Errorf("return[%v]", r.TotalCount)
	}

	//test if the right user has been given.
	if r.Items[0].User.Login != "ami-tewari" {
		t.Errorf("return[%v]", r.Items[0].User.Login)
	}

	//test if the issue number is right.
	if r.Items[0].Number != 3 {
		t.Errorf("return[%v]", r.Items[0].Number)
	}

}
