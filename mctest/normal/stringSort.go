package main

import (
	"fmt"
	"sort"
)

func main()  {
	strslice := []string{"127.0.0.1:1", "127.0.0.1:2", "127.0.0.1:3", "127.0.0.1:4"}
	strslice = []string{}
	sort.Sort(sort.StringSlice(strslice))
	fmt.Println(strslice)
}

