package main

import (
	"fmt"
	"sync/atomic"
)

func main()  {
	
	
	var a int32
	
	for i:=0;i<10000;i++ {
		go func() {
			atomic.AddInt32(&a,1)
			if atomic.LoadInt32(&a) == 999 {
				fmt.Println("atomic")
			}
			if a == 999 {
				fmt.Println(a)
			}

		}()

	}
	
}
