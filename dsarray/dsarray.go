package dsarray

import (
	 _  "fmt"
	"personell/datastructure-algorithem/sorting"
	"math"
)

// [1,5,-1,7,5]
func TwoSumProblem(arr []int, sum int)[][]int {
	low := 0
	high := len(arr)-1
	pair := make([][]int,0)
	for low < high {
		if arr[low]+arr[high] == sum {
			v := []int{arr[low],arr[high]}
			pair = append(pair,v)
		}
		if arr[low] + arr[high] > sum {
			high = high -1
		} else {
			low = low +1
		}
	}
	return pair
}

func SearchInRotatedArray(arr []int,key int) int{
	return search(arr,0,len(arr),key)
}

func MergeSortedArray(arr1 []int, arr2 []int)[]int {
	l1 := len(arr1)
	l2 := len(arr2)
	var p1,p2 int
	var iteration int
	result := make([]int,0)
	iteration = l1+l2

	for i:=0;i<iteration;i++ {
		if l1 == p1 {
			result = append(result,arr2[p2:]...)
			break
		} else if l2 == p2 {
			result = append(result,arr1[p1:]...)
			break
		}
		if arr1[p1]<=arr2[p2] {
			result = append(result,arr1[p1])
			p1 = p1 +1
		} else if arr1[p1]>= arr2[p2] {
			result = append(result,arr2[p2])
			p2 = p2+1
		}
	}
	return result
}


func MergeSortedArray2(arr1 []int, arr2 []int)[]int {
	l1 := len(arr1)
	l2 := len(arr2)
	var p1,p2 int
	p1 = 0
	p2 = l1 
	var iteration int
	arr1 = append(arr1,arr2...)
	iteration = l1+l2

	for i:=0;i<iteration;i++ {
		if l1 == p1  ||  p2 == iteration{
			break
		}
		if arr1[p1]<=arr1[p2] {
			p1 = p1 +1
		} else if arr1[p1]>= arr1[p2] {
			temp := arr1[p2]
			arr1[p2] = arr1[p1]
			arr1[p1] = temp
			p2 = p2+1
		}
	}
	return arr1
}
//
//{ 1, 4, 45, 6, 10, 8 }
// complexity n^2  as we sorting the array first
func TripletSum(arr []int,sum int)bool {
	arr = sorting.SelectionSort(arr)
	for i:=0;i<len(arr);i++ {
		l := i+1
		r := len(arr)-1
		for l<r {
			if arr[i]+arr[l]+arr[r] == sum {
				return true
			} else if arr[i]+arr[l]+arr[r] < sum {
				l++
			} else {
				r--
			}
		}
	}
	return false
}

func FindSingle(arr []int)int {
	var t int
	for _,eachVal := range arr {
		t = eachVal^t
	}
	return t
}

// o(n) 
//other ways we can discover later
func FindMissing(arr []int)int {
	arr = sorting.SelectionSort(arr)
	for i:=0;i<len(arr);i++ {
		if arr[i]+1 != arr[i+1] {
			return arr[i]+1
		}
	}
	return -1
}
//{12, 11, 40, 12, 5, 6, 5, 12, 11}
//using map find duplicate
func FindDuplicate(arr []int)[]int {
	resultSet := make([]int,0)
	dataMap := make(map[int]int)
	for _,eachData := range arr {
		if val,ok := dataMap[eachData];ok {
			if val ==1 {
				resultSet = append(resultSet,eachData)
			}
			dataMap[eachData] = dataMap[eachData]+1
		} else {
			dataMap[eachData] = 1
		}
	}
	return resultSet
}

//idea here is use array as hash map as the value of each index can not exceed the length of array
// so we will traverse the array and mark as -ve and it value is -ve it means it duplicate
func FindDuplicateInPlace(arr []int)[]int {

	result := make([]int,0)
	for i:=0;i<len(arr);i++ {
		if arr[int(math.Abs(float64(arr[i])))] >= 0 {
			arr[int(math.Abs(float64(arr[i])))] = -arr[int(math.Abs(float64(arr[i])))]
		} else {
			result = append(result,int(math.Abs(float64(arr[i]))))
		}
	}
	return result
}

// o(n) 
func MaximumSubarray(arr []int)[]int {

	var max,maxSoFar int
	var start, end,s int
	for i:=0;i<len(arr);i++ {
		maxSoFar = maxSoFar + arr[i]
		if max < maxSoFar {
			max = maxSoFar
			start = s
			end = i 
		}
		if maxSoFar < 0 {
			maxSoFar = 0
			s = s+1
		}
	}
	return arr[start:end+1]
}

func MaxProfit(arr []int)int {

	var buy bool
	var minima,profite int
	for i:=0;i<len(arr);i++ {
		if buy {
			if i == len(arr)-1 {
				buy = false
				profite += arr[i] - minima 
			} else if arr[i] > arr[i+1] {
				buy = false
				profite += arr[i] - minima 
			}			
		} else if i < len(arr)-1 && arr[i]  < arr[i+1] {
			buy  = true
			minima = arr[i]
		}
	}
	return profite
}

func MergeIntervals(arr []Interval ) []Interval{

	arr = sortInterval(arr)
	index := 0
	for i:=1;i<len(arr);i++ {
		if arr[index].E >= arr[i].S {
			//merge interval
			arr[index].E = max(arr[index].E,arr[i].E)
			arr[index].S = min(arr[index].S,arr[i].S)
		//	arr = append(arr[:i],arr[i+1:]...) 
		} else {
			arr[index+1] = arr[i]
			index++
		}
	}
	return arr[:index+1]
}

func SoldierMeadalDistribution(arr []int)[]int {
	
	result := make([]int,0)
	result = append(result,1) 
	for i :=1;i<len(arr);i++  {
		if arr[i-1]<arr[i] {
			result = append(result,result[i-1]+1)
		} else {
			result = append(result,1)
		}
	}
	for j := len(arr)-2;j>=0;j-- {
		if arr[j]>arr[j+1] {
			result[j] = max(result[j],result[j+1]+1)
		}
	}
	return result
}

func check(t []int)int {

	var p int
	for i:=0;i<len(t)-1;i++ {
		if t[i]> t[i+1] {
			p = p+1
		}else {
			return p
		}
	}
	return p
}

func search(arr []int,l int,h int,key int) int{
	if l >h {
		return -1
	}
	mid := (l+h)/2
	if key == arr[mid] {
		return mid
	}
	if arr[l] <= arr[mid] {
		//it mean out first phase of array is sorted 
		if key >= arr[l] && key <= arr[mid] {
			return search(arr,l,mid-1,key)
		} 
		return search(arr,mid+1,h,key)
	}

	if arr[l]>= arr[mid] && key <= arr[h] {
		//first half not sorted
		return search(arr,mid+1,h,key)
	}
	return search(arr,l,mid-1,key)
}