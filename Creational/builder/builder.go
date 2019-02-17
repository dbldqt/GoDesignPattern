package builder

//定义汽车建造者基础类
type ICarBuilder interface {
	buildWheel() string
	buildEngine() string
	SetBlueprint(bp *CarBluePrint) bool
	BuildCar() ICar
}
//定义基础的carBuilder类
type CarBuilder struct {
	carBlueprint *CarBluePrint
}

func (cb CarBuilder) buildWheel() string {
	return cb.carBlueprint.GetWheel()
}

func (cb CarBuilder) buildEngine() string {
	return cb.carBlueprint.GetEngine()
}

func (cb CarBuilder) SetBlueprint(bp *CarBluePrint) bool{
	if bp != nil{
		cb.carBlueprint = bp
		return true
	}else{
		return false
	}
}

func (cb CarBuilder) BuildCar() ICar{
	return NewCar(cb.carBlueprint.GetWheel(),cb.carBlueprint.GetEngine())
}
