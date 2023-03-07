package DSA

import "fmt"

//常用的查找有四种
//顺序（线性）查找
//二分/折半查找
//插值查找
//斐波那契查找

//线性查找
func seqSearch(arr []int, value int) int {
	//线性查找是逐一比对，发现有相同值，就返回下标
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return arr[i]
		}
	}
	return -1
}

//二分查找
//二分查找必须在一个有序数组中查询
//思路
//首先确定该数组的中间下标 mid=(left+right)/2
//然后让需要查找的数findVal和arr[mid]比较
// findVal>arr[mid]，说明要查找的数在mid的右边,因此需要递归的向右查找
// findVal<arr[mid]，说明要查找的数据在mid的左边,因此需要递归的向左查找
// find==arr[mid]，说明找到，返回
//什么时候结束递归
//找到目标值
//递归完整个数组，仍然没有找到findVal
//left>right

//递归写法
func binarySearchRecursion(arr []int, left, right, findVal int) int {
	//没有找到
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	midVal := arr[mid]
	if findVal > midVal { //向右递归
		return binarySearchRecursion(arr, mid+1, right, findVal)
	} else if findVal < midVal {
		return binarySearchRecursion(arr, left, mid-1, findVal)
	} else {
		return mid
	}
}

//迭代写法,返回目标元素在数组中的下标
//在小->大的数组中查找
//左闭右开写法
func binarySearch(arr []int, val int) int {

	l, r := -1, len(arr)-1
	mid := 0
	for ; l <= r; {
		mid = l + (r-l)/2
		if arr[mid] > val { //在mid左边
			r = mid
		} else if arr[mid] < val-1 { //在mid右边
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func SearchDemo() {
	test := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(binarySearch(test, 3))
}
