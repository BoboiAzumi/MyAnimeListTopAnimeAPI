package main

import (
	"myanimelisttop/src"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(2)
	src.Run()
}