package main

import (
	"fmt"
	"time"
)

type user struct {
	firstname string
	lastname  string
	dob       string
	createdAt time.Time
}

func main() {
	userFirstname := get_user_data("Enter your First Name: ")
	userLastname := get_user_data("Enter your Last Name: ")
	userDob := get_user_data("Enter your Date Of Birth: ")

	var appUser user

	// New Instance of our struct type
	appUser = user{
		firstname: userFirstname,
		lastname:  userLastname,
		dob:       userDob,
		createdAt: time.Now(),
	}

	//.. Do something with that data
	// fmt.Println(userFirstname, userLastname, userDob)
	outputUserDetails(&appUser) // We are passing address of our struct object
	// Else a new copy of struct is created in memory when this func is called

}

func outputUserDetails(u *user) {
	fmt.Println(u.firstname, u.lastname, u.dob, u.createdAt)
}

func get_user_data(prompt_text string) string {
	fmt.Print(prompt_text)
	var value string
	fmt.Scan(&value)
	return value
}
