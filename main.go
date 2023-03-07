package main

import "fmt"

func main() {
	//DSA.SparseArray()
	//DSA.ArrayQueueDemo()
	//DSA.HeroNodeDemo()
	//DSA.ArrayStackDemo()
	//DSA.RecursionDemo()
	//DSA.SortDemo()
	//DSA.SearchDemo()
	test:=[]int{3,3,3,3,3,3}
	fmt.Println(searchRange(test,3))

}

func searchRange(nums []int, target int) []int {
	l,r:=0,len(nums)
	mid:=0
	res:=make([]int,2)
	res[0]=-1
	res[1]=-1
	for l<r {
		mid=(r-l)/2+l
		if nums[mid]>target{//在左边区间
			r=mid
		}else if nums[mid]<target{//在右边区间
			l=mid+1
		}else{//找到目标值
			res[0]=mid
			res[1]=mid
			if mid>0 || mid<len(nums)-1{//左右端点有一个或两个存在
				for l=mid-1;l>=0 && nums[l]==target;l--{
					fmt.Println(l)
				}
				for r=mid+1;r<len(nums) && nums[mid]==target;r++{
					fmt.Println(r)
				}
				res[0]=l+1
				res[1]=r-1
			}
			return res
		}
	}
	return res
}