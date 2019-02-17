package abstract_factory

//定义奔驰车基础类，实现ICar接口
type BenzCar struct {

}

func (benz BenzCar) GetBrand() string {
	return "Benz"
}

func (benz BenzCar) GetModel() string {
	return "not define"
}

//定义奔驰suv，继承奔驰基础类
type BenzSuv struct {
	BenzCar
}

func (benzsuv BenzSuv) GetModel() string {
	return "BenzSuv"
}

//定义奔驰MPV，继承BenzCar基类
type BenzMpv struct {
	BenzCar
}

func (benzmpv BenzMpv) GetModel() string {
	return "BenzMpv"
}

//定义奔驰汽车工厂
type BenzFactory struct {

}

func (benzf BenzFactory) CreateSuv() ICar {
	return BenzSuv{}
}

func (benzf BenzFactory) CreateMpv() ICar {
	return BenzMpv{}
}