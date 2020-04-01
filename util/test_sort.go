package util

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

//maxSize 设置数组随机产生的最大长度
const maxSize  = 100

//randomArr 生成一个随机长度的随机数组
func RandomArr()[]int{
	size := rand.Intn(maxSize) + 1
 	arr := make([]int,size)
 	for i,_ := range arr{
 		arr[i] = rand.Intn(size) - (size >> 1)
	}
	return arr
}

//rightSort 返回一个绝对正确的有序数组
func rightSort(arr []int) []int{
	sort.Ints(arr)
	return arr
}

//equalArr 校验测试数组和正确数组是否相同
func equalArr(test,right []int)bool{
	//if len(test) != len(right){
	//	return false
	//}
	//for i,_ := range right{
	//	if right[i] != test[i]{
	//		return false
	//	}
	//}
 	return reflect.DeepEqual(test,right)
}

//TestSort(testTime 测试次数,testFunc 被测试方法) 排序对数器
func TestSort(testTime int,testFunc func([]int)[]int) error{
	rand.Seed(time.Now().UnixNano())
	pastTime := 0
	for testTime > pastTime{
		testArr := RandomArr()
		copyArr := make([]int,len(testArr))
		copy(copyArr,testArr)
		testRet := testFunc(copyArr)
		copyArr2 := make([]int,len(testArr))
		copy(copyArr2,testArr)
		rightRet := rightSort(copyArr2)
		if !equalArr(testRet,rightRet){
			msg := fmt.Sprintf("Test Past times : %d \n---- Test ERROR ----\nTest Eg : \n %v\nTest Ret : \n %v\n---- Test MustBe ----\nRight Ret : \n %v",pastTime,testArr,testRet,rightRet)
			return errors.New(msg)
		}
		pastTime++
	}
	return nil
}