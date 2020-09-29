package rendering

import "math/rand"
import tm "github.com/buger/goterm"
import "time"

const BALL = "O"
const BLOCK = "█"

type Coord struct {
    X int
    Y int
    Color int
    State int
}

type Board []Coord

func CreateEmptyBoard(width int, height int) Board {
    var board Board
    x := 0
    y := 0

    for x < width {
        for y < height {
            board = append(board, Coord{X: x, Y: y, Color: tm.WHITE, State: 0})
            y++
        }
        x++
        y = 0
    }
    return board
}

func MutateRandomBoard(board Board) Board {
    var newBoard Board
    rand.Seed(time.Now().UnixNano())
    for _, coord := range board {
        nb := rand.Intn(10 - 1) + 1
        if (nb > 5) {
            newBoard = append(newBoard, Coord{ X: coord.X, Y: coord.Y, Color: tm.BLUE, State: 1 })
        } else {
            newBoard = append(newBoard, Coord{ X: coord.X, Y: coord.Y, Color: tm.WHITE, State: 0 })
        }
    }

    return newBoard
}

func getState(board Board, x int, y int) int {
    for _, coord := range board {
        if coord.X == x && coord.Y == y {
            return coord.State
        }
    }
    return 0
}

func GetNextStep(board Board) Board {
    var newBoard Board
    for _, coord := range board {
        aliveN := 0
        n1 := getState(board, coord.X - 1, coord.Y)
        n2 := getState(board, coord.X - 1, coord.Y - 1)
        n3 := getState(board, coord.X, coord.Y - 1)
        n4 := getState(board, coord.X + 1, coord.Y)
        n5 := getState(board, coord.X + 1, coord.Y - 1)
        n6 := getState(board, coord.X + 1, coord.Y + 1)
        n7 := getState(board, coord.X - 1, coord.Y + 1)
        n8 := getState(board, coord.X, coord.Y + 1)
        aliveN = n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8

        if (coord.State == 0 && aliveN == 3) || (coord.State == 1 && (aliveN == 2 || aliveN == 3)){
            newBoard = append(newBoard, Coord{ X: coord.X, Y: coord.Y, Color: tm.BLUE, State: 1 })
        } else {
            newBoard = append(newBoard, Coord{ X: coord.X, Y: coord.Y, Color: tm.WHITE, State: 0 })
        }
    }
    return newBoard
}
