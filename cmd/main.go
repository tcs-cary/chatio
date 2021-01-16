package main

import (
	"fmt"
	"os"
	// "time"

	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

func main() {
	fmt.Println("Welcome to chat.io!")
	var user string
	fmt.Print("What's your name? ")
	_, err := fmt.Scanln(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" ")
	userMF := "[%s]: %s" // userMessageFormat
	fmt.Printf(userMF, user, "Hello, World")

	err = termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		draw()
	}

}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func draw() { // box at bottom of terminal to type into
	const coldef = termbox.ColorDefault
	const edit_box_width = 30

	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - edit_box_width) / 2

	// unicode box drawing chars around the edit box
	if runewidth.EastAsianWidth {
		termbox.SetCell(midx-1, midy, '|', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy, '|', coldef, coldef)
		termbox.SetCell(midx-1, midy-1, '+', coldef, coldef)
		termbox.SetCell(midx-1, midy+1, '+', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy-1, '+', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy+1, '+', coldef, coldef)
		fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '-'})
		fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '-'})
	} else {
		termbox.SetCell(midx-1, midy, '│', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy, '│', coldef, coldef)
		termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
		termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy-1, '┐', coldef, coldef)
		termbox.SetCell(midx+edit_box_width, midy+1, '┘', coldef, coldef)
		fill(midx, midy-1, edit_box_width, 1, termbox.Cell{Ch: '─'})
		fill(midx, midy+1, edit_box_width, 1, termbox.Cell{Ch: '─'})
	}

	termbox.SetCursor(midx, midy)

	tbprint(midx+6, midy+3, coldef, coldef, "Press ESC to quit")
	termbox.Flush()
}
