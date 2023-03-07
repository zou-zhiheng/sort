package DSA

import (
	"fmt"
	"math"
)

func MazeDemo() {
	mazeMap := make([][]int, 0)
	for i := 0; i < 8; i++ {
		mazeMap = append(mazeMap, []int{7: 0})
	}

	//使用1表示墙
	//上下全部置为1
	for i := 0; i < 7; i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}
	//左右全部置为1
	for i := 0; i < 8; i++ {
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}

	//设置挡板
	mazeMap[3][1] = 1
	mazeMap[3][2] = 1

	//地图打印
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j], " ")
		}
		fmt.Println()
	}

	//使用递归回溯来给小球找路
	//说明
	//1 map表示地图
	//2 i,j表示从地图的哪个位置出发
	//设置策略(down->right->up->left),如果该点走不通再回溯
	setWay(mazeMap, 1, 1)
	//地图打印
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j], " ")
		}
		fmt.Println()
	}

}

func setWay(mazeMap [][]int, i, j int) bool {

	if mazeMap[6][5] == 2 { //找到出口
		return true
	} else {
		if mazeMap[i][j] == 0 { //当前点没有走过
			//按照策略(down->right->up->left)走
			mazeMap[i][j] = 2 //假定该点可以走通
			if setWay(mazeMap, i+1, j) {
				return true
			} else if setWay(mazeMap, i, j+1) {
				return true
			} else if setWay(mazeMap, i-1, j) {
				return true
			} else if setWay(mazeMap, i, j-1) {
				return true
			} else {
				//说明该点走不通是，该点置为3
				mazeMap[i][j] = 3
				return false
			}
		} else { //可能是1，2，3
			return false
		}
	}

}

//八皇后
func Queen8() {
	max := 8
	array := make([]int, max)
	count:=0

	//判断是否冲突
	var isConflict = func(n int) bool {

		for i := 0; i < n; i++ {
			if array[i] == array[n] || math.Abs(float64(n-i)) == math.Abs(float64(array[i]-array[n])) {
				return false
			}
		}

		return true
	}

	//放入皇后
	var check func(n int)
	check = func(n int) {
		if n == max {
			count++
			return
		}

		//依次放入皇后，并判断是否冲突
		for i := 0; i < max; i++ {
			//把皇后n放到该行的第i列
			array[n] = i
			//判断是否冲突
			if isConflict(n) { //不冲突
				check(n + 1)
			}
			//如果冲突皇后移向下一列
		}
	}
	check(0)
	fmt.Println(count)

}

func RecursionDemo() {
	//MazeDemo()
	Queen8()
}
