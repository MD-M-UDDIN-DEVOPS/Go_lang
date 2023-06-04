package main

import (
	"fmt"
)

func main() {

	var minhaj string = "Hello Minhaj"

	fmt.Println("minhaj")

}

##package main

import "fmt"

func main() {
        var a, b bool = true, false
        fmt.Println(a && b)
        fmt.Println(a || b)
}


## package main

import "fmt"

func main() {
        var a, b bool = false, false
        fmt.Println(a && b)
        fmt.Println(a || b)
}

func main() {
	// a program to determine positive / negative / zero
	number := 10
	if number > 0 {
	  fmt.Printf("Positive\n")
	}	else if number < 0 {
	  fmt.Printf("Negative\n")
	} else {
	  fmt.Printf("Zero\n")
	}
  }