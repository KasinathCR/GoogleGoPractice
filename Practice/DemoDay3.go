package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MyMethod)
	http.ListenAndServe(":8006", nil)
}

//MyMethod - Just an Index method to demostrate the HTML tag support in Go!
func MyMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<p>Hi! Welcome to Day 3 of Go Language Tutorial<br/>
It's fun learning Go Language<br/>
I am getting <strong>better</strong> at learning Go!</p>`)
	fmt.Fprintf(w, "<p>I just entered a <strong>%s</strong> here</p>", "place-holder")
}
