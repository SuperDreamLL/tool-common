package main

import (
	"fmt"
	"strconv"
)

func main()  {

	str := make([]string,0,10)

	for i := 0 ; i < 10; i++ {
		str = append(str, strconv.Itoa(i),strconv.Itoa(i+100))
	}

	fmt.Println(str)
}
