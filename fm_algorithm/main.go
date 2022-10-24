package main

import "fmt"

func main() {
	m := make(map[string][]int)
	for i := 0; i < 5; i++ {
		m["a"] = append(m["a"], i)
	}
	for _, iterm := range m["a"] {
		fmt.Println(iterm)
	}
}
