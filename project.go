package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/DayanaVV/GolangProject/pkg/slidingBlocks3x3"
	"github.com/DayanaVV/GolangProject/pkg/slidingBlocks4x4"
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

	
	var board slidingBlocks3x3.SlidingBlocksBoard;
	var board2 slidingBlocks4x4.SlidingBlocksBoard;

	if sizeOfBoard == "3" || sizeOfBoard == "4" {

		if (sizeOfBoard == "3" && choice == "numbers") || (sizeOfBoard == "3" && choice == "letters"){
			board.InitializeByHand(size, choice);
			board.PrintMatrix(size);
		} else if (sizeOfBoard == "4" && choice == "numbers") || (sizeOfBoard == "4" && choice == "letters") {
			board2.InitializeByHand(size,choice)
			board2.PrintMatrix(size)
		} else {
			fmt.Print("Incorect input. Choose between numbers and letters")
			return
		}
	} else {
		fmt.Print("Incorect number. Choose between 3 and 4")
		return
	}
}
