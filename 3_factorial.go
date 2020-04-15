package main

import (
	"fmt"
)

func factorial(num int) (int) {
    if num == 0 {
	return 1
    } else {
	return num * factorial(num-1);  	
    }
}
func main() {
    var num int
    num = 5
    fact := factorial(5)	
    fmt.Printf("Factorial of %d is %d",num, fact)
}
