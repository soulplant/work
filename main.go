package main

import (
	"code.google.com/p/goncurses"
	"log"
)

type Cursor struct {
	x, y int
	// True if we should enter new lines on the right.
	end bool
	// X coordinate to enter a new line on.
	startX int
}

type Editor struct {
	window *goncurses.Window
	file   *File
	cursor Cursor
}

func NewEditor() *Editor {
	window, err := goncurses.Init()
	goncurses.Echo(false)
	if err != nil {
		log.Fatal("failed to create screen")
	}
	return &Editor{window, &File{}, Cursor{}}
}

func (e *Editor) Close() {
	goncurses.End()
}

func (e *Editor) YX() (int, int) {
	return e.window.CursorYX()
}

func (e *Editor) Draw() {
	e.window.Move(0, 0)
	for _, line := range e.file.lines {
		e.window.Println(line)
	}
	e.refreshCursor()
}

func (e *Editor) MoveY(delta int) {
	e.cursor.y = e.file.ConstrainY(e.cursor.y + delta)
	if e.cursor.end {
		e.MoveCursorToLineEnd()
	} else {
		e.cursor.x = e.cursor.startX
	}
	e.refreshCursor()
}

func (e *Editor) MoveX(delta int) {
	e.cursor.x = e.file.ConstrainX(e.cursor.y, e.cursor.x+delta)
	e.cursor.startX = e.cursor.x
	e.cursor.end = false
	e.refreshCursor()
}

func (e *Editor) refreshCursor() {
	y, x := e.file.Constrain(e.cursor.y, e.cursor.x)
	e.window.Move(y, x)
}

func (e *Editor) GetChar() goncurses.Key {
	return e.window.GetChar()
}

func (e *Editor) MoveCursorToLineEnd() {
	y := e.cursor.y
	e.MoveCursorTo(y, len(e.file.lines[y])-1)
	e.cursor.end = true
}

func (e *Editor) MoveCursorToLineStart() {
	y, _ := e.YX()
	e.MoveCursorTo(y, 0)
}

func (e *Editor) MoveCursorTo(y, x int) {
	e.cursor.y, e.cursor.x = e.file.Constrain(y, x)
	e.refreshCursor()
}

func main() {
	e := NewEditor()
	e.file.lines = []string{"this", "is", "a", "test"}
	defer e.Close()
	done := false
	for !done {
		e.Draw()
		c := e.GetChar()
		switch c {
		case 'j':
			e.MoveY(1)
		case 'k':
			e.MoveY(-1)
		case 'h':
			e.MoveX(-1)
		case 'l':
			e.MoveX(1)
		case '$':
			e.MoveCursorToLineEnd()
		case '0':
			e.MoveCursorToLineStart()
		case 'q':
			done = true
		}
	}
}
