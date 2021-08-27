package main
import "fmt"
func checkSortedArray(arr []int){
	sortedArray := true
	for i:=0; i<=len(arr)-1; i++{
		for j:=0; j<len(arr)-1-i; j++{
			if arr[j]> arr[j+1]{
				sortedArray = false
				break
			}
		}
	}
	if sortedArray{
		fmt.Println("给定数组已排序。")
	} else {
		fmt.Println("给定数组未排序。")
	}
}

func main(){
	checkSortedArray([]int{1, 3, 5, 6, 7, 8})
	checkSortedArray([]int{1, 3, 5, 9, 4, 2})
	checkSortedArray([]int{9, 7, 4, 2, 1, -1})
}