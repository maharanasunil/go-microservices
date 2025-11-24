package pointers

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age // pointer variable

	/* NOTES:
	- If we add &var_name => We get the address of the variable
	- If we have address and we want to see what value we have => *address_of_variable
	*/

	fmt.Println("Pointer address of Age variable:", agePointer)

	// de-referencing
	fmt.Println("Dereferenced Age variable:", *agePointer)

	adultYears := editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", adultYears)

	fmt.Println("Age Updated:", age)
}

func editAgeToAdultYears(age *int) int {
	// return *age - 18
	*age = *age - 18
	return *age
}
