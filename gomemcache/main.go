package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("127.0.0.1:8080")
	err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	if err != nil {
		fmt.Println(err.Error())
	}

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(it)
}
