package slidingBlocks4x4

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
	boardWithNumbers4x4 [4][4]int
	dir                 []Direction
}

var boardWithLetters4x4 [4][4]string
var endXClass int
var endYClass int

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func convertToInt(tiles [4][4]string, size int) [4][4]int {
	var result [4][4]int
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
func (sl *SlidingBlocksBoard) InitializeByHand(size int, choice string,  endX int, endY int) [4][4]int {
	endXClass=endX
	endYClass=endY

	input := bufio.NewReader(os.Stdin)
	var number int
	if choice == "numbers" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print("Enter tile with position [", i, ",", j, "] : ")
				userNumber, _ := input.ReadString('\n')
				userNumber = strings.TrimRight(userNumber, "\r\n")
				number, _ = strconv.Atoi(userNumber)
				sl.boardWithNumbers4x4[i][j] = number
			}
		}
	} else if choice == "letters" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print("Enter tile with position [", i, ",", j, "] : ")
				letter, _ := input.ReadString('\n')
				letter = strings.TrimRight(letter, "\r\n")
				boardWithLetters4x4[i][j] = letter
			}
		}
		sl.boardWithNumbers4x4 = convertToInt(boardWithLetters4x4, size)
	}
	return sl.boardWithNumbers4x4
}

func (sl *SlidingBlocksBoard) InitializeRandom(size int, endX int, endY int) [4][4]int {
	endXClass=endX
	endYClass=endY

	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(16)[:16]
	
	m:=0;
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sl.boardWithNumbers4x4[i][j] = arr[m]
			m++;
		}
	}
	return sl.boardWithNumbers4x4
}

func (sl *SlidingBlocksBoard) InitializeRandomForString(size int, endX int, endY int) [4][4]string {
	sl.boardWithNumbers4x4=sl.InitializeRandom(size, endX, endY)
	boardWithLetters4x4=sl.convertIntToString(size)
	return boardWithLetters4x4
}

func (sl *SlidingBlocksBoard) PrintMatrix(size int, choice string) {
	if choice=="numbers"{
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print(sl.boardWithNumbers4x4[i][j], " ")
			}
			fmt.Println()
		}
	} else if choice=="letters"{
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print(boardWithLetters4x4[i][j], " ")
			}
			fmt.Println()
		}
	}
}

func (sl *SlidingBlocksBoard) manhattanDistance(size int) int {
	var path int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = sl.boardWithNumbers4x4[i][j]
			if currentTile != 0 {
				var rightRow = (currentTile - 1) / size
				var rightColumn = (currentTile - 1) % size
				path += int(math.Abs(float64(i-rightRow)) + math.Abs(float64(j-rightColumn)))
			}
		}
	}
	return path
}

func (sl *SlidingBlocksBoard) isReachedDestination(tiles [4][4]int, size int) bool {
	if endXClass==0 && endYClass==0 {
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
	}else if endXClass==3 && endYClass==3{
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				var currentTile = tiles[i][j]
				if currentTile != 0 {
					var rightRow = (currentTile-1) / size
					var rightColumn = (currentTile-1) % size
					if i != rightRow || j != rightColumn {
						return false
					}
				}
			}
		}
	}
	return true
}

func swapTiles(copiedBoard [4][4]int, indexFirstRow int, indexFirstColumn int, indexSecondRow int, indexSecondColumn int) [4][4]int {
	var temp = copiedBoard[indexFirstRow][indexFirstColumn]
	copiedBoard[indexFirstRow][indexFirstColumn] = copiedBoard[indexSecondRow][indexSecondColumn]
	copiedBoard[indexSecondRow][indexSecondColumn] = temp
	return copiedBoard
}

func (sl *SlidingBlocksBoard) getMove(currentTile int, size int) Direction {
	startX, startY := sl.findStartPosition(size)

	if startX > 0 && currentTile == sl.boardWithNumbers4x4[startX-1][startY] {
		return DOWN
	} else if startX < size-1 && currentTile == sl.boardWithNumbers4x4[startX+1][startY] {
		return UP
	} else if startY > 0 && currentTile == sl.boardWithNumbers4x4[startX][startY-1] {
		return RIGHT
	} else if startY < size-1 && currentTile == sl.boardWithNumbers4x4[startX][startY+1] {
		return LEFT
	}
	return -1
}

func (sl *SlidingBlocksBoard) returnMove(copiedBoard [4][4]int, step Direction, size int) [4][4]int {
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
			var currentTile = sl.boardWithNumbers4x4[i][j]
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
			if sl.boardWithNumbers4x4[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (sl *SlidingBlocksBoard) visitedMoves(size int) map[Direction][4][4]int {
	var allMoves = []Direction{}

	startX, startY := sl.findStartPosition(size)

	allMoves = sl.getAllMoves(startX, startY, size)[:]

	visited := make(map[Direction][4][4]int)

	for i := 0; i < len(allMoves); i++ {
		var move = allMoves[i]
		puzzleCopy := sl.copyOfPuzzle(size)
		visited[move] = sl.returnMove(puzzleCopy, move, size)
	}
	//	fmt.Println("map:", visited)
	return visited
}

func (sl *SlidingBlocksBoard) copyOfPuzzle(size int) [4][4]int {
	var copyOfMatrix [4][4]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyOfMatrix[i][j] = sl.boardWithNumbers4x4[i][j]
		}
	}
	return copyOfMatrix
}

func (sl *SlidingBlocksBoard) distance(size int, moves []Direction) int {
	return len(moves) + sl.manhattanDistance(size)
}

func (sl *SlidingBlocksBoard) popMinDistance(states map[Direction][4][4]int) map[Direction][4][4]int {
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

func (sl *SlidingBlocksBoard) AStar(size int) [4][4]int {
	var states = make(map[Direction][4][4]int)
	var moves []Direction

	var keys []Direction

	for k, _ := range states {
		keys = append(keys, k)
	}

	for len(states) > 0 {
		var puzzle = sl.popMinDistance(states)
		var puz = puzzle[0]
		if sl.isReachedDestination(puz, size) {
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
			var child = SlidingBlocksBoard{boardWithNumbers4x4: visit, dir: copy_paths}

			var current_distance = sl.distance(size, child.dir)
			fmt.Println(current_distance)
			//states[current_distance]=child;
			//states.push((child, current_distance));
			i++
		}
	}
	return sl.boardWithNumbers4x4
}

func (sl *SlidingBlocksBoard) UserPlay(size int, choice string){
	input := bufio.NewReader(os.Stdin)
	if choice=="numbers"{
		for !sl.isReachedDestination(sl.boardWithNumbers4x4, size){
			fmt.Println("Enter the direction you want to move the blank position (0): ")
			userDirection, _ := input.ReadString('\n')
			userDirection = strings.TrimRight(userDirection, "\r\n")
			direction,_:=strconv.Atoi(userDirection)
			directions:=Direction(direction)
			sl.boardWithNumbers4x4=sl.returnMove(sl.boardWithNumbers4x4, directions, size)
			sl.PrintMatrix(size, choice)
		}
	} else if choice=="letters"{
		for !sl.isReachedDestination(sl.boardWithNumbers4x4, size){
			fmt.Println("Enter the direction you want to move the blank position (0): ")
			userDirection, _ := input.ReadString('\n')
			userDirection = strings.TrimRight(userDirection, "\r\n")
			direction,_:=strconv.Atoi(userDirection)
			directions:=Direction(direction)
			sl.boardWithNumbers4x4=sl.returnMove(sl.boardWithNumbers4x4, directions, size)
			boardWithLetters4x4=sl.convertIntToString(size)
			sl.PrintMatrix(size, choice)
		}
	}
	fmt.Println("Congratulations! You solved the puzzle!")
}

func (sl * SlidingBlocksBoard) convertIntToString(size int) [4][4]string{
	for i:=0; i<size;i++{
		for j:=0;j<size;j++{
			if sl.boardWithNumbers4x4[i][j]!=0{
				boardWithLetters4x4[i][j]=string(sl.boardWithNumbers4x4[i][j]+96)
			}else {
				boardWithLetters4x4[i][j]="0"
			}
		}
	}
	return boardWithLetters4x4
}
