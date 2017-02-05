package main

import (
	"fmt"

	"github.com/hypersolid/duckmap"
)

func main() {
	m := duckmap.NewMap()

	m.Set(4, "this")
	m.Set(5, "that")

	m.Delete(5)

	fmt.Println(4, m.Get(4).(string))
	fmt.Println(5, m.Get(5))
}
