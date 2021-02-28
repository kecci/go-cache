package main

import (
	"fmt"

	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(20).
		LRU().
		Build()
	err := gc.Set("key", "ok")
	if err != nil {
		fmt.Println(err.Error())
	}
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
