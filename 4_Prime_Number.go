package main

import (
	"fmt"
)

func main() {
    var num int
    num = 11
    prime := true
    for i:=2;i<num;i++ {
	if num%i == 0 {
	    prime = false
	    break
	}
    }

if prime {
	fmt.Printf("Number %d is prime",num)
} else {
	fmt.Printf("Number %d is not prime",num)
}
}
