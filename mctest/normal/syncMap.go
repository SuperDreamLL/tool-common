package main

import (
	"bitbucket.org/funplus/golib/deepcopy"
	"fmt"
	"sync"
)

func main()  {

	for i:=0;i<10;i++ {
		globalForbidPlayers.Store(i,true)
	}

	globalForbidPlayers.Range(func(key, value interface{}) bool {
		k := key.(int)
		v := value.(bool)
		fmt.Println(k,v)
		if k == 2 {
			globalForbidPlayers.Delete(k)
		}
		return true
	})

	//fmt.Println(get())
	
}

func get() map[uint64]bool {
	return deepcopy.Copy(globalForbidPlayers).(map[uint64]bool)
}

var (
	globalForbidPlayers sync.Map
)
