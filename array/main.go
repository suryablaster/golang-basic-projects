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
	arr := [5]int{
		1, 2, 3, 4, 5,
	}

	for index, val := range arr {
		fmt.Println(index, "--", val)
	}

	fmt.Println(len(arr))
}
