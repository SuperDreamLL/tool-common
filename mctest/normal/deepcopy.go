package main

import (
	"fmt"
	"github.com/mohae/deepcopy"
	"sync"
	"time"
)
//TODO 重要！！！ deepcopy包copy的结构体中变量是私有的小写开头，是无法copy的

type forbidPlayer struct {
	Players 	map[uint64]bool
	UpdateTime 	int64
	lock 		sync.RWMutex
}
type forbidPlayersCache struct {
	forbidMap	map[int]*forbidPlayer
	lock 		sync.RWMutex
}

func main()  {
	player := &forbidPlayer{
		Players:map[uint64]bool{1:true,2:true},
		UpdateTime:time.Now().Unix(),
	}
	cache := &forbidPlayersCache{
		forbidMap: map[int]*forbidPlayer{1:player},
	}
	a := deepcopy.Copy(cache.forbidMap).(map[int]*forbidPlayer)
	fmt.Println(a[1])
}
