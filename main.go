package main

import "fmt"
import (
	"github.com/golang/example/stringutil"
)

func Sum(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

func Reverse(x string, y string) string {
	return stringutil.Reverse(x + " " + y)
}

func main() {
	fmt.Println(Sum(5, 5))
	fmt.Println(Multiply(5,1))
	fmt.Println(Reverse("hello", "aslan"))
}