package main

import (
	"fmt"
)

func main() {
	row := 5
	limit := 1 
	
	for i:=0;i<=row;i++ {
		for j:=0;j<limit;j++ {
		  fmt.Print("* ")
		}
		limit++
	  fmt.Println()
	}	
}
