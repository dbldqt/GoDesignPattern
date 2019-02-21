package flyweight

import (
	"fmt"
)

type IShape interface {
	Draw()
}

type Circle struct {
	id int
	color string
	x int
	y int
	radius int
}

func (cir *Circle) Draw() {
	fmt.Printf("Circle::Draw() id:%d color:%s x:%d y:%d radius:%d",cir.id,cir.color,cir.x,cir.y,cir.radius)
	fmt.Println()
}

func (cir *Circle) SetX(x int){
	cir.x = x
}

func (cir *Circle) SetY(y int){
	cir.y = y
}

func (cir *Circle) SetRadius(radius int){
	cir.radius = radius
}

func NewCircle(color string) *Circle{
	return &Circle{color:color}
}

