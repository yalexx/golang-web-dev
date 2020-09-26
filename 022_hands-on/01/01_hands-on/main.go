package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func def(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "d")
}

func dog(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "dog")
}

func cat(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "cat")
}

func me(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Yanko")
}
