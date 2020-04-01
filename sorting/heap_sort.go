package sorting

//HeapSort 堆排序
func HeapSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		arr = heapInsert(arr, i)
	}
	for heapSize := len(arr) - 1; heapSize > 0; heapSize-- {
		//最后把堆顶和最后一个数值交换
		arr[heapSize], arr[0] = arr[0], arr[heapSize]
		arr = heapiFy(arr, 0, heapSize)
		//heapSize--重新划分构建堆的范围
		//如此循环每次划分的数组最后的位置为最大值
	}
	//循环完成从小到大的排序
	return arr
}

//heapInsert 建一个大项堆
func heapInsert(arr []int, index int) []int {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
	return arr
}

//heapiFy 维护一个大项堆
func heapiFy(arr []int, index, heapSize int) []int {
	left := 2*index + 1
	for left < heapSize {
		var largest int
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] < arr[index] {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*index + 1
	}
	return arr
}
