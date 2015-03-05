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
	return f.ConstrainY(y), f.ConstrainX(y, x)
}

func (f *File) ConstrainY(y int) int {
	if f.IsEmpty() {
		return 0
	}
	return max(0, min(len(f.lines)-1, y))
}

func (f *File) ConstrainX(y, x int) int {
	if f.IsEmpty() {
		return 0
	}
	return max(0, min(len(f.lines[y])-1, x))
}

func (f *File) IsEmpty() bool {
	return len(f.lines) == 0
}
