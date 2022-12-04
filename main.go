package main

import "fmt"

var str string = "gyaru"

func main() {

	fmt.Printf("%v, %T", str, str)
	str := 50
	fmt.Println("----")
	fmt.Printf("%v, %T", str, str)

}
