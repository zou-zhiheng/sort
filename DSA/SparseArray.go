package DSA

import "fmt"

func SparseArray() {
	chessArr1 := [11][11]int{}
	chessArr1[1][2] = 1
	chessArr1[2][3] = 2

	//有效数字个数
	var sum int
	//获取有效数字个数
	for i := 0; i < len(chessArr1); i++ {
		for j := 0; j < len(chessArr1[i]); j++ {
			if chessArr1[i][j] != 0 {
				sum++
			}
		}
	}

	chessArr2 := make([][3]int, sum+1)
	//稀疏数组赋值
	chessArr2[0][0] = 11
	chessArr2[0][1] = 11
	chessArr2[0][2] = sum
	count := 0
	for i := 0; i < len(chessArr1) && count <= sum; i++ {
		for j := 0; j < len(chessArr1[i]); j++ {
			if chessArr1[i][j] != 0 {
				count++
				chessArr2[count][0] = i
				chessArr2[count][1] = j
				chessArr2[count][2] = chessArr1[i][j]
			}

		}
	}
	for i:=0;i<len(chessArr2);i++{
		for j:=0;j<3;j++{
			fmt.Print(chessArr2[i][j]," ")
		}
		fmt.Println()
	}


	//稀疏数组转置为二维数组
	//获取行列
	chessArr3 := make([][11]int, chessArr2[0][1])
	for i := 1; i <= chessArr2[0][2]; i++ {
		chessArr3[chessArr2[i][0]][chessArr2[i][1]]=chessArr2[i][2]
	}

	//原始数组打印
	for i:=0;i<len(chessArr3);i++{
		for j:=0;j<len(chessArr3[i]);j++{
			fmt.Print(chessArr3[i][j]," ")
		}
		fmt.Println()
	}

}
