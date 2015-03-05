package main

type Window interface {
	CursorYX() (int, int)
	Move(y, x int)
	Println(args ...interface{})
}

type Cursor struct {
	x, y int
	// True if we should enter new lines on the right.
	end bool
	// X coordinate to enter a new line on.
	startX int
}

type Editor struct {
	window Window
	file   *File
	cursor Cursor
}

func NewEditor(w Window, f *File) *Editor {
	return &Editor{w, f, Cursor{}}
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

func (e *Editor) MoveCursorToLineEnd() {
	y := e.cursor.y
	e.MoveCursorTo(y, len(e.file.lines[y])-1)
	e.cursor.end = true
}

func (e *Editor) MoveCursorToLineStart() {
	y, _ := e.window.CursorYX()
	e.MoveCursorTo(y, 0)
}

func (e *Editor) MoveCursorTo(y, x int) {
	e.cursor.y, e.cursor.x = e.file.Constrain(y, x)
	e.refreshCursor()
}
