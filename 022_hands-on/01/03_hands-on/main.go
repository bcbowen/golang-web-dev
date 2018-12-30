package main

import (
	"html/template"
	"io"
	"net/http"
)

/*

ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.

1. Take the previous program in the previous folder and change it so that:
* a template is parsed and served
* you pass data into the template

*/

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("something.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ello")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "something.gohtml", "Ben")

}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "woof, man")

}
