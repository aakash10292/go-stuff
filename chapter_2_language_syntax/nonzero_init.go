package main

import "fmt"

func main() {
	a := 10
	b := "Hello"
	c := 3.14
	d := false

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)
}
