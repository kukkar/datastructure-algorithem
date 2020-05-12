package main
import (
	"personell/datastructure-algorithem/dsarray"
	"fmt"
)

func main() {
	input := []int{4, 5, 6, 7, 8, 9, 1, 2, 3} 
	dsarray.TwoSumProblem(input,6)
	fmt.Printf("reuslt %d",dsarray.SearchInRotatedArray(input,6))
	ar1 := []int{1, 3, 10, 11}
	ar2 := []int{2, 4, 6, 8}
	fmt.Printf("result %v",dsarray.MergeSortedArray2(ar1,ar2))
	triplet := []int{ 1, 4, 45, 6, 10, 8 }
	fmt.Printf("triplet sum %v",dsarray.TripletSum(triplet,22))
	findSingle := []int{1,1,4,2,4}
	fmt.Printf("dd %d",dsarray.FindSingle(findSingle))
	findMissing := []int{1,2,3,5,6,7,8}
	fmt.Printf("missing %d",dsarray.FindMissing(findMissing)) 
	findDuplicate := []int{1,1,4,2,4}
	fmt.Printf("duplicate %v", dsarray.FindDuplicate(findDuplicate)) 
	fmt.Printf("duplicate %v", dsarray.FindDuplicateInPlace(findDuplicate)) 
	maxArray := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Printf("maxSubArray %v",dsarray.MaximumSubarray(maxArray)) 

	maxProfit := []int{11,1,30,2,12,2,21}
	fmt.Printf("profite %d",dsarray.MaxProfit(maxProfit)) 
	mInterval := make([]dsarray.Interval, 0)

	mInterval = append(mInterval,dsarray.Interval{
		S : 6,
		E : 8,
	})
	mInterval = append(mInterval,dsarray.Interval{
		S:1,
		E:9,
	})
	mInterval = append(mInterval,dsarray.Interval{
		S:2,
		E:4,
	})
	mInterval = append(mInterval,dsarray.Interval{
		S:9,
		E:7,
	})
	fmt.Printf("data %v", dsarray.MergeIntervals(mInterval))
}