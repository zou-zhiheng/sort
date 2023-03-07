package DSA

import (
	"errors"
	"fmt"
)

//冒泡
func bubbleSort(array []int) {

	flag := false
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				flag = true
				array[j], array[j+1] = array[j+1], array[j]
			}
		}

		if !flag {
			break
		} else {
			flag = false
		}
	}
}

//选择排序
//每次从无序表选择最大值（最小值）放入有序表
func selectSort(array []int) {

	var enumSort func(array []int) error
	enumSort = func(array []int) error {

		if len(array) == 0 {
			return errors.New("空数组")
		}
		if len(array) == 1 {
			return nil
		}

		val := array[0]
		index := 0
		for i := 1; i < len(array); i++ {
			//查找index值
			if val > array[i] {
				val = array[i]
				index = i
			}

		}

		array[0], array[index] = array[index], array[0]

		return nil
	}

	for i := range array {
		_ = enumSort(array[i:])
	}

}

//把n个待排元素看成为一个有序表
//每次将无序表第一个元素插入到有序表
//(大->小）
func insertSort(arr []int) {
	if arr == nil {
		fmt.Println("空数组")
		return
	} else if len(arr) == 1 {
		return
	}
	//有序表
	lastIndex := 0 //有序表最后一个元素的下标
	//无序表依次拿出元素插入有序表
	for i := lastIndex + 1; i < len(arr); i++ {
		//在有序表中查找插入点
		for j := 0; j < lastIndex+1; j++ {
			if arr[i] >= arr[j] { //找到插入点,大于最大值
				//记录下要插入的值，有序表元素后移
				val := arr[i]
				//有序表元素后移
				for k := lastIndex; k >= j; k-- {
					arr[k+1] = arr[k]
				}
				arr[j] = val
				lastIndex++
				break
			}
		}
	}

}

//希尔排序
//把记录按下标的一定增量分组，对每组使用直接插入排序算法排序，
//随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，
//整个文件恰好被分为一组算法终止
//交换式实现
func shellSortChange(arr []int) {
	if arr == nil {
		fmt.Println("空数组")
		return
	} else if len(arr) == 1 {
		return
	}
	//增量
	gap := len(arr) / 2
	for ; gap > 0; gap /= 2 {
		//每组依次比较,第i-gap+1组开始
		for i := gap; i < len(arr); i++ { //组数len(arr)-gap
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] < arr[j+gap] {
					arr[j], arr[j+gap] = arr[j+gap], arr[j]
				}
			}
		}
	}

}

//移位法
func shellSort(arr []int) {
	if arr == nil {
		fmt.Println("空数组")
		return
	} else if len(arr) == 1 {
		return
	}
	//增量
	gap := len(arr) / 2
	for ; gap > 0; gap /= 2 {
		//每组依次比较,第i-gap+1组开始
		for i := gap; i < len(arr); i++ { //组数len(arr)-gap
			j := i
			temp := arr[j]
			if arr[j] > arr[j-gap] {
				for ; j-gap >= 0 && temp > arr[j-gap]; {
					//移动,按步长后移
					arr[j] = arr[j-gap]
					j -= gap
				}
				//找到temp的插入位置
				arr[j] = temp
			}
		}
	}
}

//快速排序,迭代法实现
func quickSort(arr []int, left, right int) {

	pivot := arr[(left+right)/2]
	l, r := left, right
	for ; l < r; {
		//在pivot的左边一直找，找到大于等于pivot的值，才退出
		for ; arr[l] < pivot; {
			l += 1
		}
		//在pivot的左边一直找，找到大于等于pivot的值，才退出
		for ; arr[r] > pivot; {
			r -= 1
		}
		//如果l>=r说明pivot的左边全部是小于等于pivot值，
		//右边全部是大于等于pivot值
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		//如果交换完后，发现arr[l]==pivot值 ,r--,前移
		if arr[l] == pivot {
			r--
		}
		//如果交换完后，发现arr[r]==pivot值 ,l++,后移
		if arr[r] == pivot {
			l++
		}

	}

	if l == r { //退出循环
		l++
		r--
	}

	//向左递归
	if left < r {
		quickSort(arr, left, r)
	}
	//向右递归
	if right > l {
		quickSort(arr, l, right)
	}

}

//归并
//大->小
func merge(arr []int, left, mid, right int, temp []int) {
	i := left    //初始化，左边有序序列的初始索引
	j := mid + 1 //初始化j，右边有序序列的初始索引
	t := 0       //指向temp数组的当期索引

	//temp数组填充
	for ; i <= mid && j <= right; {
		if arr[i] < arr[j] {
			temp[t] = arr[j]
			j++
		} else {
			temp[t] = arr[i]
			i++
		}
		t++
	}

	//剩下一个左有序或右有序数组
	if i <= mid { //剩下右有序数组
		for ; i <= mid; i++ {
			temp[t] = arr[i]
			t++
		}
	} else { //剩下左有序数组
		for ; j <= right; j++ {
			temp[t] = arr[j]
			t++
		}
	}

	//将temp数组的值拷贝到arr
	t = 0
	for i = left; i <= right; {
		arr[i] = temp[t]
		i++
		t++
	}

}
func mergeSort(arr []int, left, right int, temp []int) {

	if left < right {
		mid := (left + right) / 2
		//向左递归分解
		mergeSort(arr, left, mid, temp)
		//向右递归分解
		mergeSort(arr, mid+1, right, temp)
		//合并
		merge(arr, left, mid, right, temp)
	}

}

func radixSort(arr []int) {
	//定义一个二维数组，表示10个桶，每个桶就是一个一纬数组
	//1、二维数组包含长度为10的一纬数组
	//2、为了防止在放入数的时候，数据溢出，则每个一维数组（桶），大小为len(arr)
	//3、名明确，基数排序是使用空间换时间的经典算法
	bucket := make([][10]int, len(arr))

	//为了记录每个桶中，实际存放了多少个数据，定义一个一维数组来记录各个桶的每次放入的数据个数
	//可以这里理解
	//比如：bucketElementCounts:=make([]int,10)
	bucketElementCounts := make([]int, 10)

	//第一轮（针对每个元素的个位进行排序处理）
	for j := 0; j < len(arr); j++ {
		//取出每个元素的个位的值
		digitOfElement := arr[j] % 10
		//放到对应的桶中
		bucket[digitOfElement][bucketElementCounts[digitOfElement]] = arr[j]
		bucketElementCounts[digitOfElement]++
	}

	//按照这个桶的顺序（一维数组的下标依次取出数据，放入原来数组）
	index := 0
	for k := 0; k < len(bucketElementCounts); k++ {
		//如果桶中，有数据 ，我们才放入到原数组
		if bucketElementCounts[k] != 0 {
			//循环该桶即第K个桶(即第k个一维数组)，放入
			for l := 0; l < bucketElementCounts[k]; l++ {
				//取出元素放入到arr
				arr[index+1]=bucket[k][l]
			}
		}
	}

}

func SortDemo() {
	test := []int{-9, 78, 0, 23, 24, -567, 70, 85, -1, 900, 456}
	//bubbleSort(test)
	//selectSort(test)
	//insertSort(test)
	//shellSort(test)
	quickSort(test, 0, len(test)-1)
	//temp := make([]int, len(test))
	//mergeSort(test, 0,  len(test)-1, temp)

	fmt.Println(test)
}
