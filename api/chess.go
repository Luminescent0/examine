package api

import (
	"fmt"
)

var Board [10][9]int

const (
	RedGen     = iota + 1  //帅
	RedSold                //仕
	RedMin                 //相
	RedHorse               //马
	RedVeh                 //车
	RedArm                 //兵
	RedGun                 //炮
	BlackGen   = iota + 11 //帅
	BlackSold              //仕
	BlackMin               //象
	BlackHorse             //马
	BlackVeh               //车
	BlackArm               //兵
	BlackGun               //炮

)

func Operate(choose int, client *Client, cases int) {
	switch choose {
	case 1:
		HorseMove(client, cases)
	case 2:
		MinMove(client, cases)
	case 3:
		SoldMove(client, cases)
	}

}
func InitBoard() {
	Board[0][0] = BlackVeh
	Board[0][8] = BlackVeh
	Board[0][1] = BlackHorse
	Board[0][7] = BlackHorse
	Board[0][2] = BlackMin
	Board[0][6] = BlackMin
	Board[0][3] = BlackSold
	Board[0][5] = BlackSold
	Board[0][4] = BlackGen
	Board[2][1] = BlackGun
	Board[2][7] = BlackGun
	for i := 0; i < 9; i = i + 2 {
		Board[3][i] = BlackArm
	}
	Board[9][0] = RedVeh
	Board[9][8] = RedVeh
	Board[9][1] = RedHorse
	Board[9][7] = RedHorse
	Board[9][2] = RedMin
	Board[9][6] = RedMin
	Board[9][3] = RedSold
	Board[9][5] = RedSold
	Board[9][4] = RedGen
	Board[7][1] = RedGun
	Board[7][7] = RedGun
	for i := 0; i < 9; i = i + 2 {
		Board[6][i] = RedArm
	}
	fmt.Println(Board)
}

func FindChess(Chess int) (i, j int) {
	for i, v := range Board {
		for j, v2 := range v {
			if v2 == Chess {
				fmt.Printf("board[%v][%v]", i, j)
				return i, j
			}
		}
	}
	fmt.Println(i, j)
	return i, j
}
func HorseMove(client *Client, cases int) {
	if client.Typ == 0 {
		i, j := FindChess(RedHorse)
		Board[i][j] = 0
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
		if Board[i][j] == 0 || Board[i][j] > 8 { //防止吃掉自己的子（。
			Board[i][j] = RedHorse
		} else {
			return
		}
	} else {
		i, j := FindChess(BlackHorse)
		fmt.Printf("board[%v][%v]", i, j)
		Board[i][j] = 0
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
		if Board[i][j] < 8 {
			Board[i][j] = BlackHorse
		}
	}
}

func MinMove(client *Client, cases int) {
	if client.Typ == 0 {
		i, j := FindChess(RedMin)
		fmt.Printf("board[%v][%v]", i, j)
		Board[i][j] = 0
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
		if Board[i][j] == 0 || Board[i][j] > 8 {
			Board[i][j] = RedMin
		} else {
			return
		}
	} else {
		i, j := FindChess(BlackMin)
		fmt.Printf("board[%v][%v]", i, j)
		Board[i][j] = 0
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
		if Board[i][j] < 8 {
			Board[i][j] = BlackMin
		} else {
			return
		}
	}
}
func SoldMove(client *Client, cases int) {
	if client.Typ == 0 {
		i, j := FindChess(RedSold)
		switch cases {
		case 1:
			i = i + 1
			j = j + 1
		case 2:
			i = i - 1
			j = j + 1
		case 3:
			i = i - 1
			j = j - 1
		case 4:
			i = i + 1
			j = j - 1
		}
		if i > 10 || i < 8 || j < 3 || j > 5 {
			return
		} else if Board[i][j] < 8 {
			return
		}
		Board[i][j] = RedSold
	} else {
		i, j := FindChess(BlackSold)
		switch cases {
		case 1:
			i = i + 1
			j = j + 1
		case 2:
			i = i - 1
			j = j + 1
		case 3:
			i = i - 1
			j = j - 1
		case 4:
			i = i + 1
			j = j - 1
		}
		if i > 2 || i < 0 || j < 3 || j > 5 {
			return
		} else if Board[i][j] > 8 {
			return
		}
		Board[i][j] = BlackSold
	}
}
