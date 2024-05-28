package main

import "fmt"

type student struct {
	name    string
	age     int
	marks   int
	marsks2 int
}

func (s student) add(a, b int) int {
	return s.marks + s.marsks2
}

func main() {
	mp := map[int]string{
		1: "surya",
		2: "raju",
	}
	for index, val := range mp {
		fmt.Println(index, "--", val)
	}
}
