package main

import (
	"fmt"
	"time"
)

func main() {
	var x string = "Hello World"
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	
	fmt.Println(x)
	
  	fmt.Print("Enter a number: ")
  	var input float64
  	fmt.Scanf("%f", &input)

	output := input * 2

  	fmt.Println(output)
  	i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
  	for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
  	for i := 1; i <= 10; i++ {
       	if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }
    
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
			var x [5]int
    		x[4] = 100
		    fmt.Println(x)
	}
//http://golang-book.ru/chapter-06-arrays-slices-maps.html
}
