package main

import (
	"fmt"
	"net/http"
)

const pattern = "/hello"

func hellohandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "This is hello")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w,"This is about") // this show in console
	fmt.Fprintln(w, "This is about") // this show in browser
}

func main() {
	mux := http.NewServeMux() //this is router

	mux.HandleFunc(pattern, hellohandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server is running on port: 1234")

	err := http.ListenAndServe(":1234", mux)
	// if err == nill means no error // if err != nill means error

	//means error exist
	if err != nil {
		fmt.Println("Error prevents to start the server", err)
	}
}
