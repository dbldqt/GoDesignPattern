package facade
//定义图形画图接口
type IShape interface {
	Draw()
}
//定义长方形
type Rectangle struct {

}

func (rect *Rectangle) Draw(){
	println("Rectangle::Draw()")
}
//定义正方形
type Square struct {

}

func (squa *Square) Draw(){
	println("Square::Draw()")
}

//定义圆形
type Circle struct {

}

func (cir *Circle) Draw(){
	print("Circle::Draw()")
}


//定义外观类
type ShapeMaker struct {
	rectangle *Rectangle
	square *Square
	circle *Circle
}

func (sm *ShapeMaker) DrawRectangle(){
	sm.rectangle.Draw()
}

func (sm *ShapeMaker) DrawSquare(){
	sm.square.Draw()
}

func (sm *ShapeMaker) DrawCircle(){
	sm.circle.Draw()
}

func NewShapeMaker() *ShapeMaker{
	return &ShapeMaker{
		&Rectangle{},
		&Square{},
		&Circle{},
	}
}