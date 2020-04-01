package sorting


//SelectSort 选择排序
func SelectSort(arr []int)[]int{
	if arr == nil || len(arr) < 2{
		return arr
	}
	for i := 0; i < len(arr); i++{
		for j := i + 1; j < len(arr);j++{
			if arr[j] < arr[i]{
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	return arr
}