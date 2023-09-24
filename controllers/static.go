package controllers

import (
	"html/template"
	"net/http"

	"github.com/ethanjmarchand/joncalhoun/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is the best way to learn code?",
			Answer:   "You have to start coding!",
		},
		{
			Question: "What are the three rules of real estate?",
			Answer:   "Location, Location, Location",
		},
		{
			Question: "What is the average air speed velocity of a swallow?",
			Answer:   "African, or American?",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
