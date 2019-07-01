package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(this(4,3))
}

	func this(x, y int)int{
		return x*y
	}