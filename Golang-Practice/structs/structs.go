package main

import (
	"fmt"

	"example.com/golang-practice/user"
)

func main() {
	userFirstname := get_user_data("Enter your First Name: ")
	userLastname := get_user_data("Enter your Last Name: ")
	userDob := get_user_data("Enter your Date Of Birth: ")

	var appUser *user.User
	// New Instance of our struct type
	appUser, err := user.New(userFirstname, userLastname, userDob)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(userFirstname, userLastname, userDob)
	// OutputUserDetails(&appUser) // We are passing address of our struct object
	// Else a new copy of struct is created in memory when this func is called

	// We added a new method to our struct
	appUser.PrintUserData()
	appUser.ClearUserName()
	appUser.PrintUserData()

}

func get_user_data(prompt_text string) string {
	fmt.Print(prompt_text)
	var value string
	fmt.Scanln(&value)
	// Scan will wait till you enter a value even if you keep going t next line
	// Scanln will make sure to stop taking value when you move to next line
	return value
}
