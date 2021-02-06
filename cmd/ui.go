package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type ui struct {
	username   string
	isLoggedIn bool
	editBox    EditBox
}

func (ui ui) draw() {
	if ui.isLoggedIn {
		ui.drawChatroom()
		return
	}
	ui.drawLogin(30)
}

func (ui ui) drawChatroom() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - 30) / 2
	tbprint(midx+6, midy+3, coldef, coldef, ui.username)
	termbox.Flush()
}

func (ui ui) drawLogin(edit_box_width int) {
	const coldef = termbox.ColorDefault
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

	ui.editBox.Draw(midx, midy, edit_box_width, 1)
	termbox.SetCursor(midx+ui.editBox.CursorX(), midy)

	tbprint(midx+6, midy+3, coldef, coldef, "Press ESC to quit")
	termbox.Flush()
}
