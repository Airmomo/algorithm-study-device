package sorting

//MergeSort 归并排序
func MergeSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	mergeProcess(&arr, 0, len(arr)-1)
	return arr
}

func mergeProcess(arr *[]int, L, R int) {
	if L == R {
		return
	}
	//mid := L + ((R - L) >> 1)
	mid := (R + L) >> 1
	mergeProcess(arr, L, mid)
	mergeProcess(arr, mid+1, R)
	merge(arr, L, R, mid)
}

func merge(arr *[]int, L, R, mid int) {
	help := make([]int, R - L + 1)
	i := 0
	p1 := L
	p2 := mid + 1
	for p1 <= mid && p2 <= R {
		if (*arr)[p1] < (*arr)[p2] {
			help[i] = (*arr)[p1]
			p1++
		} else {
			help[i] = (*arr)[p2]
			p2++
		}
		i++
	}
	for p1 <= mid {
		help[i] = (*arr)[p1]
		i++
		p1++
	}
	for p2 <= R {
		help[i] = (*arr)[p2]
		i++
		p2++
	}
	for i := 0;i<len(help);i++{
		(*arr)[L + i] = help[i]
	}
}
