package main

import "fmt"

func main() {
	arr1 := []int{2, 4, 3, 7, 9, 8}
	arr2 := []int{0, 1, 5, 5}
	mergeSort(soort(arr1), soort(arr2))
}
//sorts a given splice with O(n)
func soort(arr []int)[]int{
	for i:=0;i<len(arr)-1;i++ {
		if arr[i]>arr[i+1]{
			arr[i],arr[i+1]=arr[i+1],arr[i]
		}
	}
	return arr
}
//merges two splice while comparing its elements one by one with complexity of O(n)
func mergeSort(arr1 []int, arr2 []int){
	i,j:=0,0
	arr3:=make([]int,len(arr1)+len(arr2))

	for k:=0;k<len(arr3);k++ {
		if i>=len(arr1) {
			arr3[k]=arr2[j]
			j++
			continue
		}else if j>=len(arr2){
			arr3[k]=arr1[i]
			i++
			continue
		}
		if arr1[i]<arr2[j] {
			arr3[k]=arr1[i]
			i++
		}else{
			arr3[k]=arr2[j]
			j++
		}
	}
	fmt.Println(arr3)
}
