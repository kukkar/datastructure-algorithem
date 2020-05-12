package dsarray
import (

	"fmt"
)

func sortInterval(arr []Interval)[]Interval {

	for i:=0;i<len(arr);i++ {
		minIndex := i
		for j:=i;j<len(arr);j++ {
			if arr[minIndex].S > arr[j].S {
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
	return arr
}

func max(a,b int)int {
	if a>=b {
		return a
	} 
	return b
}

func min(a,b int)int {
fmt.Printf("iput %d second %d",a,b)
	if a <=b {
		fmt.Printf("return value %d",a)
		return a
	}
	fmt.Printf("resutn %d",b)
	return b
}

func deleteIndex (arr []int)[]int {


	return nil
}