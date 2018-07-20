package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", About)
	http.ListenAndServe(":8005", nil)
}

//IndexHandler - To display content when accessing Go Web Server
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First Program based on Go's Web Server!\n")
	fmt.Fprintln(w, "Enjoying Learning Go!")
}

//About - To display About page content
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Displaying About Page")
}
