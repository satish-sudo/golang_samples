package main

import "fmt"

func main() {
	defer fmt.Println("DONE with slice, MAP, defer tests")
	fmt.Println("Slice Test")
	var slice []string
	/*
	slice := make([]string, 3)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	*/

	val := 97
	for i:=0; i<6; i++ {
		slice = append(slice, string(val))
		val++
	}
	/*
	slice = append(slice, "a")
	slice = append(slice, "b")
	slice = append(slice, "c")
	slice = append(slice, "d")
	slice = append(slice, "e")
	slice = append(slice, "f")
	*/

	fmt.Printf("\nSlice Elements: %v",slice)

	len := lengthOfSlice(slice)
	fmt.Printf("\nSlice Length:%v",len)

	updateSlice(slice, 4, "eee")
	fmt.Printf("\nSlice Elements After Update: %v",slice)


	/*
	MAP TEST
	*/

	fmt.Println("\n\nMAP Test")
	//Create map for student record, key=roll number and value=full name
	student := make(map[int] string)
	student[5] = "s1"
	student[10] = "s2"
	student[15] = "s3"

	//Print roll number and full name of each student
	printMap(student)

	//Modify full name whose roll number is 10
	student[10] = "s4"
	fmt.Println("After updating name of student whose roll number is 10")
	printMap(student)

	//Remove student record from map whose  name is "s4"
	deleteFromMap(student, "s4")
	fmt.Println("Elements after deleting s4")
	printMap(student)
}

//Slice Functions
func lengthOfSlice(s []string) int {
	return len(s)
}

func updateSlice(s []string, pos int, val string) {
	fmt.Println("\nUpdated slice element")
	s[pos] = val
}

//MAP Functions
func printMap(student map[int]string) {

	fmt.Printf("\nMap Elements:\n")
	for k, v := range student {
		fmt.Printf("Key : %v Value: %v\n", k, v)
	}
}


func deleteFromMap(stud map[int]string, val string) {
	defer fmt.Println("Deleted element from MAP")
	for k, v := range stud {
		if v == val {
			delete(stud, k)
		}
	}
}
