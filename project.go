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
	var endX int
	var endY int
	if size==3{
		fmt.Print("Enter which end position you want 0,0 or 2,2: ")
		endPosition, _ := input.ReadString('\n')
		endPosition = strings.TrimRight(endPosition, "\r\n")
		end:=strings.Split(endPosition, ",")
		endX,_=strconv.Atoi(end[0])
		endY,_=strconv.Atoi(end[1])
	} else if size==4{
		fmt.Print("Enter which end position you want 0,0 or 3,3: ")
		endPosition, _ := input.ReadString('\n')
		endPosition = strings.TrimRight(endPosition, "\r\n")
		end:=strings.Split(endPosition, ",")
		endX,_=strconv.Atoi(end[0])
		endY,_=strconv.Atoi(end[1])
	}


	

	
	var board slidingBlocks3x3.SlidingBlocksBoard;
	var board2 slidingBlocks4x4.SlidingBlocksBoard;
	var tiles =[4][4]int{{0,1,2,3}, {4,5,6,7}, {8,9,10,11},{12,13,14,15}}
	//var tiles =[3][3]int{{0,1,2}, {3,4,5}, {6,7,8}}
	if sizeOfBoard == "3" || sizeOfBoard == "4" {

		if (sizeOfBoard == "3" && choice == "numbers") || (sizeOfBoard == "3" && choice == "letters"){
			board.InitializeByHand(size, choice, endX, endY);
			board.PrintMatrix(size);
		//	fmt.Println(board.IsReachedDestination(tiles,size));
		} else if (sizeOfBoard == "4" && choice == "numbers") || (sizeOfBoard == "4" && choice == "letters") {
			board2.InitializeByHand(size,choice, endX, endY)
			board2.PrintMatrix(size)
			fmt.Println(board2.IsReachedDestination(tiles,size));
		} else {
			fmt.Print("Incorect input. Choose between numbers and letters")
			return
		}
	} else {
		fmt.Print("Incorect number. Choose between 3 and 4")
		return
	}
}
