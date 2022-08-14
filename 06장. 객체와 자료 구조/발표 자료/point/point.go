package point

type Point interface {
	X() int
	Y() int
	SetCartesian(x, y int)
}

// go 에서는 이름의 앞 글자가 대문자면 public 입니다.
// 반대로 소문자라면 private 입니다.
type point struct {
	x int
	y int
}

func New() Point {
	return &point{0, 0}
}

func (p *point) X() int {
	return p.x
}

func (p *point) Y() int {
	return p.y
}

func (p *point) SetCartesian(x, y int) {
	p.x = x
	p.y = y
}
