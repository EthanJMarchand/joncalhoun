package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func faqhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("param.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, id)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homehandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqhandler)
	r.Get("/param/{id}", paramHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 | Page not found", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
