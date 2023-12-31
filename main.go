package main

import (
	"fmt"
	"net/http"

	"github.com/ethanjmarchand/joncalhoun/controllers"
	"github.com/ethanjmarchand/joncalhoun/templates"
	"github.com/ethanjmarchand/joncalhoun/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{}
	usersC.Templates.Signup = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.Signup)
	r.Post("/users", usersC.Create)

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 | Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
