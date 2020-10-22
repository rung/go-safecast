package main

import (
	"fmt"
	"github.com/rung/go-safecast"
)

func main() {
	i := 2147483647
	converted, err := safecast.Int32(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(converted)
}
