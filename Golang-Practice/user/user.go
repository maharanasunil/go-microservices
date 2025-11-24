package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	dob       string
	createdAt time.Time
}

// Method 1) This function belongs to the struct user
func (u User) PrintUserData() {
	fmt.Println(u.firstname, u.lastname, u.dob, u.createdAt)
}

// Method 2) This function takes the user pointer and clears the first and last name
func (u *User) ClearUserName() {
	u.firstname = ""
	u.lastname = ""
}

// Method for our struct
func OutputUserDetails(u *User) {
	fmt.Println(u.firstname, u.lastname, u.dob, u.createdAt)
}

// Function 1) This function returns a pointer of our new created struct
func New(firstname, lastname, dob string) (*User, error) {
	if firstname == "" || lastname == "" || dob == "" {
		return nil, errors.New("ERROR!!! FirstName, LastName and DOB are required")
	}

	return &User{
		firstname: firstname,
		lastname:  lastname,
		dob:       dob,
		createdAt: time.Now(),
	}, nil
}
