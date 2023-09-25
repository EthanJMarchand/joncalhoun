package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		Signup Template
	}
}

func (u Users) Signup(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Signup.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
