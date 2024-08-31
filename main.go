package main

import (
	"fmt"
	"os"
)

func main() {
	v := os.Args[0:]
	if len(v) != 2 {
		return
	}
	file, err := os.ReadFile(v[1])
	if err != nil {
		fmt.Println(err)
	}
	c:=string(file)
	
}
