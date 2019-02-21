package bridge

type Circle struct {
	drawer IDraw
}

func (cir *Circle) Draw(){
	cir.drawer.DrawCircle()
}

func NewCircle(drawer IDraw) *Circle{
	return &Circle{drawer}
}
