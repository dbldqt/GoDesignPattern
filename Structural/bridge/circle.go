package bridge

import "fmt"

//定义画圆的接口
type IDraw interface {
	DrawCircle()
}

//定义红圆
type RedCircle struct {

}

func (rc *RedCircle) DrawCircle() {
	fmt.Println("draw red circle")
}


//定义绿圆
type GreenCircle struct {

}

func (gc GreenCircle) DrawCircle() {
	fmt.Println("draw green circle")
}

