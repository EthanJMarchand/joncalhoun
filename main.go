package main

import (
	"fmt"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Welcome to my new site!</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>Contact Page</h1>
		<p>contact me at <a href="mailto:ethan@ethanmarchand.com">email me</a></p>
	`)
}

func faqhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>FAQ</h1>
		<h2><Strong>Q.</strong> What is the best way to get ahold of you?</h2>
		<p><strong>A.</strong> Phone!</p>
		<hr>
		<h2><Strong>Q.</strong> What is the best way to get ahold of you?</h2>
		<p><strong>A.</strong> Phone!</p>
		<hr>
		<h2><Strong>Q.</strong> What is the best way to get ahold of you?</h2>
		<p><strong>A.</strong> Phone!</p>
		<hr>
		<h2><Strong>Q.</strong> What is the best way to get ahold of you?</h2>
		<p><strong>A.</strong> Phone!</p>

	`)
}

// func pathhandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homehandler(w, r)
// 	case "/contact":
// 		contacthandler(w, r)
// 	default:
// 		http.Error(w, "404 | Page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homehandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqhandler(w, r)
	default:
		http.Error(w, "404 | Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", router)

}
