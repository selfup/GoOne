package main

import "fmt"

// Rectangle is a simple struct for learning
type Rectangle struct {
	length, width int
}

// Gostruct prints out the Rectangle type attributes
func main() {
	r := Rectangle{23, 50}
	fmt.Println("Default rectangle is: ", r)

	o := Rectangle{60, 90}

	fmt.Println(o.length)
	fmt.Println(o.width)
}
