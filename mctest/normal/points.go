package main

import "fmt"

func main()  {

	a := &point{
		a:1,
	}

	b := a
	b.a = 2

	fmt.Println(a, b)

	c := &point{}
	d := &point{}

	fmt.Println(c == d, c, d)
}

type point struct {
	a int
}
