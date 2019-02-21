package decorator

type ShapeDecorator struct {
	shape IShape
}

func (sd *ShapeDecorator) Draw(){
	sd.shape.Draw()
}

type RedShapeDecorator struct {
	ShapeDecorator
}

func (rsd *RedShapeDecorator) Draw(){
	rsd.shape.Draw()
	println("border color: Red")
}

