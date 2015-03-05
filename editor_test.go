package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type FakeWindow struct {
	width  int
	height int
	chars  [][]byte
	cursor Cursor
}

func NewFakeWindow(width, height int) *FakeWindow {
	chars := make([][]byte, height)
	for i := range chars {
		chars[i] = make([]byte, width)
	}
	return &FakeWindow{
		chars:  chars,
		width:  width,
		height: height,
	}
}

func (w *FakeWindow) CursorYX() (int, int) {
	return w.cursor.y, w.cursor.x
}

func (w *FakeWindow) Move(y, x int) {
	w.cursor.y = y
	w.cursor.x = x
}

func (w *FakeWindow) Println(args ...interface{}) {
	var words []string
	for _, arg := range args {
		str := fmt.Sprintf("%s", arg)
		words = append(words, str)
	}
	str := strings.Join(words, " ")
	for i := range str {
		w.chars[w.cursor.y][w.cursor.x] = str[i]
		w.cursor.x++
	}
	w.cursor.y++
	w.cursor.x = 0
}

func (w *FakeWindow) HasStringAt(y, x int, s string) bool {
	return bytes.Equal(w.chars[y][x:x+len(s)], []byte(s))
}

func TestFakePrintln(t *testing.T) {
	w := NewFakeWindow(10, 10)
	w.Println("Hi")
	if !w.HasStringAt(0, 0, "Hi") {
		t.Errorf("Failed to print 'Hi'")
	}
	if w.cursor.y != 1 {
		t.Errorf("Cursor didn't move down")
	}
	w.Println("There")
	if !w.HasStringAt(1, 0, "There") {
		t.Errorf("Failed to print 'There'")
	}
}

func newFile(lines ...string) *File {
	return &File{
		lines: lines,
	}
}

func TestEditorDown(t *testing.T) {
	w := NewFakeWindow(10, 10)
	e := NewEditor(w, newFile("this", "is", "a", "test"))
	e.MoveY(1)
	if e.cursor.y != 1 {
		t.Errorf("Failed to move cursor down")
	}
	e.MoveY(100)
	if e.cursor.y != 3 {
		t.Errorf("Downwards movement should be limited to the file bounds.")
	}
}
