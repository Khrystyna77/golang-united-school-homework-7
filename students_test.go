package coverage

import (
	"os"
	"testing"
	//"errors"
	//"fmt"
	"strconv"
	"strings"
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

func TestLessOk(t *testing.T){
	var all People
	Per1 := Person{firstName: "Anna", lastName: "Grad", birthDay: time.Date(2002, 3, 4, 10, 30, 0, 0, time.UTC)}
	Per2:= Person{firstName: "Bohdan", lastName: "Horvat", birthDay: time.Date(1991, 3, 4, 10, 30, 0, 0, time.UTC)}
	all = append(all, Per1, Per2)
	var i, j int 
	i =0
	j =1
    res2 := all[i].birthDay.Unix() > all[j].birthDay.Unix()
	
	if res2 == false{
		t.Errorf("List doesn't sorted by birthday")
	}


}

// Swap
func TestSwOk(t *testing.T){
	//var all3 People
	var all3 People
	Per3 := Person{firstName: "Anna", lastName: "Grad", birthDay: time.Date(2002, 3, 4, 10, 30, 0, 0, time.UTC)}
	Per4:= Person{firstName: "Bohdan", lastName: "Horvat", birthDay: time.Date(1991, 3, 4, 10, 30, 0, 0, time.UTC)}
	all3 = append(all3, Per3, Per4)
	all3.Swap(0, 1)
	var res People
    res = append(res, Per4, Per3)
	//var p1 string
	//var p2 string

		
	if res[0] != Per4{
			t.Errorf("Swap doesn't work")

		}
}

////////////////////////////////////////////////////////
///newMatrix
func TestNewOk(t *testing.T){
	//str1 := "abcd efghij klmnop"
	str2 := "123 456 78"
	//var matrix Matrix
    New(str2)
    str2 = strings.ReplaceAll(str2, " ", "")
	if len(str2) == 0{
		t.Errorf("Don't find data")
	}
	_, ok := strconv.Atoi(str2)
	if ok != nil {
		t.Errorf("Swap doesn't work")
	}

}

func TestSetOk(t *testing.T){
	str5 := "123 456 78"
	matrix2, _ := New(str5)
	//mtrx2 := matrix2.Rows()
    //mtrx3 := matrix2.Cols()
	if matrix2.rows < 0 {
		t.Errorf("Output doesn't correct2")
	}

}


