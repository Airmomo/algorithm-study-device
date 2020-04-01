package sorting

//InsertSort 插入排序
func InsertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j+1] < arr[j]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
	return arr
}
