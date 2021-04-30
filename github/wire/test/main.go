package main

import (
	"fmt"
	"github.com/google/wire"
)

type Conf struct {
	A int
}

type BigConf struct {
	B *Conf
}


func CreateConf(a int) *Conf {
	return &Conf{
		A:a,
	}
}

func TestConf(c *Conf) *BigConf {

	return &BigConf{
		B:c,
	}
}


var (
	c = wire.NewSet(CreateConf,TestConf)
)

func main()  {
	temp := InitTest(1)

	fmt.Println(temp.B)
}