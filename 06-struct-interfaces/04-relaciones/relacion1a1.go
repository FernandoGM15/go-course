package main

import "fmt"

type User struct {
	name   string
	email  string
	active bool
}

// Relación 1-1
type Student struct {
	user User
	code string
}

func main() {
	fer := User{
		name:   "Fernando Garcia",
		email:  "fgarcia@gmail.com",
		active: true,
	}

	vane := User{
		name:   "Vanessa Quiroz",
		email:  "vquiroz@gmail.com",
		active: true,
	}

	ferStudent := Student{
		user: fer,
		code: "admin123.$",
	}

	fmt.Println(vane)
	fmt.Println(ferStudent.user.name)
}
