package main

import (
	"fmt"
	"github.com/airmomo/algorithm/sorting"
	"github.com/airmomo/algorithm/util"
	"time"
)

func main(){
	fmt.Println("---- Test START ----")

	startTime := time.Now().UnixNano()

	//TestSort 排序类问题对数器
	err := util.TestSort(50000, sorting.HeapSort)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Test Past")
	}

	endTime := time.Now().UnixNano()

	//seconds:= float64((endTime - startTime) / 1e9)
	milliSeconds:= float64((endTime - startTime) / 1e6)
	//nanoSeconds:= float64(endTime - startTime)

	fmt.Printf("---- Test END ---- Wrost Time : %.2f ms",milliSeconds)
}
