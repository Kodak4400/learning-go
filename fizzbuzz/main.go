package main

import (
	"fizzbuzz/pkg"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		r := pkg.FizzBuzz(i)
		fmt.Println(r)
	}
}
