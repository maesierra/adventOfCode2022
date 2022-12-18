package common

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) ManhattanDistance(other Point) int {
	return IntAbs(p.X - other.X) + IntAbs(p.Y - other.Y)
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

type Line struct {
	X1 int
	Y1 int 
	X2 int 
	Y2 int
}

func (l Line) IsHorizontal() bool{
	return l.Y1 == l.Y2
}

func (l Line) IsVertical() bool{
	return l.X1 == l.X2
}

func (l Line) Draw() []Point {
	points := []Point{}
	direction := 1
	if l.IsHorizontal() {		
		if l.X2 < l.X1 {
			direction = -1
		}
		for x := l.X1; x != l.X2; x+= direction {
			points = append(points, Point{x, l.Y1})	
		}	
		points = append(points, Point{l.X2, l.Y2})	
		return points
	} else if l.IsVertical() {
		if l.Y2 < l.Y1 {
			direction = -1
		}
		for y := l.Y1; y != l.Y2; y+= direction {
			points = append(points, Point{l.X1, y})	
		}	
		points = append(points, Point{l.X2, l.Y2})	
		return points
	}
	panic("line not straight")
}

type BoundingBox struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (b BoundingBox) Width() int {
	return IntAbs(b.X2 - b.X1) + 1
}

func (b BoundingBox) Height() int {
	return IntAbs(b.Y2 - b.Y1) + 1
}

func (b BoundingBox) Contains(x, y int) bool {
	return b.X1 <= x && b.X2 >= x && b.Y1 <= y && b.Y2 >= y
}

func (b BoundingBox) Move(x, y int) BoundingBox {
	return BoundingBox{
		X1: b.X1 + x, 
		Y1: b.Y1 + y, 
		X2: b.X2 + x, 
		Y2: b.Y2 + y,
	}
}

func (r1 BoundingBox) Intersects(r2 BoundingBox) bool {
	return !(r2.X1 > r1.X2 || r2.X2 < r1.X1 || r2.Y2 < r1.Y1 || r2.Y1 > r1.Y2);
	
}

func NewBoundingBox(x, y, height, width int) BoundingBox {
	return BoundingBox{X1: x, Y1: y, X2: x + width, Y2: y + height}
}