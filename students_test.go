package coverage

import (
	"testing"
	"os"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

//per1 := Person{firstName: "Anna", lastName: "Grad", birthDay: time.Date(2002, 3, 4, 10, 30, 0, 0, time.UTC)}
//Per2:= Person{firstName: "Bohdan", lastName: "Horvat", birthDay: time.Date(1991, 3, 4, 10, 30, 0, 0, time.UTC)}
//Per3 := Person{firstName: "Vika", lastName: "Lestless", birthDay: time.Date(1991, 3, 4, 10, 30, 0, 0, time.UTC)}
//All := People{Per1, Per2, Per3}

func TestLenOK(t *testing.T){
	var all People
	Per1 := Person{firstName: "Anna", lastName: "Grad", birthDay: time.Date(2002, 3, 4, 10, 30, 0, 0, time.UTC)}
	Per2:= Person{firstName: "Bohdan", lastName: "Horvat", birthDay: time.Date(1991, 3, 4, 10, 30, 0, 0, time.UTC)}
	all = append(all, Per1, Per2)
	res := all.Len()
	if res != 2 {
		t.Errorf("Please pay attention to input")
	}
}


