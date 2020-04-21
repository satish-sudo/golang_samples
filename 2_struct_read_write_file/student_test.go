package main

import "testing"

func TestStudent (t *testing.T) {
	s:= newStudent("satish", "FE", 28, 12345)
	if s.Age > 30 {
		t.Log("Age is greater than 30 for class FE")
		t.FailNow()
	}

	s.updateContact(45678)
	if s.Contact != 45678 {
		t.Log("Invalid contact Number")
		t.FailNow()
	}

	t.Log("Tests passed successfully")
}
