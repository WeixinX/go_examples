package main

import (
	"testing"
)

var input = map[string][]string{
	"a": []string{"c", "e"},
	"b": []string{"c", "d"},
	"c": []string{"e", "f"},
	"d": []string{"f"},
	"e": []string{"g"},
	"f": []string{"g", "h"},
	"g": []string{},
	"h": []string{},
}

func TestNewFM(t *testing.T) {
	fm := NewFM(input)
	fm.HPG.Print()
	fm.PrintRecord()
}

func TestFM_Run(t *testing.T) {
	fm := NewFM(input)
	fm.Run()
	fm.PrintRecord()
	//for k, index := range fm.HPG.indexs {
	//	fmt.Println(k, index)
	//}
}
