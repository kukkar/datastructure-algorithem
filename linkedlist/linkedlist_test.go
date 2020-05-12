package linkedlist

import (
	"testing"
	"reflect"
)


func TestReverseLinkedList(t *testing.T) {
	var ll LL
	input := []int{4,3,2,1}

	for _,eachData := range input {
		ll.Insert(eachData)
	}
	ll.ReverseLinkedList()
	data := ll.TraverseLinkedList()
	if reflect.DeepEqual(input,data) {
		t.Logf("success expected %v got %v",input,data)
	}else {
		t.Errorf("Wrong Data expected %v got %v",input,data)
	}
}

func TestCheckPalindrome(t *testing.T) {
	var ll LL
	input := []int{1,2,5,2,1}
	for _,eachData := range input {
		ll.Insert(eachData)
	}
	expected := true
	output := ll.CheckPalindrome()
	if expected == output {
		t.Logf("success expected %v got %v",expected,output)
	}else {
		t.Errorf("Wrong Data expected %v got %v",expected,output)
	}
}

func TestAdd2Numbers(t *testing.T) {
	var l1 LL
	var l2 LL
	input1 := []int{6,4,9,5,7}
	input2 := []int{4,8}
	for _,eachData := range input1 {
		l1.Insert(eachData)
	}
	for _,eachData := range input2 {
		l2.Insert(eachData)
	}
	expected := []int{6,5,0,0,5}
	output := Add2Numbers(&l1,&l2)
	outputArray :=  output.TraverseLinkedList()
	if reflect.DeepEqual(outputArray,expected) {
		t.Logf("success expected %v got %v",expected,outputArray)
	}else {
		t.Errorf("Wrong Data expected %v got %v",expected,outputArray)
	}
}

func TestReverseTillK(t *testing.T){
	var ll LL
	input := []int{4,3,2,1}
	expected := []int{3,2,1,4}
	for _,eachData := range input {
		ll.Insert(eachData)
	}
	ll.ReverseTillK(3)
	data := ll.TraverseLinkedList()
	if reflect.DeepEqual(expected,data) {
		t.Logf("success expected %v got %v",expected,data)
	}else {
		t.Errorf("Wrong Data expected %v got %v",expected,data)
	}
}

func TestRemoveLoop(t *testing.T) {
	var ll LL
	input := []int{4,3,2,1}
	for _,eachData := range input {
		ll.Insert(eachData)
	}
	//generate loop
	ll.Head.Next.Next.Next.Next = ll.Head
	expected := []int{1,2,3,4}
	ll.RemoveLoop()
	output := ll.TraverseLinkedList()
	if reflect.DeepEqual(expected,output) {
		t.Logf("success expected %v got %v",expected,output)
	}else {
		t.Errorf("Wrong Data expected %v got %v",expected,output)
	}
}

func TestFindMid(t *testing.T) {
	var ll LL
	input := []int{6,5,4,3,2,1}
	for _,eachData := range input {
		ll.Insert(eachData)
	}
	expected := 4
	output := ll.FindMid()
	if expected != output {
		t.Errorf("Wrong Data expected %v got %v",expected,output)
	}else {
		t.Logf("success expected %v got %v",expected,output)
	}
}

func TestFindNth(t *testing.T) {
	var ll LL
	input := []int{6,5,4,3,2,1}
	for _,eachData := range input {
		ll.Insert(eachData)
	}
	expected := 3
	output := ll.FindNth(3)
	if expected != output {
		t.Errorf("Wrong Data expected %v got %v",expected,output)
	}else {
		t.Logf("success expected %v got %v",expected,output)
	}
}



// func TestReArrange(t *testing.T){
// 	var ll LL
// 	input := []int{4,3,2,1}
// 	expected := []int{1,4,2,3}
// 	for _,eachData := range input {
// 		ll.Insert(eachData)
// 	}
// 	ll.ReArrange()
// 	data := ll.TraverseLinkedList()
// 	if reflect.DeepEqual(expected,data) {
// 		t.Logf("success expected %v got %v",expected,data)
// 	}else {
// 		t.Errorf("Wrong Data expected %v got %v",expected,data)
// 	}
// }
