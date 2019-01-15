package boardLetters4x4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

var boardWithLetters4x4 [4][4]string

func Initialize(size int) [4][4]string {
	input := bufio.NewReader(os.Stdin)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("Enter tile with position [", i, ",", j, "] : ")
			letter, _ := input.ReadString('\n')
			letter = strings.TrimRight(letter, "\r\n")

			boardWithLetters4x4[i][j] = letter
		}
	}
	return boardWithLetters4x4
}

func PrintMatrix(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(boardWithLetters4x4[i][j], " ")
		}
		fmt.Println()
	}
}

/*
func manhattanDistance(size int) int {
}

func isReachedDestionation(size int) bool {
}

func swapTiles(indexFirstRow int, indexFirstColumn int, indexSecondRow int, indexSecondColumn int) {
}

func findStartPosition(size int) (int, int) {
}

func getMove(currentTile int, size int) Direction {
}

func returnMove(step Direction, size int) [3][3]string {
}

func getAllMoves(startX int, startY int, size int) []Direction {
}

func copyMoves(path []Direction) []Direction {
}

func visistedMoves(size int, startX int, startY int) map[[3][3]int]Direction {
}

func copyOfPuzzle(size int) [3][3]string {
}

   func popMinDistance(map[[3][3]int]int) [3][3]string {

	func AStar(size int) {

}*/
