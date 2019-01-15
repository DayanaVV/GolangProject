package boardNumbers4x4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

var boardWithNumbers4x4 [4][4]int

func Initialize(size int) [4][4]int {
	input := bufio.NewReader(os.Stdin)
	var number int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("Enter tile with position [", i, ",", j, "] : ")
			userNumber, _ := input.ReadString('\n')
			userNumber = strings.TrimRight(userNumber, "\r\n")
			number, _ = strconv.Atoi(userNumber)
			boardWithNumbers4x4[i][j] = number
		}
	}
	return boardWithNumbers4x4
}

func PrintMatrix(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(boardWithNumbers4x4[i][j], " ")
		}
		fmt.Println()
	}
}

func manhattanDistance(size int) int {
	var path int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = boardWithNumbers4x4[i][j]
			if currentTile != 0 {
				var rightRow = (currentTile - 1) / size
				var rightColumn = (currentTile - 1) % size
				path += int(math.Abs(float64(i-rightRow)) + math.Abs(float64(j-rightColumn)))
			}
		}
	}
	return path
}

func isReachedDestionation(size int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = boardWithNumbers4x4[i][j]
			if currentTile != 0 {
				var rightRow = (currentTile - 1) / size
				var rightColumn = (currentTile - 1) % size
				if i != rightRow || j != rightColumn {
					return false
				}
			}
		}
	}
	return true
}

func swapTiles(indexFirstRow int, indexFirstColumn int, indexSecondRow int, indexSecondColumn int) {
	var temp = boardWithNumbers4x4[indexFirstRow][indexFirstColumn]
	boardWithNumbers4x4[indexFirstRow][indexFirstColumn] = boardWithNumbers4x4[indexSecondRow][indexSecondColumn]
	boardWithNumbers4x4[indexSecondRow][indexSecondColumn] = temp
}

func findStartPosition(size int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if boardWithNumbers4x4[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func getMove(currentTile int, size int) Direction {
	startX, startY := findStartPosition(size)

	if startX > 0 && currentTile == boardWithNumbers4x4[startX-1][startY] {
		return DOWN
	} else if startX < size-1 && currentTile == boardWithNumbers4x4[startX+1][startY] {
		return UP
	} else if startY > 0 && currentTile == boardWithNumbers4x4[startX][startY-1] {
		return RIGHT
	} else if startY < size-1 && currentTile == boardWithNumbers4x4[startX][startY+1] {
		return LEFT
	}
	return -1
}

func returnMove(step Direction, size int) [4][4]int {
	startX, startY := findStartPosition(size)

	switch step {
	case LEFT:
		swapTiles(startX, startY, startX, startY+1)
		break
	case RIGHT:
		swapTiles(startX, startY, startX, startY-1)
		break
	case UP:
		swapTiles(startX, startY, startX+1, startY)
		break
	case DOWN:
		swapTiles(startX, startY, startX-1, startY)
		break
	}
	return boardWithNumbers4x4
}

func getAllMoves(startX int, startY int, size int) []Direction {
	var allMoves = []Direction{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = boardWithNumbers4x4[i][j]
			var dir = getMove(currentTile, size)
			if dir != -1 {
				allMoves = append(allMoves, dir)
			}
		}
	}

	return allMoves
}

func copyMoves(path []Direction) []Direction {
	var copy = []Direction{}
	copy = path[:]
	return copy
}

/*
func visistedMoves(size int, startX int, startY int) map[[4][4]int]Direction {
}

func copyOfPuzzle(size int) [4][4]int {
	var copyOfMatrix [4][4]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyOfMatrix[i][j] = boardWithNumbers4x4[i][j]
		}
	}
	return copyOfMatrix
}

   func popMinDistance(map[[4][4]int]int) [4][4]int {

	func AStar(size int) {

}*/
