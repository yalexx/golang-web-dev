package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

func index(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(writer, "Welcome!\n")
}
