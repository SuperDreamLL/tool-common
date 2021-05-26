package main

import (
	"fmt"
	"sync"
)

type RankConfCache struct {
	confMap 	map[int]int
	confLock 	sync.RWMutex
}

type test struct {
	a int
}

func main()  {
	a := make([]*RankConfCache,10)
	fmt.Println(a)

	for i:=0; i<10;i++ {
		a[i] = &RankConfCache{}
	}
	fmt.Println(a)

	// 看懂了，这是强转的意思，例如  int()
	t := test(struct {
		a int
	}{})
	fmt.Println(t)
}
