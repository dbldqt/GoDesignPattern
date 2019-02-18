package builder
//定义导演类
/**
场景类只要找到导演类（也就是车间主任了）说给我制造一辆这样的宝马车，车间主任马上通晓你的意图，
设计了一个蓝图，然后命令建造车间拼命加班加点建造，最终返回给你一件最新出品的产品
 */
type Director struct {
	CarBuilder ICarBuilder
}

func mopCreateCar(carBuilder ICarBuilder,engine string,wheel string) ICar{
	blueprint := CarBluePrint{}
	blueprint.SetEngine(engine)
	blueprint.SetWheel(wheel)
	carBuilder.SetBlueprint(&blueprint)
	return carBuilder.BuildCar()
}

func (d Director) CreateBenzSuv() ICar {
	return mopCreateCar(d.CarBuilder,"benz的引擎","benz的轮胎")
}

func (d Director) CreateBWMVan() ICar{
	return mopCreateCar(d.CarBuilder,"BMW的引擎","BMW的轮胎")
}

func (d Director) CreateComplexCar() ICar{
	return mopCreateCar(d.CarBuilder,"BMW的引擎","benz的轮胎")
}






