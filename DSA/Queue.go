package DSA

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	maxSize int   //表示数组最大容量
	front   int   //队列头，第一个元素的位置
	rear    int   //指向最后一个元素的后一个位置，预留出一个位置做约定
	arr     []int //该数组用于存放数据，模拟队列
}

func newArrayQueue(maxSize int) ArrayQueue {
	return ArrayQueue{maxSize: maxSize, front: 0, rear: 0, arr: make([]int, maxSize)}
}

//判断队列是否满
func (r *ArrayQueue) isFull() bool {
	return (r.rear+1)%r.maxSize == r.front
}

//判断队列是否为空
func (r *ArrayQueue) isEmpty() bool {
	return r.front == r.rear
}

//加入队列
func (r *ArrayQueue) addQueue(n int) {
	if r.isFull() {
		fmt.Println("队列已满，不能加入新元素")
		return
	}
	r.arr[r.rear] = n
	r.rear = (r.rear + 1) % r.maxSize

}

//出队列
func (r *ArrayQueue) getQueue() (int, error) {
	if r.isEmpty() {
		return 0, errors.New("队列为空")
	}

	res := r.arr[r.front]
	r.front = (r.front + 1) % r.maxSize

	return res, nil
}

//队列展示
func (r *ArrayQueue) showQueue() {
	if r.isEmpty() {
		fmt.Println("空队列")
		return
	}
	dataLen := (r.rear + r.maxSize - r.front) % r.maxSize
	for i := r.front; i < dataLen; i++ {
		fmt.Printf("a[%d]=%d\n", i%r.maxSize, r.arr[i%r.maxSize])
	}
}

//获取有效数据的长度
func (r *ArrayQueue) size() int {
	return (r.rear + r.maxSize - r.front) % r.maxSize
}

//展示头部front数据
func (r *ArrayQueue) headQueue() (int, error) {
	if r.isEmpty() {
		return 0, errors.New("空队列")
	}
	return r.arr[r.front], nil
}

func ArrayQueueDemo() {
	test := newArrayQueue(100)
	test.addQueue(1)
	test.addQueue(2)
	test.showQueue()
	front, err := test.headQueue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(test.front)
		fmt.Println(front)
	}

	fmt.Println("size", test.size())
	t1,err:=test.getQueue()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(t1,test.size())
	t1,err=test.getQueue()
	if err!=nil{
		fmt.Println(err)
	}
	t1,err=test.getQueue()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(t1,test.size())

}
