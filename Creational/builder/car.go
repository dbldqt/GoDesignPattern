package builder
//定义汽车接口
type ICar interface {
	GetWheel() string
	GetEngine() string
}
//定义汽车实体
/**
	简单定义产品的属性，明确对产品的描述。我们继续来思考，因为我们的产品是比较抽象的，它没有指定引擎的型号，也没有指定车轮的牌子，
	那么这样的组合方式有很多，完全要靠建造者来建造，建造者说要生产一辆奔驰SUV那就得用奔驰的引擎和奔驰的车轮，该建造者对于一个具体的产品来说是绝对的权威
 */
type Car struct {
	wheel string
	engine string
}

func (car Car) GetWheel() string {
	return car.wheel
}

func (car Car) GetEngine() string {
	return car.engine
}

func (car Car) GetInfo() string{
	return "车轮是"+car.wheel+",引擎是"+car.engine
}

func NewCar(wheel string,engine string) ICar{
	car := Car{
		wheel,
		engine,
	}
	return car
}



