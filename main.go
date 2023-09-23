package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}

func faqhandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

type Data struct {
	Id string
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := Data{
		Id: id,
	}
	w.Header().Set("content-type", "text/html; charset=utf-8")
	tplpath := filepath.Join("templates", "param.gohtml")
	tpl, err := template.ParseFiles(tplpath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
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
