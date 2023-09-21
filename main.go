package main

import (
	"fmt"
	"net/http"
)

func Handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>Hello world!</p>")
}

func main() {
	http.HandleFunc("/", Handlerfunc)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", nil)

}
