// +build wireinject

package main

import "github.com/google/wire"

func InitTest(a int) *BigConf {
	wire.Build(c)
	return &BigConf{}
}
