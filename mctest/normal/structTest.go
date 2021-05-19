package main

import (
	"fmt"
	"sync"
)

type RankConfCache struct {
	confMap 	map[int]int
	confLock 	sync.RWMutex
}

func main()  {
	a := make([]*RankConfCache,10)
	fmt.Println(a)

	for i:=0; i<10;i++ {
		a[i] = &RankConfCache{}
	}
	fmt.Println(a)
}
