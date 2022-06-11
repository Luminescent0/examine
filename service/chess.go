package service

import (
	"examine/api"
	"fmt"
)

var board [10][9]int
const (
	RedGen = 1 //帅
	RedSold = 2//仕
	RedMin = 3//相
	RedHorse = 4//马
	RedVeh = 5//车
	RedArm = 6//兵
	RedGun= 7//炮
	BlackGen = 11 //帅
	BlackSold = 12 //仕
	BlackMin = 13 //象
	BlackHorse = 14 //马
	BlackVeh = 15 //车
	BlackArm = 16//兵
	BlackGun = 17//炮

)

func Operate(choose int,client *api.Client,cases int) {
	switch choose {
	case 1:
		HorseMove(client,cases)

	case 2:
		MinMove(client,cases)
	case 3:
		SoldMove(client,cases)
	}

}
func InitBoard() {
	board[0][0]= BlackVeh
	board[0][8]= BlackVeh
	board[0][1]= BlackHorse
	board[0][7]= BlackHorse
	board[0][2]= BlackMin
	board[0][6]= BlackMin
	board[0][3]= BlackSold
	board[0][5]= BlackSold
	board[0][4]= BlackGen
	board[2][1]= BlackGun
	board[2][7]= BlackGun
	for i:=0;i<9;i=i+2{
		board[3][i]= BlackArm
	}
	board[9][0]= RedVeh
	board[9][8]= RedVeh
	board[9][1]= RedHorse
	board[9][7]= RedHorse
	board[9][2]= RedMin
	board[9][6]= RedMin
	board[9][3]= RedSold
	board[9][5]= RedSold
	board[9][4]= RedGen
	board[7][1]= RedGun
	board[7][7]= RedGun
	for i:=0;i<9;i=i+2{
		board[6][i]= RedArm
	}
	fmt.Println(board)
}

func FindChess(Chess int) (i,j int) {
	for i,v := range board {
		for j,v2 := range v{
			if v2 == Chess{
				fmt.Printf("board[%v][%v]",i,j)
				return i,j
			}
		}
	}
	fmt.Println(i,j)
	return i,j
}
func HorseMove(client *api.Client,cases int) {
	if client.Typ == 0 {
		i,j:= FindChess(RedHorse)
		board[i][j] = 0
		switch cases { //这里爆格的话终端会报错，但是不能返回给用户（。
		case 1:
			i = i + 1
			j = j + 2
		case 2:
			i = i + 2
			j = j + 1
		case 3:
			i = i - 1
			j = j + 2
		case 4:
			i = i - 2
			j = j + 1
		case 5:
			i = i - 2
			j = j - 1
		case 6:
			i = i - 1
			j = j - 2
		case 7:
			i = i + 1
			j = j - 2
		case 8:
			i = i + 2
			j = j + 1
		}
		if board[i][j] == 0 || board[i][j] > 8 { //防止吃掉自己的子（。
			board[i][j] = RedHorse
		} else {
			return
		}
	} else {
		i,j:= FindChess(BlackHorse)
		fmt.Printf("board[%v][%v]", i, j)
		board[i][j] = 0
		switch cases{ //这里爆格的话终端会报错，但是不能返回给用户（。
		case 1:
			i = i + 1
			j = j + 2
		case 2:
			i = i + 2
			j = j + 1
		case 3:
			i = i - 1
			j = j + 2
		case 4:
			i = i - 2
			j = j + 1
		case 5:
			i = i - 2
			j = j - 1
		case 6:
			i = i - 1
			j = j - 2
		case 7:
			i = i + 1
			j = j - 2
		case 8:
			i = i + 2
			j = j + 1
		}
		if board[i][j] < 8 {
			board[i][j] = BlackHorse
		} else {
			return
		}
	}
}

func MinMove(client *api.Client,cases int) {
	if client.Typ == 0 {
		i,j:= FindChess(RedMin)
		fmt.Printf("board[%v][%v]",i,j)
		board[i][j]=0
		switch cases {
		case 1:
			i = i + 2
			j = j + 2
		case 2:
			i = i - 2
			j = j + 2
		case 3:
			i = i + 2
			j = j - 2
		case 4:
			i = i - 2
			j = j - 2
		}
		if board[i][j]==0|| board[i][j]>8 {
			board[i][j]= RedMin
		}else {
			return
		}
	}else {
		i,j:= FindChess(BlackMin)
		fmt.Printf("board[%v][%v]",i,j)
		board[i][j]=0
		switch cases {
		case 1:
			i=i+2
			j=j+2
		case 2:
			i=i-2
			j=j+2
		case 3:
			i=i+2
			j=j-2
		case 4:
			i=i-2
			j=j-2
		}
		if board[i][j]<8 {
			board[i][j]= BlackMin
		}else {
			return
		}
	}
}
func SoldMove(client *api.Client,cases int) {
	if client.Typ==0 {
		i,j:= FindChess(RedSold)
		switch cases {
		case 1:
			i=i+1
			j=j+1
		case 2:
			i=i-1
			j=j+1
		case 3:
			i=i-1
			j=j-1
		case 4:
			i=i+1
			j=j-1
		}
		if i>10||i<8||j<3||j>5 {
			return
		}else if board[i][j] < 8 {
			return
		}
		board[i][j]= RedSold
	}else{
		i,j:= FindChess(BlackSold)
		switch cases {
		case 1:
			i=i+1
			j=j+1
		case 2:
			i=i-1
			j=j+1
		case 3:
			i=i-1
			j=j-1
		case 4:
			i=i+1
			j=j-1
		}
		if i>2||i<0||j<3||j>5 {
			return
		}else if board[i][j] > 8 {
			return
		}
		board[i][j]= BlackSold
	}
}
