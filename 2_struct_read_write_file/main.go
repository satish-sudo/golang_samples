package main

import "fmt"

func main() {
	s:= newStudent("satish", "FE", 30, 12345)
	s.print()

	//Update Contact Info
	s.updateContact(67890)

	fmt.Println("\nAfter Updating Contact number ")
	s.print()

	var fileName string = "student.json"
	s.writeToFile(fileName)
	s1 := &Student{}
	s1.readFromFile(fileName)
	fmt.Println("\nRead data from file:")
	s1.print()
}

