package main

import (
	"Ploshad/geom"
	"Ploshad/run"
	"fmt"
)

func main() {

	r := geom.NewRect(2, 2)
	sRect := run.RunS(r)

	c := geom.NewCirc(3)
	sCirc := run.RunS(c)

	fmt.Println(r)
	fmt.Println(c)
	fmt.Println("sRect ", sRect)
	fmt.Println("sCirc ", sCirc)
}
