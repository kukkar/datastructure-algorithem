package dsarray

import (
	"testing"
)

func TestTwoSumProblem(t *testing.T) {

	input := []int{4, 5, 6, 7, 8, 9, 1, 2, 3} 
	result := TwoSumProblem(input,6)
	expected := make([][]int,0)
	t.Logf("success Extected %v recieved %v",expected,result)
}


func TestSearchInRotatedArray(t *testing.T) {

	input := []int{4, 5, 6, 7, 8, 9, 1, 2, 3} 
	search := 2
	value :=  SearchInRotatedArray(input,search)
	if value != search {
		t.Errorf("failed expected %d recieved %d",search,value)
	}else {
		t.Logf("success")
	}
}


func TestSoldierMeadalDistribution(t *testing.T) {

	input := []int{1 ,2 ,4 ,5, 5, 5, 4, 3, 2, 1} 
	//input := []int{1,0,2}
	output := []int{1,2,3,4,1,5,4,3,2,1}
	//output := []int{2,1,2}
	answ :=  SoldierMeadalDistribution(input)
	if !Equal(output,answ) {
		t.Errorf("failed expected %v recieved %v",output,answ)
	}else {
		t.Logf("success exptec %v asnw %v",output,answ)
	}
}


func TestCookDishSample1(t *testing.T) {
	n := 5
	i := 3
	input := []string{"FATOIL","FIBERSPINCH","CARBPOTATO","FATII","FIBERRW"}
	expected := "---FATOIL:FIBERSPINCH:FATII-"
	output := CookDish(n,i,input)
	if output != expected {
		t.Errorf("failed exptected %s and output %s",expected,output)
	} else {
		t.Logf("success expected %s rec %s",expected,output)
	}
}

func TestCookDishSample3(t *testing.T) {
	n := 12
	i := 4
	input := []string{"FATII","FIBERII","CARBII","FATII","FIBERII","FATII","FIBERII","CARBII","CARBII","FATII","FIBERII","CARBII"}
	expected := "-----FATII:FIBERII:FATII:FATII--CARBII:FIBERII:CARBII:CARBII---"
	output := CookDish(n,i,input)
	
	if output != expected {
		t.Errorf("failed exptected %s and output %s",expected,output)
	} else {
		t.Logf("success expected %s rec %s",expected,output)
	}
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

