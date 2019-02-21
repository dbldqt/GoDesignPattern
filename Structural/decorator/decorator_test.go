package decorator

import "testing"

func Test(t *testing.T){
	circle := Circle{}
	redCircle := RedShapeDecorator{ShapeDecorator{&Circle{}}}
	redRectangle := RedShapeDecorator{ShapeDecorator{&Rectangle{}}}

	circle.Draw()
	redCircle.Draw()
	redRectangle.Draw()
}







