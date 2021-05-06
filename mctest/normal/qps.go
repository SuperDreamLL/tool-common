package main

import (
	"fmt"
	"time"
)

func main()  {
	selectFun()

}

// 通过time.Sleep 做定时，sleep的时间是一致的但是每次间隔时间会大于10ms
func forearchFun()  {
	fmt.Println(time.Now().Unix())

	var start int64
	count := 0
	nowTime := time.Now()
	start = nowTime.UnixNano()


	for  {
		startTemp := time.Now()
		count += 1
		time.Sleep(time.Millisecond * 10)

		fmt.Println(time.Now().Sub(startTemp))
		if time.Now().Sub(nowTime).Milliseconds() >= 1000 {
			break
		}
	}
	fmt.Println(time.Now().Unix(),time.Now().UnixNano()-start)
	fmt.Println(count)
}

// 通过time.Tick 做定时，更加准确对应时刻执行
func selectFun()  {
	fmt.Println(time.Now().Unix())
	var start int64
	count := 0
	c := time.Tick(time.Millisecond * 10)
	start = time.Now().UnixNano()
	for {
		startTime := time.Now()
		select {
		case <-c:
			count += 1
			if count >= 100 {
				goto finish
			}
		}
		fmt.Println(time.Now().Sub(startTime))
	}


	//for time.Now().UnixNano() - start < 10000000000 {
	//	count += 1
	//	time.Sleep(time.Millisecond * 10)
	//}
finish:
	fmt.Println(time.Now().Unix(),time.Now().UnixNano()-start)
	fmt.Println(count)
}
