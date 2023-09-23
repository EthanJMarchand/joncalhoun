package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ethanjmarchand/joncalhoun/controllers"
	"github.com/ethanjmarchand/joncalhoun/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	hometmpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/", controllers.StaticHandler(hometmpl))

	contacttmpl, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/contact", controllers.StaticHandler(contacttmpl))

	faqtmpl, err := views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/faq", controllers.StaticHandler(faqtmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 | Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
