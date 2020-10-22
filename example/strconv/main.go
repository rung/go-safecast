package main

import (
	"fmt"
	"github.com/rung/go-safecast"
)

func main() {
	s := "2147483647"
	converted, err := safecast.Atoi32(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(converted)
}
