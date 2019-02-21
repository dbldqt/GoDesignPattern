package flyweight

import (
	"testing"
	"math/rand"
	"time"
)
var colors = []string{"Red","Green","Blue","White","Black"}

func TestCircle_Draw(t *testing.T) {
	for i:=0;i<20;i++{
		circle := GetShapeFactory().GetCircle(getRandomColor())
		circle.SetX(getRandomNum())
		circle.SetY(getRandomNum())
		circle.SetRadius(getRandomNum())
		circle.Draw()
	}
}

func getRandomNum() int{
	 rand.Seed(time.Now().UnixNano())
	 return rand.Int()
}

func getRandomColor() string{
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(4)
	return colors[rand]
}

