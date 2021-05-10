package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(1<<8,math.MaxUint8)// 左移8位就是9位，-1就是8位的最大值
	fmt.Println(1<<16)
	fmt.Println(1<<2)
}
