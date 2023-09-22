package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
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

	// u1 := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "John Wayne",
	// 	Age:  32,
	// }

	u1 := User{
		Name: "Bruce Wayne",
		Age:  37,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}
}
