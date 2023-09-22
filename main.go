package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Welcome to my new site!!!</p>")
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

func main() {
	r := chi.NewRouter()
	r.Get("/", homehandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqhandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 | Page not found", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
