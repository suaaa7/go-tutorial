package main

import (
	f "fmt"
	. "strings"
	"errors"
)

func div(i int, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return
	}
	result = i / j
	return
}

func main() {
	f.Println(ToUpper("hello world"))
	result, _ := div(2, 0)
	f.Println(result)
}
