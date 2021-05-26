package main

import (
	"fmt"
	"sync/atomic"
)

var (
	pointList = atomic.Value{}

)

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

	fmt.Println("-----------")

	p := &point{
		a:1,
	}
	pointMap := make(map[uint32]*point,1)
	pointMap[1] = p

	pointList.Store(pointMap)

	pointAto := pointList.Load().(map[uint32]*point)
	fmt.Println(pointAto[1])
	pointAto[1].a = 2

	fmt.Println(pointList.Load().(map[uint32]*point)[1])
}

type point struct {
	a int
}
