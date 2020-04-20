package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Name  	string	`json:"name"`
	Class 	string	`json:"class"`
	Age  	int		`json:"age"`
	Contact int		`json:"contact"`
}

func newStudent (name string, class string, age int, contact int) *Student {
	s:= Student {
		Name:name,
		Class:class,
		Age:age,
		Contact:contact,
	}
	return &s
}
func (s *Student) print() {
	fmt.Println("\nStudent Info")
	fmt.Printf("Name: %s  Class:%s  Age:%d  Contact:%d", s.Name, s.Class, s.Age, s.Contact)
}

func(s *Student) updateContact(argContact int) {
	//s.contact = argContact
	ps := &s
	(*ps).Contact = argContact
}

func (s *Student) writeToFile(fileName string) {
	data, _ := json.Marshal(s)
    err := ioutil.WriteFile(fileName, data, 0777)
    if err != nil {
    	fmt.Println("Failed to write data to file")
	}
}

func (s *Student) readFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("\nFailed to open file for read")
	}

	err = json.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Println("\nFailed to unmarshal data")
	}
}