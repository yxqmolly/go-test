package main

import "fmt"

func find(arr *[6]int, leftIndex int, rightIndex int, findval int){
	if leftIndex > rightIndex  {
		fmt.Println("找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if((*arr)[middle] > findval){
		find(arr, leftIndex, middle - 1, findval)
	}else if((*arr)[middle] == findval){
		fmt.Println("找到了")
	}else{
		find(arr, middle + 1, rightIndex, findval)
	}
}
/**
 * @Description: 二分查找（指针）
 * @param arr
 * @param leftIndex
 * @param rightIndex
 * @param findval
 */
func findPtr(arr [6]*int, leftIndex int, rightIndex int, findval int){
	if  leftIndex > rightIndex{
		fmt.Println("找不到")
	}
	middle := (leftIndex + rightIndex) / 2
	if(*arr[middle] > findval){
		findPtr(arr,leftIndex, middle - 1, findval)
	}else if(*arr[middle] == findval){
		fmt.Println("找到了")
	}else {
		findPtr(arr, middle + 1, rightIndex, findval)
	}
}

func main(){
	arr := [6]int{1,3,5,7,9,11}
	find(&arr, 0, len(arr) - 1, 7)
	var arrPtr [6]*int
	for i := range arr{
		arrPtr[i] = &arr[i]
	}
	findPtr(arrPtr, 0, len(arrPtr) - 1, 7)
}