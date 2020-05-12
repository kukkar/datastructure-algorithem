package dsstack

import (
	"testing"
	"reflect"
)

func TestFindNextGreater(t *testing.T) {

	input := []int{4,5,2,25}
	expected := [][]int{{4,5},{2,25},{5,25},{25,-1}}
	output := FindNextGreater(input)
	if reflect.DeepEqual(expected,output) {
		t.Logf("success expected %v output %v",expected,output)
	}else {
		t.Errorf("expected %v output %v",expected,output)
	}

}