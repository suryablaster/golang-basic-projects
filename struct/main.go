package main

import "fmt"

type raju interface {
	Demo() string
}
type flagshupmob struct {
	name       string
	price      int
	offerprice int
}

func (mob flagshupmob) Demo() string {
	return "jai shree ram"
}
func (mob flagshupmob) diff() int {
	return mob.price - mob.offerprice
}

func main() {

	mob1 := flagshupmob{
		name:       "pixel",
		price:      100000,
		offerprice: 90000,
	}

	fmt.Println(mob1.Demo())
}
