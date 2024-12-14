package main 
import (
	"fmt"
)

// PLAYER setup

type Player struct {
	pType string
}

func createPlayers() []Player {
	players := []Player{
		{pType: "X"},
		{pType: "O"},
	}

	return players
}

func (p Player) displayPlayer() {
	fmt.Println(p.pType)
}

// BOARD setup

type Board struct {
	tiles [][]string
}

func createBoard() Board{
	b := Board{
		tiles: [][]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
	}
	return b
}

func (b Board) displayBoard() {
	for i := 0; i < 3; i++ {
		fmt.Println(b.tiles[i][0], " |", b.tiles[i][1], "|", b.tiles[i][2])
		if (i != 2){	
			fmt.Println("---|---|---")
		}
	}
}

func (b Board) checkMove(row, col int) bool {
	if b.tiles[row][col] == " " {
		return true
	}
	return false
}

func (b Board) checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.tiles[i][j] == " "{
				return false
			}
		}
	}
	return true 
}

func (b *Board) updateBoard(row, col int, value string) {
	if row < len(b.tiles) && col < len(b.tiles[row]) {
		b.tiles[row][col] = value
	}

}

// will check if there is a win on the current snapshot of the board
func (b Board) checkWin() bool{
	
	//checking the rows
	for i := 0; i < 3; i++ {
		if (b.tiles[i][0] == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] && b.tiles[i][0] != " "){
			return true
		}
	}
	
	// checking the columns
	for i := 0; i < 3; i++ {
		if (b.tiles[0][i] == b.tiles[1][i] && b.tiles[2][i] == b.tiles[i][2] && b.tiles[0][i] != " "){
			return true
		}
	}

	// checking the diagonals
	if b.tiles[0][0] == b.tiles[1][1] && b.tiles[1][1] == b.tiles[2][2] && b.tiles[0][0] != " "{
		return true
	}
	if b.tiles[0][2] == b.tiles[1][1] && b.tiles[1][1] == b.tiles[2][0] && b.tiles[2][0] != " "{
		return true
	}

	// otherwise return false
	return false
}

// AI PLAYER setup
type botPlayer struct {
	pType string
	opponentType string
	nextMove []int
	firstMove []int
}

func (p botPlayer) getPlayer() string {
	return p.pType
}

func (p botPlayer) getNextMove() []int {
	return p.nextMove
}

func (p botPlayer) setOpponentType(s string) {
	p.opponentType = s
}

func isBoardClear(b Board, p botPlayer) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Processing %d %d tile.\n", i, j)
			if b.tiles[i][j] != " " {
				p.firstMove = []int{i,j}
				return false
			}
		}
	}
	return true
}

func checkRow(row int, x, y, z, pType, rowType string) []int {
	if rowType == "row"{	
		if x == pType && y == pType {
			return []int{row, 2}
		} else if x == pType && z == pType {
			return []int{row, 1} 
		} else if z == pType && y == pType {
			return []int{row, 0}
		}
	}
	if rowType == "col"{
		if x == pType && y == pType {
			return []int{2, row}
		} else if x == pType && z == pType {
			return []int{1, row} 
		} else if z == pType && y == pType {
			return []int{0, row}
		}
	}
	return nil
}

func isTwoInRow(b Board, p botPlayer) []int {
	// This function finds if there are two in a row or column and 
	// returns the one to stop the win 
	fmt.Println("Checking opponents...")
	for row := 0; row < 3; row++{
		x := b.tiles[row][0]
		y := b.tiles[row][1]
		z := b.tiles[row][2]

		fmt.Printf("Row: %s %s %s \n", x, y, z)

		if x != " " && y != " " && z != " "{
			// do nothing 
		} else {
			// check opponents on board
			check := checkRow(row, x, y, z, p.opponentType, "row")
			fmt.Printf("Move: %v \n", check)
			if check == nil {
				// do nothing
			} else {
				fmt.Println("Returning move")
				return check
			}
		}
	}

	for col := 0; col < 3; col++{
		x := b.tiles[0][col]
		y := b.tiles[1][col]
		z := b.tiles[2][col]
		fmt.Printf("Row: %s %s %s \n", x, y, z)
		
		if x != " " && y != " " && z != " "{
			// do nothing
		} else {
			// check opponents on board
			check := checkRow(col, x, y, z, p.opponentType, "col")
			fmt.Printf("Move: %v \n", check)
			if check == nil {
				// do nothing
			} else {
				return check
			}
		}
	}

	fmt.Println("Checking bot...")

	for row := 0; row < 3; row++{
		x := b.tiles[row][0]
		y := b.tiles[row][1]
		z := b.tiles[row][2]
		fmt.Printf("Row: %s %s %s \n", x, y, z)

		if x != " " && y != " " && z != " "{
			// do nothing 
		} else {
			// check placements on board
			check := checkRow(row, x, y, z, p.pType, "row")
			fmt.Printf("Move: %v \n", check)
			if check == nil {
				// do nothing
			} else {
				return check
			}
		}
	}

	for col := 0; col < 3; col++{
		x := b.tiles[0][col]
		y := b.tiles[1][col]
		z := b.tiles[2][col]
		fmt.Printf("Row: %s %s %s \n", x, y, z)

		if x != " " && y != " " && z != " "{
			// do nothing
		} else {
			// check placements on board
			check := checkRow(col, x, y, z, p.pType, "col")
			fmt.Printf("Move: %v \n", check)
			if check == nil {
				// do nothing
			} else {
				return check
			}
		}
	}

	return nil
}

func calculateNextMove(b Board, p botPlayer, turn int) []int {
	if turn < 2{	
		fmt.Println("Processing...")
		if isBoardClear(b, p) {
			return []int{0,2}
		} else if b.checkMove(0,0){
			return []int{0,0}
		} else if b.checkMove(2,2){
			return []int{2,2}
		} else if b.checkMove(2,0){
			return []int{2,0}
		}
	} 
	
	if turn >= 2{	
		fmt.Println("Processing")
		nextMove := isTwoInRow(b, p)
		if nextMove == nil {
			if b.checkMove(0,0){
				return []int{0,0}
			} else if b.checkMove(0,2){
				return []int{0,2}
			} else if b.checkMove(2,0){
				return []int{2,0}
			} else if b.checkMove(2,2){
				return []int{2,2}
			} else if b.checkMove(0,1){
				return []int{0,1}
			} else if b.checkMove(1,0){
				return []int{1,0}
			} else if b.checkMove(1,2){
				return []int{1,2}
			} else if b.checkMove(2,1){
				return []int{2,1}
			} else if b.checkMove(1,1){
				return []int{1,1}
			}
		} else {
			return nextMove
		}
	}
	
	return nil
}

func main() {
	board := createBoard()
	// players := createPlayers()
	realPlayer := Player{
		pType: "O",
	}

	bot := botPlayer{
		pType: "X",
		opponentType: "O",
		nextMove: []int{1, 1}, 
		firstMove: []int{2, 2},
	}

	PLAYER_X := bot
	PLAYER_O := realPlayer

	// var currentPlayer = PLAYER_X
	// game loop 
	turn := 0

	for {	
		board.displayBoard()
		
		var row, col int

		if turn % 2 == 0{
			// real player move
			for {
				// get move 
				fmt.Printf("Player %s, enter your move (row and column-ex: '1 2'): ", PLAYER_O.pType)
				fmt.Scanln(&row, &col)
				if board.checkMove(row, col) {
					break
				} else {
					fmt.Println("Invalid move. Try again!")
				}
			}
			
			// make the move
			board.updateBoard(row, col, PLAYER_O.pType)
		} else {
			fmt.Println("Bot: Calculating turn...")
			botMove := calculateNextMove(board, PLAYER_X, turn)
			fmt.Println(botMove)
			if board.checkMove(botMove[0], botMove[1]) {
				fmt.Println("Bot Move OK")
			}


			// make the move
			board.updateBoard(botMove[0], botMove[1], PLAYER_X.pType)
		}
		

		// display board
		board.displayBoard()

		// check for the win
		if board.checkWin() {
			fmt.Println("WIN")
			break
		} else {
			fmt.Println("NO WIN")
		}
		if board.checkDraw() {
			fmt.Println("DRAW")
			break
		}

		// switch players
		// if currentPlayer == PLAYER_X {
		// 	currentPlayer = PLAYER_O
		// } else {
		// 	currentPlayer = PLAYER_X
		// }
		turn++
	}
}


/*
	col
row  0   1   2
0 	   |   |
	---|---|---
1	 x |   |
	---|---|---
2	   |   |

*/