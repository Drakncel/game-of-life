package main

import "./utils/rendering"
import tm "github.com/buger/goterm"
import "time"

func render (board []rendering.Coord) {
    // TOTRY: Maybe use a box
    for _, coord := range board {
        tm.MoveCursor(coord.X, coord.Y)
        tm.Println(tm.Color(tm.Bold(rendering.BLOCK), coord.Color))
        tm.Flush()
    }
}

func main() {
    //https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
    board := rendering.CreateEmptyBoard(20, 10)
    board = rendering.MutateRandomBoard(board)

    // GET UPDATE FROM GAME LOOP
    // RENDER EVERY SECOND ON TERM
    tm.Clear()
    for {
        board = rendering.GetNextStep(board)
        render(board)
        time.Sleep(time.Second / 10);
    }
}
