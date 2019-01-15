package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"project"
)

type Direction int

var boardWithNumbers3x3 [3][3]int

var boardWithNumbers4x4 [4][4]int

//var boardWithLetters3x3 [3][3]char

//var boardWithLetters3x3 [4][4]char

const (
	LEFT Direction = iota
	RIGHT
	UP    
	DOWN 
)

func main() {

	input := bufio.NewReader(os.Stdin)

	fmt.Print("Enter size of the board you want to play: ")
	sizeOfBoard, _ := input.ReadString('\n')
	sizeOfBoard = strings.TrimRight(sizeOfBoard, "\r\n")

	fmt.Print("Enter whether you want to play with numbers or letters: ")
	choice, _ := input.ReadString('\n')
	choice = strings.TrimRight(choice, "\r\n")

	var size int
	size, _ = strconv.Atoi(sizeOfBoard)

	if sizeOfBoard == "3" || sizeOfBoard == "4" {
		if sizeOfBoard == "3" && choice == "numbers" {
			initialize(size)
			//	manhattanDistance(size)
			//	swapTiles(0, 0, 1, 1)
			//startX, startY := findStartPosition(size)
			//fmt.Println(startX, startY)
			var path = []Direction{3, 1, 0, 2}
			fmt.Println(path)

		}
		printMatrix(size)
		/*} else if sizeOfBoard == "4" && choice == "numbers" {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Print("Enter numbers: ")
				userNumber, _ := input.ReadString('\n')
				userNumber = strings.TrimRight(userNumber, "\r\n")
				number, _ = strconv.Atoi(userNumber)
				boardWithNumbers4x4[i][j] = number
			}
		}
		fmt.Println(boardWithNumbers4x4)*/

	} else {
		fmt.Print("Incorect number. Choose between 3 and 4")
		return
	}

}

func initialize(size int) [3][3]int {
	input := bufio.NewReader(os.Stdin)
	var number int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("Enter tile with position [", i, ",", j, "] : ")
			userNumber, _ := input.ReadString('\n')
			userNumber = strings.TrimRight(userNumber, "\r\n")
			number, _ = strconv.Atoi(userNumber)
			boardWithNumbers3x3[i][j] = number
		}
	}
	return boardWithNumbers3x3
}

func printMatrix(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(boardWithNumbers3x3[i][j], " ")
		}
		fmt.Println()
	}
}

func manhattanDistance(size int) int {
	var path int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = boardWithNumbers3x3[i][j]
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
			var currentTile = boardWithNumbers3x3[i][j]
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
	var temp = boardWithNumbers3x3[indexFirstRow][indexFirstColumn]
	boardWithNumbers3x3[indexFirstRow][indexFirstColumn] = boardWithNumbers3x3[indexSecondRow][indexSecondColumn]
	boardWithNumbers3x3[indexSecondRow][indexSecondColumn] = temp
	fmt.Println(boardWithNumbers3x3)
}

func findStartPosition(size int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if boardWithNumbers3x3[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func getMove(currentTile int, size int) Direction {
	startX, startY := findStartPosition(size)
	
	if (startX > 0 && currentTile == boardWithNumbers3x3[startX-1][startY]) {
		return DOWN
	} else if (startX < size-1 && currentTile == boardWithNumbers3x3[startX+1][startY]) {
		return UP
	} else if (startY > 0 && currentTile == boardWithNumbers3x3[startX][startY-1]) {
		return RIGHT
	} else if (startY < size-1 && currentTile == boardWithNumbers3x3[startX][startY+1]) {
		return LEFT
	}
	return -1
}

func returnMove(step Direction, size int) [3][3]int {
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
	return boardWithNumbers3x3
}

func getAllMoves(startX int, startY int, size int) []Direction {
	var allMoves = []Direction{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var currentTile = boardWithNumbers3x3[i][j]
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

func visistedMoves(size int, startX int, startY int) map[[3][3]int]Direction {
	var allMoves = []Direction{}
	allMoves=getAllMoves(startX, startY, size)[:]

	visited:=make(map[[3][3]int]Direction)

	//List<Pair<int[][], Direction>> visited = new ArrayList<>()

	for i := 0; i < len(allMoves); i++ {
		var move = allMoves[i]
		//puzzleCopy := copyOfPuzzle(size)
		visited[returnMove(move, size)]=move

	//	visited.add(new Pair(returnMove(move, row, column, copy), move));
	}
	return visited
}
func copyOfPuzzle(size int) [3][3]int {
	var copyOfMatrix [3][3]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			copyOfMatrix[i][j] = boardWithNumbers3x3[i][j];
		}
	}
	return copyOfMatrix;
}
/*
   func popMinDistance(map[[3][3]int]int) [3][3]int {
		var minIndex = 0
		states:=make(map[[3][3]int]int)
		minDistance := states[0]
	
		for i := 0; i < len(states); i++ {
			var p map[[3][3]int]int
			p = states[i]
            dist := p.getValue()

            if dist < minDistance {
                minIndex = i
                minDistance = dist
            }
		}
		states=delete(states, minIndex)
        return states
    }
/*
	func aStar(size int) {
       // List<Pair<Puzzle, Integer>> states = new ArrayList<>();
        //List<Direction> moves = new ArrayList<>();

        //finds the zero position's indexes
        startX, startY = findStartPosition(size)
        

        //create initial puzzle
     //   Puzzle puzzle = new Puzzle(matrix, moves, row);
     //   states.add(new Pair(puzzle, 0));

        //pop minimal distance - path + manhatan
        while (states.size() > 0) {
            Puzzle p = popMinDistance(states);

            if (p.isSolved()) {
                System.out.println(p.path.size()+ "\n");
                for (int i = 0; i < p.path.size(); i++) {
                    String str=p.path.get(i).name();
                    System.out.println(str.toLowerCase() + "\n");
                }
               
                return;
            }

            //visit all available moves and calculate distance and add to state visited and pop to final condition
            List<Pair<int[][], Direction>> visited = visistedMoves(p.matrix, row, column, startX, startY);

            for (int i = 0; i < visited.size(); i++) {
                Pair<int[][], Direction> visit = visited.get(i);
                List<Direction> copyPaths = new ArrayList<>();

                //copy previos path and add last Direction
                copyPaths.addAll(p.path);
                copyPaths.add(visit.getValue());

                //crete a new puzzle with children nodes of the puzzle
                Puzzle child = new Puzzle(visit.getKey(), copyPaths, p.dimension);
                states.add(new Pair(child, child.distance()));
            }
        }
    }*/
