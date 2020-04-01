package sorting

//BubbleSort 冒泡排序
func BubbleSort(arr []int)[]int{
	if arr == nil || len(arr) < 2{
		return arr
	}
	for end := len(arr)-1 ; end > 0 ;end--{
		for i := 0; i < end ;i++{
			if arr[i] > arr[i+1]{
				arr[i],arr[i+1] = arr[i+1],arr[i]
			}
		}
	}
	return arr
}