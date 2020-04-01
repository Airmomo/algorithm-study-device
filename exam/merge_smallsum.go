package main

import (
	"errors"
	"fmt"
	"github.com/airmomo/algorithm/util"
	"math/rand"
	"time"
)

//求一个数组的小和问题
//在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。

func main() {
	fmt.Println("---- Test START ----")

	startTime := time.Now().UnixNano()

	//testMergeSmallSum 数组小和对数器
	err := testMergeSmallSum(50000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Test Past")
	}

	endTime := time.Now().UnixNano()

	//seconds:= float64((endTime - startTime) / 1e9)
	milliSeconds := float64((endTime - startTime) / 1e6)
	//nanoSeconds:= float64(endTime - startTime)

	fmt.Printf("---- Test END ---- Wrost Time : %.2f ms", milliSeconds)

}

//rightSum 绝对正确的结果
func rightSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				sum += arr[i]
			}
		}
	}
	return sum
}

//testMergeSmallSum 对数器
func testMergeSmallSum(testTime int) error {
	rand.Seed(time.Now().UnixNano())
	pastTime := 0
	for testTime > pastTime {
		testArr := util.RandomArr()
		//testArr := []int{4,3,1,4}

		copyArr := make([]int,len(testArr))
		copy(copyArr,testArr)
		testRet := MergeSort(copyArr)

		copyArr2 := make([]int,len(testArr))
		copy(copyArr2,testArr)
		rightRet := rightSum(copyArr2)

		if testRet != rightRet {
			msg := fmt.Sprintf("Test Past times : %d \n---- Test ERROR ----\nTest Eg : \n %v\nTest Ret : \n %v\n---- Test MustBe ----\nRight Ret : \n %v", pastTime, testArr, testRet, rightRet)
			return errors.New(msg)
		}
		pastTime++
	}
	return nil
}

//MergeSort 归并排序
func MergeSort(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return sortProcess(&arr, 0, len(arr)-1)
}

func sortProcess(arr *[]int, L, R int) int {
	if L == R {
		return 0
	}
	//mid := L + ((R - L) >> 1)
	mid := (R + L) >> 1
	return sortProcess(arr, L, mid) + sortProcess(arr, mid+1, R) + merge(arr, L, R, mid)
}

func merge(arr *[]int, L, R, mid int) int {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid + 1
	ret := 0
	for p1 <= mid && p2 <= R {
		if (*arr)[p1] < (*arr)[p2] {
			//小和规则：左边的值 < 当前的值
			//复合规则则进行一次累加
			ret += (*arr)[p1] * (R - p2 + 1)
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
	return ret
}
