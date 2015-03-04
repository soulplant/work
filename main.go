package main

import (
	"code.google.com/p/goncurses"
	"fmt"
	"log"
)

type Cursor struct {
	x, y int
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
	e.MoveCursorTo(e.cursor.y, e.cursor.x)
}

func (e *Editor) MoveCursorBy(dy, dx int) {
	e.cursor.y += dy
	e.cursor.y = e.file.ConstrainY(e.cursor.y)
	e.cursor.x += dx
	e.refreshCursor()
}

func (e *Editor) refreshCursor() {
	e.window.Move(e.cursor.y, e.cursor.x)
}

func (e *Editor) GetChar() goncurses.Key {
	return e.window.GetChar()
}

func (e *Editor) MoveCursorToLineEnd() {
	y, _ := e.YX()
	e.MoveCursorTo(y, len(e.file.lines[y])-1)
}

func (e *Editor) MoveCursorToLineStart() {
	y, _ := e.YX()
	e.MoveCursorTo(y, 0)
}

func (e *Editor) MoveCursorTo(y, x int) {
	y, x = e.file.Constrain(y, x)
	e.window.Move(y, x)
}

func main() {
	fmt.Printf("here we go")
	e := NewEditor()
	e.file.lines = []string{"this", "is", "a", "test"}
	defer e.Close()
	done := false
	for !done {
		e.Draw()
		c := e.GetChar()
		switch c {
		case 'j':
			e.MoveCursorBy(1, 0)
		case 'k':
			e.MoveCursorBy(-1, 0)
		case 'h':
			e.MoveCursorBy(0, -1)
		case 'l':
			e.MoveCursorBy(0, 1)
		case '$':
			e.MoveCursorToLineEnd()
		case '0':
			e.MoveCursorToLineStart()
		case 'q':
			done = true
		}
	}
}
