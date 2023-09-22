package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta UserMeta
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
		Name: "Bruce Wayne",
		Age:  37,
		Bio:  `<script>alert("Haha, you've been hacked!");</script>`,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}
}
