package facade

import "testing"

func Test(t *testing.T){
	shapeMaker := NewShapeMaker()
	//使用外观类提供的简单方法进行操作
	shapeMaker.DrawCircle()
	shapeMaker.DrawRectangle()
	shapeMaker.DrawSquare()
}
