package sorting

import (
	"math/rand"
)

//QuickSort (随机)快速排序
func QuickSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	return quickProcess(arr, 0, len(arr)-1)
}

func quickProcess(arr []int, L, R int) []int {
	if L < R {
		//随机
		//将数组中随机某一位数与最末尾的数值进行交换,打乱原有样本,使其复杂度具有随机性
		randId := L + rand.Intn(R-L)
		arr[randId], arr[R] = arr[R], arr[randId]
		//快速排序
		p := partitionC(arr, L, R)
		quickProcess(arr, L, p[0]-1)
		quickProcess(arr, p[1]+1, R)
	}
	return arr
}

//partitionC 相比 partition 节约了一个变量空间
func partitionC(arr []int, L, R int) []int {
	less := L - 1
	more := R
	for L < more {
		if arr[L] < arr[R] {
			less++
			arr[L], arr[less] = arr[less], arr[L]
			L++
		} else if arr[L] > arr[R] {
			more--
			arr[L], arr[more] = arr[more], arr[L]
		} else { //arr[L] == arr[flag]
			L++
		}
	}
	arr[R], arr[more] = arr[more], arr[R]
	return []int{less + 1, more}
}

//partition 将大于 flag 的数值放在 flag 的右边，小于的放在左边
func partition(arr []int, L, R int) []int {
	less := L - 1
	more := R + 1
	flag := arr[R]
	for L < more {
		if arr[L] < flag {
			less++
			arr[L], arr[less] = arr[less], arr[L]
			L++
		} else if arr[L] > flag {
			more--
			arr[L], arr[more] = arr[more], arr[L]
		} else { //arr[L] == arr[flag]
			L++
		}
	}
	//arr[R],arr[more] = arr[more],arr[R]
	return []int{less + 1, more - 1}
}
