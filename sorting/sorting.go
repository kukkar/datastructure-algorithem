package sorting

//complexity n^2
func SelectionSort(arr []int)[]int {
	for i:=0;i<len(arr);i++ {
		minIndex := i
		for j:=i+1;j<len(arr);j++ {
			if arr[j] <= arr[minIndex] {
				minIndex = j
			}
		}
		t := arr[minIndex]
		arr[minIndex] =  arr[i]
		arr[i] = t
	}
	return arr
}