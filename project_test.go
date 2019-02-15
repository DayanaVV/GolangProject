package main

import (
	"github.com/DayanaVV/SlidingBlocksGoGame/pkg/slidingBlocks3x3"
	"github.com/DayanaVV/SlidingBlocksGoGame/pkg/slidingBlocks4x4"
	"reflect"
	"testing"
)

func TestManhattan(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	total := board.ManhattanDistance(3)
	if total != 2 {
		t.Errorf("ManhattanDistance from board3x3 was incorrect, got: %d, want: %d.", total, 2)
	}

	board.New([3][3]int{{5, 7, 2}, {1, 3, 6}, {8, 4, 0}})
	total = board.ManhattanDistance(3)
	if total != 12 {
		t.Errorf("ManhattanDistance from board3x3 was incorrect, got: %d, want: %d.", total, 12)
	}

	board2.New([4][4]int{{9, 5, 8, 6}, {3, 2, 1, 0}, {4, 15, 14, 7}, {11, 13, 12, 10}})
	total = board2.ManhattanDistance(4)
	if total != 36 {
		t.Errorf("ManhattanDistance from board4x4 was incorrect, got: %d, want: %d.", total, 36)
	}
	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}})
	total = board2.ManhattanDistance(4)
	if total != 3 {
		t.Errorf("ManhattanDistance from board4x4 was incorrect, got: %d, want: %d.", total, 3)
	}

}

func TestIsReachedDestination(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	total := board.IsReachedDestination([3][3]int{{5, 7, 2}, {1, 3, 6}, {8, 4, 0}}, 3)
	if total != false {
		t.Errorf("IsReachedDestination from board3x3 was incorrect, got: %#v, want: %#v.", total, false)
	}

	total = board.IsReachedDestination([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}, 3)
	if total != true {
		t.Errorf("IsReachedDestination from board3x3 was incorrect, got: %#v, want: %#v.", total, true)
	}

	total = board2.IsReachedDestination([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 0}}, 4)
	if total != true {
		t.Errorf("IsReachedDestination from board4x4 was incorrect, got: %#v, want: %#v.", total, true)
	}

	total = board2.IsReachedDestination([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}}, 4)
	if total != false {
		t.Errorf("IsReachedDestination from board4x4 was incorrect, got: %#v, want: %#v.", total, false)
	}
}

func TestSwapTiles(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	total := board.SwapTiles([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}}, 2, 0, 2, 1)
	if total != [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}} {
		t.Errorf("SwapTiles from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}})
	}

	total = board.SwapTiles([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}}, 0, 1, 0, 2)
	if total != [3][3]int{{1, 3, 2}, {4, 5, 6}, {7, 0, 8}} {
		t.Errorf("SwapTiles from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 3, 2}, {4, 5, 6}, {7, 0, 8}})
	}

	total2 := board2.SwapTiles([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}}, 3, 0, 3, 1)
	if total2 != [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}} {
		t.Errorf("SwapTiles from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}})
	}

	total2 = board2.SwapTiles([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 0}, {12, 13, 14, 15}}, 2, 3, 2, 2)
	if total2 != [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 0, 11}, {12, 13, 14, 15}} {
		t.Errorf("SwapTiles from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 0, 11}, {12, 13, 14, 15}})
	}

}

func TestGetMove(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}})
	total := board.GetMove(8, 3)
	if total != 0 {
		t.Errorf("GetMove from board3x3 was incorrect, got: %d, want: %d.", total, 0)
	}

	board.New([3][3]int{{1, 2, 3}, {4, 5, 0}, {7, 8, 6}})
	total = board.GetMove(6, 3)
	if total != 2 {
		t.Errorf("GetMove from board3x3 was incorrect, got: %d, want: %d.", total, 2)
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}})
	total2 := board2.GetMove(15, 4)
	if total2 != 0 {
		t.Errorf("GetMove from board4x4 was incorrect, got: %d, want: %d.", total2, 0)
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 0}, {13, 14, 15, 12}})
	total2 = board2.GetMove(12, 4)
	if total2 != 2 {
		t.Errorf("GetMove from board4x4 was incorrect, got: %d, want: %d.", total2, 2)
	}
}

func TestRuturnMove(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	total := board.ReturnMove([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}}, 0, 3)
	if total != [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}} {
		t.Errorf("ReturnMove from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}})
	}

	board.New([3][3]int{{1, 2, 3}, {0, 5, 6}, {4, 7, 8}})
	total = board.ReturnMove([3][3]int{{1, 2, 3}, {0, 5, 6}, {4, 7, 8}}, 2, 3)
	if total != [3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}} {
		t.Errorf("ReturnMove from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	}

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	total = board.ReturnMove([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}}, 3, 3)
	if total != [3][3]int{{1, 2, 3}, {0, 5, 6}, {4, 7, 8}} {
		t.Errorf("ReturnMove from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 2, 3}, {0, 5, 6}, {4, 7, 8}})
	}

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}})
	total = board.ReturnMove([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}}, 1, 3)
	if total != [3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}} {
		t.Errorf("ReturnMove from board3x3 was incorrect, got: %d, want: %d.", total, [3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}})
	total2 := board2.ReturnMove([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}}, 0, 4)
	if total2 != [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}} {
		t.Errorf("ReturnMove from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}})
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {0, 6, 7, 8}, {5, 9, 10, 11}, {12, 13, 14, 15}})
	total2 = board2.ReturnMove([4][4]int{{1, 2, 3, 4}, {0, 6, 7, 8}, {5, 9, 10, 11}, {12, 13, 14, 15}}, 2, 4)
	if total2 != [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {0, 9, 10, 11}, {12, 13, 14, 15}} {
		t.Errorf("ReturnMove from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {0, 9, 10, 11}, {12, 13, 14, 15}})
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {0, 9, 10, 11}, {12, 13, 14, 15}})
	total2 = board2.ReturnMove([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {0, 9, 10, 11}, {12, 13, 14, 15}}, 3, 4)
	if total2 != [4][4]int{{1, 2, 3, 4}, {0, 6, 7, 8}, {5, 9, 10, 11}, {12, 13, 14, 15}} {
		t.Errorf("ReturnMove from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {0, 6, 7, 8}, {5, 9, 10, 11}, {12, 13, 14, 15}})
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}})
	total2 = board2.ReturnMove([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 0, 14, 15}}, 1, 4)
	if total2 != [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}} {
		t.Errorf("ReturnMove from board4x4 was incorrect, got: %d, want: %d.", total2, [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}})
	}
}

func TestGetAllMoves(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	var states = []slidingBlocks3x3.Direction{3, 0}
	total := board.GetAllMoves(3)
	if !reflect.DeepEqual(total, states) {
		t.Errorf("GetAllMoves from board3x3 was incorrect, got: %d, want: %d.", total, states)
	}

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}})
	var states2 = []slidingBlocks3x3.Direction{3, 1, 0}
	total2 := board.GetAllMoves(3)
	if !reflect.DeepEqual(total2, states2) {
		t.Errorf("GetAllMoves from board3x3 was incorrect, got: %d, want: %d.", total2, states2)
	}

	board.New([3][3]int{{1, 2, 3}, {4, 5, 0}, {7, 8, 6}})
	var states3 = []slidingBlocks3x3.Direction{3, 1, 2}
	total3 := board.GetAllMoves(3)
	if !reflect.DeepEqual(total3, states3) {
		t.Errorf("GetAllMoves from board3x3 was incorrect, got: %d, want: %d.", total3, states3)
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}})
	var states4 = []slidingBlocks4x4.Direction{3, 1, 0}
	total4 := board2.GetAllMoves(4)
	if !reflect.DeepEqual(total4, states4) {
		t.Errorf("GetAllMoves from board4x4 was incorrect, got: %d, want: %d.", total4, states4)
	}

	board2.New([4][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}})
	var states5 = []slidingBlocks4x4.Direction{0, 2}
	total5 := board2.GetAllMoves(4)
	if !reflect.DeepEqual(total5, states5) {
		t.Errorf("GetAllMoves from board4x4 was incorrect, got: %d, want: %d.", total5, states5)
	}

	board2.New([4][4]int{{8, 0, 1, 2}, {3, 4, 5, 6}, {7, 9, 10, 11}, {12, 13, 14, 15}})
	var states6 = []slidingBlocks4x4.Direction{1, 0, 2}
	total6 := board2.GetAllMoves(4)
	if !reflect.DeepEqual(total6, states6) {
		t.Errorf("GetAllMoves from board4x4 was incorrect, got: %d, want: %d.", total6, states6)
	}

	board2.New([4][4]int{{6, 7, 10, 15}, {5, 4, 8, 3}, {0, 12, 2, 1}, {14, 9, 13, 11}})
	var states7 = []slidingBlocks4x4.Direction{3, 0, 2}
	total7 := board2.GetAllMoves(4)
	if !reflect.DeepEqual(total7, states7) {
		t.Errorf("GetAllMoves from board4x4 was incorrect, got: %d, want: %d.", total7, states7)
	}

}
func TestStartPosition(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{5, 7, 2}, {1, 3, 6}, {8, 4, 0}})
	total, total2 := board.FindStartPosition(3)
	if total != 2 {
		t.Errorf("First position from board3x3 was incorrect, got: %v, want: %v.", total, 2)
	}
	if total2 != 2 {
		t.Errorf("Second postion from board3x3 was incorrect, got: %v, want: %v.", total, 2)
	}

	board.New([3][3]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 5}})
	total, total2 = board.FindStartPosition(3)
	if total != 1 {
		t.Errorf("First position from board3x3 was incorrect, got: %v, want: %v.", total, 1)
	}
	if total2 != 1 {
		t.Errorf("Second postion from board3x3 was incorrect, got: %v, want: %v.", total, 1)
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 0}})
	total3, total4 := board2.FindStartPosition(4)
	if total3 != 3 {
		t.Errorf("First position from board4x4 was incorrect, got: %v, want: %v.", total3, 3)
	}
	if total4 != 3 {
		t.Errorf("Second postion from board4x4 was incorrect, got: %v, want: %v.", total4, 3)
	}

	board2.New([4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {0, 13, 14, 15}})
	total3, total4 = board2.FindStartPosition(4)
	if total3 != 3 {
		t.Errorf("First position from board4x4 was incorrect, got: %v, want: %v.", total3, 3)
	}
	if total4 != 0 {
		t.Errorf("Second postion from board4x4 was incorrect, got: %v, want: %v.", total4, 0)
	}

}

func TestVisitedMoves(t *testing.T) {
	var board slidingBlocks3x3.SlidingBlocksBoard
	var board2 slidingBlocks4x4.SlidingBlocksBoard

	board.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {0, 7, 8}})
	var states = make(map[slidingBlocks3x3.Direction][3][3]int)
	states[3] = [3][3]int{{1, 2, 3}, {0, 5, 6}, {4, 7, 8}}
	states[0] = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 8}}

	total := board.VisitedMoves(3)
	if !reflect.DeepEqual(total, states) {
		t.Errorf("VisitedMoves from board3x3 was incorrect, got: %d, want: %d.", total, states)
	}

	board.New([3][3]int{{5, 6, 7}, {8, 0, 1}, {3, 2, 4}})
	var states2 = make(map[slidingBlocks3x3.Direction][3][3]int)
	states2[3] = [3][3]int{{5, 0, 7}, {8, 6, 1}, {3, 2, 4}}
	states2[1] = [3][3]int{{5, 6, 7}, {0, 8, 1}, {3, 2, 4}}
	states2[0] = [3][3]int{{5, 6, 7}, {8, 1, 0}, {3, 2, 4}}
	states2[2] = [3][3]int{{5, 6, 7}, {8, 2, 1}, {3, 0, 4}}

	total2 := board.VisitedMoves(3)
	if !reflect.DeepEqual(total, states) {
		t.Errorf("VisitedMoves from board3x3 was incorrect, got: %d, want: %d.", total2, states2)
	}

	board2.New([4][4]int{{6, 7, 10, 15}, {5, 4, 8, 3}, {0, 12, 2, 1}, {14, 9, 13, 11}})
	var states3 = make(map[slidingBlocks4x4.Direction][4][4]int)
	states3[3] = [4][4]int{{6, 7, 10, 15}, {0, 4, 8, 3}, {5, 12, 2, 1}, {14, 9, 13, 11}}
	states3[0] = [4][4]int{{6, 7, 10, 15}, {5, 4, 8, 3}, {12, 0, 2, 1}, {14, 9, 13, 11}}
	states3[2] = [4][4]int{{6, 7, 10, 15}, {5, 4, 8, 3}, {14, 12, 2, 1}, {0, 9, 13, 11}}

	total3 := board2.VisitedMoves(4)
	if !reflect.DeepEqual(total3, states3) {
		t.Errorf("VisitedMoves from board4x4 was incorrect, got: %d, want: %d.", total3, states3)
	}

	board2.New([4][4]int{{10, 13, 7, 0}, {14, 2, 4, 15}, {5, 6, 3, 9}, {11, 12, 8, 1}})
	var states4 = make(map[slidingBlocks4x4.Direction][4][4]int)
	states4[1] = [4][4]int{{10, 13, 0, 7}, {14, 2, 4, 15}, {5, 6, 3, 9}, {11, 12, 8, 1}}
	states4[2] = [4][4]int{{10, 13, 7, 15}, {14, 2, 4, 0}, {5, 6, 3, 9}, {11, 12, 8, 1}}

	total4 := board2.VisitedMoves(4)
	if !reflect.DeepEqual(total4, states4) {
		t.Errorf("VisitedMoves from board4x4 was incorrect, got: %d, want: %d.", total4, states4)
	}
}
