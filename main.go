package main

import (
	"code.google.com/p/goncurses"
	"log"
)

func main() {
	window, err := goncurses.Init()
	if err != nil {
		log.Fatal("failed to create screen")
	}
	goncurses.Echo(false)
	defer goncurses.End()

	e := NewEditor(window)
	e.file.lines = []string{"this", "is", "a", "test"}
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
