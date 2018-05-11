// Copyright © 2018 Yaguel v. d. Meij, Xohan O. Barbosa, Andreas Wöllert, Vegard Trydal, Vegard Marvik.
//
// See https://developer.github.com/v3/search/#search-issues for more info.
//
//The templates that are used for our web application.


package template

import (
	"html/template"
	"io"
	"log"
	"Oldtimers-redusert/Github/Oblig4/struct.func"
	"time"
)

//template for the main page.
const templ =
	`<head>
        <link href='https://fonts.googleapis.com/css?family=Abel' rel='stylesheet'>
        <link rel="stylesheet" href="FormStyle.css">
        <title>Homepage Github Search</title>
	<style>
	div.container {
                width: 100%;
                border: 1px solid gray;
	}
	header, footer {
    	padding: 1em;
    	color: white;
    	background-color: #800000;
    	clear: left;
    	text-align: center;
	}
	form {
    	text-align: center;
	}
	body {
    	text-align: center;
    	font-family: 'Abel';font-size: 20px;
    	}
	</style>
		<script>
		function validateSearch() {
		var x = document.forms["Search"]["keywords"].value;
    	if (x == "") {
        alert("You have to fill in the text box before hitting the search button!");
        return false;
    		}
		}
		</script>
    	</head>
    		<body>
                <div>
        		<header><h1>Oldtimers-redusert</h1></header>
					
					<h2>Welcome to our search page for Github's 30 most related issues</h2>

					<p><b><u>How this works:</u></b><br>
					In the text box below you write your terms to define the issue you want to find from Github's website.<br>
					<i><small>There are several ways to search for issues. <br>
					<u>The more specific you are with your search terms the more related the issue will be. </u><br> 
					You can define your search either by repository (e.g. repo:"name of repository"), an open or closed issue <br>
					(e.g. is:open/closed), or the specifics of a name of an issue (e.g. decoding json), or by being very specific <br>
					one can search by repo, whether an issue is open or closed, and including name specifics <br>
					(e.g. repo:golang/go is:closed json decoder). <br> </i></p>
					
					<p><b> All of the repositories and other types of keywords that can be used, have to be part of the existing issues from Github. </b></p>
					<form name="Search" action="/search" onsubmit="return validateSearch() "method="post">
					Write the keywords here:  <input type="text" name="keywords" required pattern="[A-Za-z/ :-]+">
            		<input type="submit" value="Search">
        			</form>
					
					<p><b> Some examples of search terms, you can copy one of them to try it out. </b><br>
					<i><b>searching by repository: </i></b> <br>
					repo:GoesToEleven/GolangTraining <br>
					repo:golang/go <br> 
					<br>
					<i><b>Searching with repository followed by named issue: </i></b> <br>
					repo:GoogleCloudPlatform/golang-samples testing <br>
					repo:golang/go html <br>
					<br>
					<i><b>Searching with repository, closed/open issue, and name of issue: </i></b> <br>
					repo:golang/go is:closed xml <br>
					repo:golang/go is:open json decoder <br>
					repo:golang/go is:closed json unmarshal <br>
					<br>
					<i><b>Searching by the name of an issue</i></b> <br>
					xml, json, decoder, testing, <br>
					unmarshal json go <br>
					<br>
					<i><b>Searching for open or closed issues</i></b> <br>
					is:open, is:closed</p>
					

				<footer>Copyright &copy; Yaguel van der Meij, Andreas Wollert, Xohan Otero Barbosa, Vegard Trydal, Vegard Marvik</footer>
        		</div>
    		</body>
`

//template for the result page.
const templ2 =
	`<head>
        <link href='https://fonts.googleapis.com/css?family=Exo' rel='stylesheet'>
        <link rel="stylesheet" href="FormStyle.css">
        <title>Github Issues Template </title>
	<style>
	div.container {
                width: 100%;
                border: 1px solid gray;
	}
	header, footer {
    	padding: 1em;
    	color: white;
    	background-color: #800000;
    	clear: left;
    	text-align: center;
	}
	form {
    	text-align: center;
	}
	body {
    	text-align: center;
    	font-family: 'Exo';font-size: 20px;
    	}
	table, th, td {
    border: 1px solid black;
    border-collapse: collapse;
	}
	th, td {
    padding: 5px;
	text-align:center;
	}
	table#t01 tr:nth-child(even) {
    background-color: #eee;
	}
	table#t01 tr:nth-child(odd) {
   background-color: #fff;
	}
	table#t01 th {
    background-color: black;
    color: white;
	}
	table {
    width: 100%;
    margin-left: auto;
    margin-right: auto;
	}
	</style>
    	</head>
          <meta http-equiv="refresh" content="600" />
    		<body>
                <div>
        		<header><h1>Oldtimers-redusert</h1></header>
					<nav class="navbar">
            		<a class="tempA" href="/1">Go back to main page</a>

					<h4>Total number of issues is: {{.TotalCount}} </h4>
					<p>30 of the most related issues are shown below. <br>
					<i><small>To get more info on the issue press on the number, you can also visit the user by pressing their name.</small></i></p>
					<table id="t01">
					<tr style='text-align: left'>
					<th>Number</th>
					<th>Title</th>
					<th>State</th>
					<th>User</th>
					<th>Hours since created</th>
					<th>Days since created</th>
					</tr>
					{{range .Items}}
					<tr>
					<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
					<td>{{.Title}}</td>
					<td>{{.State}}</td>
					<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
					<td>{{.CreatedAt | HoursAgo}}</tb>
					<td>{{.CreatedAt | DaysAgo}}</tb>
					</tr>
					{{end}}
					</table>
					<br>
					<p><i><small>If this site is empty of information, please make sure your search terms are correct. 
							For more help with keywords visit <a href='https://developer.github.com/v3/search/#search-issues'> </i><br>
							Github's developer api</a> </small></p>
					
				<footer>Copyright &copy; Yaguel van der Meij, Andreas Wollert, Xohan Otero Barbosa, Vegard Trydal, Vegard Marvik</footer>
        		</div>
    		</body>
`


//variables to store the info of the templates.
var mainpage *template.Template
var search *template.Template

func HoursAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//function to initialize the templates.
func init() {
	mainpage = template.Must(template.New("mainpage").
		Parse(templ))
	search = template.Must(template.New("search").
		Funcs(template.FuncMap{"HoursAgo": HoursAgo}).
		Funcs(template.FuncMap{"DaysAgo": DaysAgo}).
		Parse(templ2))
}

//functions to execute the templates.
func HTMLMain(wr io.Writer) {
	if err := mainpage.Execute(wr, templ); err != nil {
		log.Fatal(err)
	}
}

func HTMLSearch(wr io.Writer, result *Oblig4.Total) {
	if err := search.Execute(wr, result); err != nil {
		log.Fatal(err)
	}
}

//pattern="[A-Za-z :/-]+"