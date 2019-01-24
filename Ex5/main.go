package main

import "fmt"

type numboor int

var x numboor
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T", y)
	fmt.Println("\n", y)

}
