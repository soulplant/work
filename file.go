package main

type File struct {
	lines []string
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (f *File) Constrain(y, x int) (int, int) {
	if f.IsEmpty() {
		return 0, 0
	}
	y = max(0, min(len(f.lines)-1, y))
	x = max(0, min(len(f.lines[y])-1, x))
	return y, x
}

func (f *File) ConstrainY(y int) int {
	if f.IsEmpty() {
		return 0
	}
	return max(0, min(len(f.lines)-1, y))
}

func (f *File) IsEmpty() bool {
	return len(f.lines) == 0
}
