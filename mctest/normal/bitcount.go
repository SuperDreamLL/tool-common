package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(1<<8,math.MaxUint8)// 左移8位就是9位，-1就是8位的最大值
	fmt.Println(1<<16)
	fmt.Println(1<<2)

	a := 12
	b := a << 3
	c := b << 4

	fmt.Println(c, a << 7)

	fmt.Println(0<<1)


}
