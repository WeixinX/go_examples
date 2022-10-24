package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i <= 5; i++ {
		s.Push(strconv.Itoa(i))
	}
	fmt.Println(s.Top().(string), s.Size())
	for !s.Empty() {
		a := s.Pop().(string)
		fmt.Println(a, s.Size())
	}
}
