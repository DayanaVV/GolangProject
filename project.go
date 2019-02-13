package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/DayanaVV/GolangProject/pkg/boardLetters4x4"
	"github.com/DayanaVV/GolangProject/pkg/boardNumbers3x3"
	"github.com/DayanaVV/GolangProject/pkg/boardNumbers4x4"
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

	
	var board boardNumbers3x3.SlidingBlocksBoard;

	if sizeOfBoard == "3" || sizeOfBoard == "4" {

		if (sizeOfBoard == "3" && choice == "numbers") || (sizeOfBoard == "3" && choice == "letters"){
			board.InitializeByHand(size, choice);
			board.PrintMatrix(size);
		} else if sizeOfBoard == "4" && choice == "numbers" {
			boardNumbers4x4.Initialize(size)
			boardNumbers4x4.PrintMatrix(size)
		} else if sizeOfBoard == "4" && choice == "letters" {
			boardLetters4x4.Initialize(size)
			boardLetters4x4.PrintMatrix(size)
		} else {
			fmt.Print("Incorect input. Choose between numbers and letters")
			return
		}
	} else {
		fmt.Print("Incorect number. Choose between 3 and 4")
		return
	}
}
