package piscine

import (
	"fmt"
	"os"
	
	
	"strconv"
	
)



func main() {
	arg := os.Args[1:]
	bord := [][]int{
		//    0  1  2  
		/*0*/{1, 0, 0, 0, 0, 0, 0, 0, 0},
		/*1*/{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// i = 0 ".96.4...1" // j == 1
	// i = 1 "1...6.1.4"
	// append in table
	for i := range arg { // i = 0  //0
		for j, c := range arg[i]/* */ {//0
			if c != '.' {
				nbr, _ := strconv.Atoi(string(c)) // string --> int
				
				/*9 int*/
				
				bord[i][j] = nbr
			} else if c == '.' {
				continue
			}
		}
	}
	//
	if Check(bord) { // Check(bord) == true  || !Check(borad) == false
		//clearScreen()
		for i := 0; i < 9; i++ { // i == 0
			for j := 0; j < 9; j++ {
				if j == 8 {
					fmt.Print(bord[i][j]) //borad[0][0]
					continue
				}
				fmt.Print(bord[i][j], " ")
			}
			fmt.Print("\n")
		}
	}
}

func Check(Board [][]int) bool { // true false
	for i := range Board {
		for j := range Board[i] {
			if Board[i][j] == 0 {
				// 1 2 3 4 5 6 7 8
				for k := 1; k <= 9; k++ { // 1 
					//                1
					//   true         1
					if isValid(Board, k, i, j) {

						Board[i][j] = k



						if Check(Board) {
							return true
						}
						Board[i][j] = 0
					}
				}
				return false

			}
		}
	}
	return true
}
//                          1        0     3
func isValid(Board [][]int, num int, raw, col int) bool {
	for j := range Board {
		//    raw
		if num == Board[raw][j] {
			return false
			// colone
		} else if num == Board[j][col] {
			return false
		}
	}
	// caré // i == 5 | j == 4
	// raw = 5 - 5 % 3 = 5 - 2 = 3
	// col = 4 - 4 % 3 = 4 - 1 = 3
	// i == 3 | j == 3

	// i == 8 | j == 6
	// raw = 8 - 8 % 3 = 8 - 2 = 6
	// col = 6 - 6 % 3 = 6 - 0 = 6
	// i == 6 | j == 6
	star_raw := raw - raw%3
	star_col := col - col%3
	//       i = 6            i < 9 
	for i := star_raw; i < star_raw+3; i++ {
		for j := star_col; j < star_col+3; j++ {
			if Board[i][j] == num {
				return false
			}
		}
	}

	return true
}

//      0  1  2   3  4  5   6  7  8
//   0  1  9  6 | 0  4  0 | 0  0  1
//   1  1  2  2 | 0  6  0 | 0  0  4
//   2  5  0  4 | 8  1  0 | 3  9  0
//      ---------------------------- 3 9 6 2 4 5 7 8 1
//   3  0  0  7 | 9  5  0 | 0  4  3	 1 7 8 3 6 9 5 2 4
//   4  0  3  0 | 0  8  0 | 0  0  0  5 2 4 8 1 7 3 9 6
//   5  4  0  5 | 0  2  3 | 0  1  8  2 8 7 9 5 1 6 4 3
//      ---------------------------- 9 3 1 4 8 6 2 7 5
//   6  0  1  0 | 6  3  0 | 0  5  9  4 6 5 7 2 3 9 1 8
//   7  0  5  9 | 0  7  0 | 8  3  0  7 1 2 6 3 8 4 5 9
//   8  0  0  3 | 5  9  0 | 0  0  7  6 5 9 1 7 4 8 3 2
//                                   8 4 3 5 9 2 1 6 7
