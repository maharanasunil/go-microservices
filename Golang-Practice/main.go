package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	// := is used to declare the variable inside a function
	sum := 0

	// for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	sum = 1
	for {
		if sum > 100 {
			break
		}
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("=====================Switch case=====================")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}

	// fmt.Println("=====================DEFER=====================")
	// Defer
	// defer fmt.Println("world") // this gets stacked/ pushed into stack and then
	// // program starts to execute normally
	// fmt.Println("hello")

	// // defer with loop
	// fmt.Println("Counting")
	// for i := range 10 {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

	fmt.Println("=====================Reference Types(Pointers)=====================")
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) // read i through the pointer

	*p = 21 // set i through the pointer
	fmt.Println(i)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("=====================Structs=====================")
	fmt.Println(Vertex{1, 2})

	v := Vertex{10, 20}
	v.X = 100
	fmt.Println(v.X, v.Y)

	// Pointers and structs together
	vertex_address := &v
	// println(vertex_address)
	vertex_address.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p_v4, v2, v3)

	fmt.Println("=====================Arrays=====================")
	// Collection of same type of data types
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("=====================Slices=====================")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // [start:end-1]
	fmt.Println(s)

	// Slices experiment
	names := []string{
		"John",
		"Paul",
		"George",
		"Sunil",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXXXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	// When we dont pass any number in the [] then a slice is created
	// slice literal initialization of type struct
	slice_struct := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, true},
	}
	fmt.Println(slice_struct)

	// continue at 2:42:13

}

// Struct declaraton
type Vertex struct {
	X int
	Y int
}

// Structs initialization
var (
	v1   = Vertex{1, 2}  // has type Vertex
	v2   = Vertex{X: 1}  // Y:0 is implicit
	v3   = Vertex{}      // X:0, Y:0
	p_v4 = &Vertex{1, 2} // has type *Vertex
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
