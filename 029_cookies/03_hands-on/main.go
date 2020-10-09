package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Using cookies, track how many times a user has been to your website domain.
func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("visits")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "visits",
			Value: "1",
			Path:  "/",
		})
		fmt.Fprintln(w, "New Visitor, Welcome !")
	} else {

		fmt.Println("Val: ", c.Value)
		count, _ := strconv.Atoi(c.Value)
		count++
		// increment cookie
		// save on user pc

		http.SetCookie(w, &http.Cookie{
			Name:  "visits",
			Value: strconv.Itoa(count),
			Path:  "/",
		})

		fmt.Fprintln(w, "Returning user, cookie: ", strconv.Itoa(count))
	}

}
