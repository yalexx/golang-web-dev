package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", "Dog")
	if err != nil {
		fmt.Println("dog route error")
	}

}

func me(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "me.gohtml", "Yanko")
	if err != nil {
		fmt.Println("me route error")
	}
}
