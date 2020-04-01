
# Algorithm

Algorithm: Algorithm Test project

Algorithm 是本人用于学习数据结构的项目，提供用于学习的数据结构代码。

项目和代码的完整性将随着本人的学习进度进行更新和补充，欢迎勘误和优化。

数据结构通用对数器的存在可方便的用于自主学习和验证编写的方法是否正确。

本项目基于Go语言进行代码编写，代码格式没有特殊性，可借鉴并用于其他编程语言的学习。

## main.go 及其特色介绍
- **main 打印的信息（举例）**
 	-  ---- Test START ----
 	- 通过：Test Past
 	- 不通过：打印错误信息Error
 	 	-  ---- Test ERROR ----
 		- Test Eg :
 		 	-  [0 -2 1 -2 -2 -3]（对数器提供的测试用例）
 		- Test Ret :
 		 	-  [-2 -3 -2 -2 0 1]（错误的测试结果）
 		- ---- Test MustBe ----
 		- Right Ret :
 			-  [-3 -2 -2 -2 0 1]（正确的结果）
 	-  ---- Test END ---- Wrost Time : %.2f ms（程序(测试)执行的耗时）

- **对数器使用**
 	- 修改此行代码，选用Util层下需要的对数器方法（以Test开头命名）
 	 	- 函数参数 
 	 	 	- testTime 测试次数
 	 	 	- testFunc 被测试函数（测试排序类的函数则选用sorting目录下的测试函数）
 	- （后续对数器增多会改成自动检测测试类型模式） 
 	- 举例
 	 	- TestSort( testTime 测试次数 , testFunc 被测试函数 ) 排序对数器
 	 	- 如：err := util.TestSort( 50000 ,  sorting.HeapSort )
 	 - 自定义测试方法时需注意 
 	 	-  testFunc 的函数类型传递时需要严格遵守util层下对数器方法的接收类型，否则无法使用。
 	 	- 如规定 testFunc 为 func( [ ]int ) [ ]int 类型，那么编写的测试用例也需为相同类型。

```go
func main(){
	fmt.Println("---- Test START ----")

	startTime := time.Now().UnixNano()

	//TestSort 排序类问题对数器 (对数器使用)
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
```

## Exam 目录
用于存放数据结构的常见问题、经典问题。
程序中也具有相应的对数器，不过不通过main主程序进行测试。
是独立的问题，具有独立的对数器，直接执行或调用该路径下的方法进行测试和学习。

## Sorting目录
用于存放不同的排序方法，可通过main主程序来进行测试。

- **排序默认为从小到大排序**
 	- 如需其他排序结果需自行修改排序方法的和对数器中绝对正确的代码。

- **已有排序方法**
 	- BubbleSort 冒泡排序
 	- InsertSort 插入排序
 	- SelectSort 选择排序
 	- MergeSort 归并排序
 	- QuickSort (随机)快速排序
 	- HeapSort 堆排序
## Util 目录

Util Package : 用于存放各类数据结构的对数器。

由于Util层下的方法会被exam和相对应的数据结构类型下的测试用例使用到，所以不推荐您修改其方法名称。

您可以通过修改对数器方法下的属性来设置对数器的测试难度。

## 现有的对数器 (将据学习进度持续更新)
对数器会返回一个 error 用于打印测试用例、错误的结果和绝对正确的结果

错误信息 error 的格式和内容可在对数器中找到以“Test”开头命名的方法进行修改。

---
-  test_sort.go 排序对数器
    - 属性
        - maxsize 设置数组随机产生的最大长度
    - 方法
        - RandomArr 生成一个随机长度的随机数组
        - rightSort 返回一个绝对正确的有序数组
        - equalArr 校验测试数组和正确数组是否相同
        - TestSort 排序对数器
---

## 运行

```shell
go run main.go
```
