package main

import (
	"html/template"
	"os"
)

type User struct {
	Name         string
	Loggedin     bool
	Age          int
	Bio          string
	Meta         UserMeta
	Phonenumbers map[string]string
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	u1 := User{
		Loggedin: false,
		Name:     "Bruce Wayne",
		Age:      37,
		Bio:      `<script>alert("Haha, you've been hacked!");</script>`,
		Meta: UserMeta{
			Visits: 4,
		},
		Phonenumbers: map[string]string{
			"miss":      "666-333-3333",
			"odd_job":   "123-463-5673",
			"Agent_005": "123-423-4321",
		},
	}

	err = t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}
}
