package dsstring

import (
	"testing"
	"reflect"
)

func TestFindNonRepeatingCha(t *testing.T){


	input := "GeeksForGeeks"
	expected := "F"
	output := FindNonRepeatingCha(input)
	if expected != output {
		t.Errorf("Expected %s output %s",expected,output)
	}
	t.Logf("success %s expected %s output",expected,output)
}


func TestFindNonRepeatingChaNegativeCase(t *testing.T){

	input := "GeeksGeeks"
	expected := ""
	output := FindNonRepeatingCha(input)
	if expected != output {
		t.Errorf("Expected %s output %s",expected,output)
	}
	t.Logf("success %s expected %s output",expected,output)
}

func TestPrintAnagraTogether(t *testing.T) {

	input := []string{"cat","act","bat","abt"}

	expected := [][]string{{"cat","act"},{"bat","abt"}}
	output := PrintAnagraTogether(input)
	if reflect.DeepEqual(expected, output) {
		t.Logf("success %s expected %s output",expected,output)
	}else {
		t.Errorf("Expected %s output %s",expected,output)
	}
}

func TestLongestCommonPrefix(t *testing.T) {

	input := []string{"apple", "ape", "april"}
	expected := "ap"
	output := LongestCommonPrefix(input)
	if expected != output {
		t.Errorf("Expected %s output %s",expected,output)
	}else {
		t.Logf("success %s expected %s output",expected,output)
	}
	
}