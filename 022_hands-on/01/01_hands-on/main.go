package main

import (
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

*/

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
	io.WriteString(w, "It's me!")

}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "woof, man")

}
