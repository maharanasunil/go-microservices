package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age // pointer variable

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
