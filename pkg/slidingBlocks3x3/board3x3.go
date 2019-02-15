package slidingBlocks3x3

import (
	"bufio"
	"container/list"
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
var endXClass int = 2
var endYClass int = 2
var undoList = list.New()
var redoList = list.New()

func (sl *SlidingBlocksBoard) New(board [3][3]int) [3][3]int {
	sl.boardWithNumbers3x3 = board
	return sl.boardWithNumbers3x3
}

//initialize board by user
func (sl *SlidingBlocksBoard) InitializeByHand(size int, choice string, endX int, endY int) [3][3]int {
	endXClass = endX
	endYClass = endY

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

	if sl.checkForDuplicates(size) || !sl.isInRange(size) {
		panic("Should not have duplicate values")
	}
	return sl.boardWithNumbers3x3
}

//initialize board random
func (sl *SlidingBlocksBoard) InitializeRandom(size int, endX int, endY int) [3][3]int {
	endXClass = endX
	endYClass = endY

	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(9)[:9]
	m := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sl.boardWithNumbers3x3[i][j] = arr[m]
			m++
		}
	}

	return sl.boardWithNumbers3x3

}

//initialize board with string
func (sl *SlidingBlocksBoard) InitializeRandomForString(size int, endX int, endY int) [3][3]string {
	sl.boardWithNumbers3x3 = sl.InitializeRandom(size, endX, endY)
	boardWithLetters3x3 = sl.convertIntToString(size)
	return boardWithLetters3x3
}

//function to print board
func (sl *SlidingBlocksBoard) PrintMatrix(size int, choice string) {
	if choice == "numbers" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print(sl.boardWithNumbers3x3[i][j], " ")
			}
			fmt.Println()
		}
	} else if choice == "letters" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print(boardWithLetters3x3[i][j], " ")
			}
			fmt.Println()
		}
	}
}

//calculate manhattan distance of the board
func (sl *SlidingBlocksBoard) ManhattanDistance(size int) int {
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

//check whether the board is solved
func (sl *SlidingBlocksBoard) IsReachedDestination(tiles [3][3]int, size int) bool {

	if endXClass == 0 && endYClass == 0 {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				var currentTile = tiles[i][j]
				if currentTile != 0 {
					var rightRow = (currentTile) / size
					var rightColumn = (currentTile) % size
					if i != rightRow || j != rightColumn {
						return false
					}
				}
			}
		}
	} else if endXClass == 2 && endYClass == 2 {
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
	}
	return true
}

//swap two values in the board
func (sl *SlidingBlocksBoard) SwapTiles(copiedBoard [3][3]int, indexFirstRow int, indexFirstColumn int, indexSecondRow int, indexSecondColumn int) [3][3]int {
	var temp = copiedBoard[indexFirstRow][indexFirstColumn]
	copiedBoard[indexFirstRow][indexFirstColumn] = copiedBoard[indexSecondRow][indexSecondColumn]
	copiedBoard[indexSecondRow][indexSecondColumn] = temp
	return copiedBoard
}

//return enum (direction) according to the current tile of the board
func (sl *SlidingBlocksBoard) GetMove(currentTile int, size int) Direction {
	startX, startY := sl.FindStartPosition(size)

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

//function returning the board after swapping two tiles according to the direction
func (sl *SlidingBlocksBoard) ReturnMove(copiedBoard [3][3]int, step Direction, size int) [3][3]int {
	startX, startY := sl.FindStartPosition(size)

	switch step {
	case LEFT:
		copiedBoard = sl.SwapTiles(copiedBoard, startX, startY, startX, startY+1)
		break
	case RIGHT:
		copiedBoard = sl.SwapTiles(copiedBoard, startX, startY, startX, startY-1)
		break
	case UP:
		copiedBoard = sl.SwapTiles(copiedBoard, startX, startY, startX+1, startY)
		break
	case DOWN:
		copiedBoard = sl.SwapTiles(copiedBoard, startX, startY, startX-1, startY)
		break
	}

	return copiedBoard
}

//return all possible moves for the current board
func (sl *SlidingBlocksBoard) GetAllMoves(size int) []Direction {
	var allMoves = []Direction{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = sl.boardWithNumbers3x3[i][j]
			var dir = sl.GetMove(currentTile, size)
			if dir != -1 {
				allMoves = append(allMoves, dir)
			}
		}
	}

	return allMoves
}

//finding the coordinates of the blank position
func (sl *SlidingBlocksBoard) FindStartPosition(size int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if sl.boardWithNumbers3x3[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

//return all visited and available moves
func (sl *SlidingBlocksBoard) VisitedMoves(size int) map[Direction][3][3]int {
	var allMoves = []Direction{}

	allMoves = sl.GetAllMoves(size)[:]

	visited := make(map[Direction][3][3]int)

	for i := 0; i < len(allMoves); i++ {
		var move = allMoves[i]
		puzzleCopy := sl.copyOfPuzzle(sl)
		visited[move] = sl.ReturnMove(puzzleCopy, move, size)
	}
	//	fmt.Println("map:", visited)
	return visited
}

//copy of the game
func (sl *SlidingBlocksBoard) copyOfPuzzle(game *SlidingBlocksBoard) [3][3]int {
	sl.boardWithNumbers3x3 = game.boardWithNumbers3x3
	return sl.boardWithNumbers3x3
}

//heuristic function for the A* algorithm
func (sl *SlidingBlocksBoard) distance(size int, moves []Direction) int {
	return len(moves) + sl.ManhattanDistance(size)
}

//gets the game with minimal distance
func (sl *SlidingBlocksBoard) popMinDistance(states map[int]SlidingBlocksBoard) map[int]SlidingBlocksBoard {
	var minIndex int = 0
	var minDistance int
	var keys []int

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
	var states = make(map[int]SlidingBlocksBoard)
	var moves []Direction

	var keys []int

	for k, _ := range states {
		keys = append(keys, k)
	}

	for len(states) > 0 {
		var puzzle = sl.popMinDistance(states)
		var puz = puzzle[0]
		if sl.IsReachedDestination(puz.boardWithNumbers3x3, size) {
			for i := 0; i < len(moves); i++ {
				fmt.Print(moves[i])
			}
			fmt.Println()
			return puz.boardWithNumbers3x3
		}

		//visit all available moves and calculate distance and add to state visited and pop to final condition
		var visited = sl.VisitedMoves(size)
		i := 0
		var copy_paths []Direction
		for k, _ := range visited {
			var visit = visited[k]

			//copy previos path and add last Direction
			copy_paths[i] = Direction(keys[i])
			fmt.Println("copy", copy_paths)

			//crete a new board with children nodes of the puzzle
			var child = SlidingBlocksBoard{boardWithNumbers3x3: visit, dir: copy_paths}

			var current_distance = sl.distance(size, child.dir)
			fmt.Println(current_distance)
			states[current_distance] = child
			//states.push(current_distance, child);
			i++
		}
	}
	return sl.boardWithNumbers3x3
}

//function for user
func (sl *SlidingBlocksBoard) UserPlay(size int, choice string) {
	input := bufio.NewReader(os.Stdin)
	undoList.Init()
	redoList.Init()

	if choice == "numbers" {
		for !sl.IsReachedDestination(sl.boardWithNumbers3x3, size) {
			fmt.Println("Enter the direction you want to move the blank position (0): ")
			userDirection, _ := input.ReadString('\n')
			userDirection = strings.TrimRight(userDirection, "\r\n")

			if userDirection == "undo" {
				sl.boardWithNumbers3x3 = undo()
				sl.PrintMatrix(size, choice)
			} else if userDirection == "redo" {
				sl.boardWithNumbers3x3 = redo()
				sl.PrintMatrix(size, choice)
			} else {
				direction, _ := strconv.Atoi(userDirection)
				directions := Direction(direction)
				sl.boardWithNumbers3x3 = sl.ReturnMove(sl.boardWithNumbers3x3, directions, size)
				undoList.PushBack(sl.boardWithNumbers3x3)
				redoList.PushFront(sl.boardWithNumbers3x3)
				sl.PrintMatrix(size, choice)
			}
		}
	} else if choice == "letters" {
		for !sl.IsReachedDestination(sl.boardWithNumbers3x3, size) {
			fmt.Println("Enter the direction you want to move the blank position (0): ")
			userDirection, _ := input.ReadString('\n')
			userDirection = strings.TrimRight(userDirection, "\r\n")
			if userDirection == "undo" {
				sl.boardWithNumbers3x3 = undo()
				boardWithLetters3x3 = sl.convertIntToString(size)
				sl.PrintMatrix(size, choice)
			} else if userDirection == "redo" {
				sl.boardWithNumbers3x3 = redo()
				boardWithLetters3x3 = sl.convertIntToString(size)
				sl.PrintMatrix(size, choice)
			} else {
				direction, _ := strconv.Atoi(userDirection)
				directions := Direction(direction)
				sl.boardWithNumbers3x3 = sl.ReturnMove(sl.boardWithNumbers3x3, directions, size)
				undoList.PushBack(sl.boardWithNumbers3x3)
				redoList.PushFront(sl.boardWithNumbers3x3)
				boardWithLetters3x3 = sl.convertIntToString(size)
				sl.PrintMatrix(size, choice)
			}
		}
	}
	fmt.Println("Congratulations! You solved the puzzle!")
}

//converting int board to a string one
func (sl *SlidingBlocksBoard) convertIntToString(size int) [3][3]string {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if sl.boardWithNumbers3x3[i][j] != 0 {
				boardWithLetters3x3[i][j] = string(sl.boardWithNumbers3x3[i][j] + 96)
			} else {
				boardWithLetters3x3[i][j] = "0"
			}
		}
	}
	return boardWithLetters3x3
}

//function for converting string matrix to int one
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

//check whether the user has used equal values for initializing the board
func (sl *SlidingBlocksBoard) checkForDuplicates(size int) bool {
	var result [9]int
	m := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[m] = sl.boardWithNumbers3x3[i][j]
			m++
		}
	}

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result)-1; j++ {
			if result[i] == result[j] {
				return true
			}
		}
	}

	return false
}

//function to check whether user input is in range
func (sl *SlidingBlocksBoard) isInRange(size int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !(sl.boardWithNumbers3x3[i][j] >= 0 && sl.boardWithNumbers3x3[i][j] <= 8) {
				return false
			}
		}
	}
	return true
}

//function for using the undo command
func undo() [3][3]int {
	last := undoList.Front()
	undoList.Remove(last)
	return last.Value.([3][3]int)
}

//function for using the redo command
func redo() [3][3]int {
	last := redoList.Front()
	redoList.Remove(last)
	return last.Value.([3][3]int)
}
