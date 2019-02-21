package flyweight

import "sync"

var shapeFactory *ShapeFactory
var once sync.Once
type ShapeFactory struct {
	shapeMap map[string]*Circle
}

func (sf *ShapeFactory) GetCircle(color string) *Circle {
	circle,ok := sf.shapeMap[color]
	if ok{
		return circle
	}else{
		circle = NewCircle(color)
		circle.id = len(sf.shapeMap)
		sf.shapeMap[color] = circle
		return circle
	}
}
//go语言中没有类静态变量，只能使用对象存储，而且要使用单例模式
func GetShapeFactory() *ShapeFactory{
	once.Do(func() {
		shapeFactory = &ShapeFactory{
			make(map[string]*Circle,1),
		}
	})
	return shapeFactory
}