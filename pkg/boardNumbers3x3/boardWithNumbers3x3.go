package boardNumbers3x3

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Direction int

const (
	LEFT = Direction(iota)
	RIGHT
	UP
	DOWN
)

type SlidingBlocksBoard struct {
	boardWithNumbers3x3 [3][3]int
	dir                 []Direction
}

var boardWithLetters3x3 [3][3]string

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func convertToInt(tiles [3][3]string, size int) [3][3]int {
	var result [3][3]int
	var getAllRunes []rune
	n := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			convertedString := []rune(tiles[i][j])
			getAllRunes = append(getAllRunes, convertedString...)
			if int(getAllRunes[n]-'0') != 0 {
				result[i][j] = int(getAllRunes[n] - '0' - 48)
				n++
			} else {
				result[i][j] = int(getAllRunes[n] - '0')
				n++
			}
		}
	}
	return result
}
func (sl *SlidingBlocksBoard) InitializeByHand(size int, choice string) [3][3]int {
	input := bufio.NewReader(os.Stdin)
	var number int
	if choice == "numbers" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print("Enter tile with position [", i, ",", j, "] : ")
				userNumber, _ := input.ReadString('\n')

				userNumber = strings.TrimRight(userNumber, "\r\n")
				number, _ = strconv.Atoi(userNumber)
				sl.boardWithNumbers3x3[i][j] = number
			}
		}
	} else if choice == "letters" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print("Enter tile with position [", i, ",", j, "] : ")
				letter, _ := input.ReadString('\n')
				letter = strings.TrimRight(letter, "\r\n")

				boardWithLetters3x3[i][j] = letter
			}
		}
		sl.boardWithNumbers3x3 = convertToInt(boardWithLetters3x3, size)
	}

	return sl.boardWithNumbers3x3
}

func (sl *SlidingBlocksBoard) InitializeRandom(size int) [3][3]int {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(9)[:9]
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sl.boardWithNumbers3x3[i][j] = arr[i]
		}
	}
	return sl.boardWithNumbers3x3
}

func (sl *SlidingBlocksBoard) PrintMatrix(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(sl.boardWithNumbers3x3[i][j], " ")
		}
		fmt.Println()
	}
}

func (sl *SlidingBlocksBoard) manhattanDistance(size int) int {
	var path int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = sl.boardWithNumbers3x3[i][j]
			if currentTile != 0 {
				var rightRow = (currentTile - 1) / size
				var rightColumn = (currentTile - 1) % size
				path += int(math.Abs(float64(i-rightRow)) + math.Abs(float64(j-rightColumn)))
			}
		}
	}
	return path
}

func (sl *SlidingBlocksBoard) isReachedDestionation(tiles [3][3]int, size int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = tiles[i][j]
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

func swapTiles(copiedBoard [3][3]int, indexFirstRow int, indexFirstColumn int, indexSecondRow int, indexSecondColumn int) [3][3]int {
	var temp = copiedBoard[indexFirstRow][indexFirstColumn]
	copiedBoard[indexFirstRow][indexFirstColumn] = copiedBoard[indexSecondRow][indexSecondColumn]
	copiedBoard[indexSecondRow][indexSecondColumn] = temp
	return copiedBoard
}

func (sl *SlidingBlocksBoard) getMove(currentTile int, size int) Direction {
	startX, startY := sl.findStartPosition(size)

	if startX > 0 && currentTile == sl.boardWithNumbers3x3[startX-1][startY] {
		return DOWN
	} else if startX < size-1 && currentTile == sl.boardWithNumbers3x3[startX+1][startY] {
		return UP
	} else if startY > 0 && currentTile == sl.boardWithNumbers3x3[startX][startY-1] {
		return RIGHT
	} else if startY < size-1 && currentTile == sl.boardWithNumbers3x3[startX][startY+1] {
		return LEFT
	}
	return -1
}

func (sl *SlidingBlocksBoard) returnMove(copiedBoard [3][3]int, step Direction, size int) [3][3]int {
	startX, startY := sl.findStartPosition(size)

	switch step {
	case LEFT:
		copiedBoard = swapTiles(copiedBoard, startX, startY, startX, startY+1)
		break
	case RIGHT:
		copiedBoard = swapTiles(copiedBoard, startX, startY, startX, startY-1)
		break
	case UP:
		copiedBoard = swapTiles(copiedBoard, startX, startY, startX+1, startY)
		break
	case DOWN:
		copiedBoard = swapTiles(copiedBoard, startX, startY, startX-1, startY)
		break
	}

	return copiedBoard
}

func (sl *SlidingBlocksBoard) getAllMoves(startX int, startY int, size int) []Direction {
	var allMoves = []Direction{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = sl.boardWithNumbers3x3[i][j]
			var dir = sl.getMove(currentTile, size)
			if dir != -1 {
				allMoves = append(allMoves, dir)
			}
		}
	}

	return allMoves
}

func (sl *SlidingBlocksBoard) findStartPosition(size int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if sl.boardWithNumbers3x3[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (sl *SlidingBlocksBoard) visitedMoves(size int) map[Direction][3][3]int {
	var allMoves = []Direction{}

	startX, startY := sl.findStartPosition(size)

	allMoves = sl.getAllMoves(startX, startY, size)[:]

	visited := make(map[Direction][3][3]int)

	for i := 0; i < len(allMoves); i++ {
		var move = allMoves[i]
		puzzleCopy := sl.copyOfPuzzle(size)
		visited[move] = sl.returnMove(puzzleCopy, move, size)
	}
	//	fmt.Println("map:", visited)
	return visited
}

func (sl *SlidingBlocksBoard) copyOfPuzzle(size int) [3][3]int {
	var copyOfMatrix [3][3]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyOfMatrix[i][j] = sl.boardWithNumbers3x3[i][j]
		}
	}
	return copyOfMatrix
}

func (sl *SlidingBlocksBoard) distance(size int, moves []Direction) int {
	return len(moves) + sl.manhattanDistance(size)
}

func (sl *SlidingBlocksBoard) popMinDistance(states map[Direction][3][3]int) map[Direction][3][3]int {
	var minIndex Direction = 0
	var minDistance Direction
	var keys []Direction

	for k, _ := range states {
		keys = append(keys, k)
	}
	fmt.Println("keys", keys)
	minDistance = keys[0]
	for k, _ := range states {
		dist := k

		if dist < minDistance {
			minIndex = k
			minDistance = dist
		}
	}

	delete(states, minIndex)
	fmt.Println(states)
	return states

}

func (sl *SlidingBlocksBoard) AStar(size int) [3][3]int {
	var states = make(map[Direction][3][3]int)
	var moves []Direction

	var keys []Direction

	for k, _ := range states {
		keys = append(keys, k)
	}

	for len(states) > 0 {
		var puzzle = sl.popMinDistance(states)
		var puz = puzzle[0]
		if sl.isReachedDestionation(puz, size) {
			for i := 0; i < len(moves); i++ {
				fmt.Print(moves[i])
			}
			fmt.Println()
			return puz
		}

		var visited = sl.visitedMoves(size)
		i := 0
		var copy_paths []Direction
		for k, _ := range visited {
			var visit = visited[k]

			//copy previos path and add last Direction
			copy_paths[i] = keys[i]
			fmt.Println("copy", copy_paths)

			//crete a new puzzle with children nodes of the puzzle
			var child = SlidingBlocksBoard{boardWithNumbers3x3: visit, dir: copy_paths}

			var current_distance = sl.distance(size, child.dir)
			fmt.Println(current_distance)
			//states[current_distance]=child;
			//states.push((child, current_distance));
			i++
		}
	}
	return sl.boardWithNumbers3x3
}
