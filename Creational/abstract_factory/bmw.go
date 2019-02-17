package abstract_factory

//定义宝马车基础类，实现ICar接口
type BmwCar struct {

}

func (bmw BmwCar) GetBrand() string {
	return "BMW"
}

func (bmw BmwCar) GetModel() string {
	return "not define"
}
//定义宝马suv，继承宝马基础类
type BmwSuv struct {
	BmwCar
}

func (bwmsuv BmwSuv) GetModel() string {
	return "BmwSuv"
}

//定义宝马MPV，继承BmwCar基类
type BmwMpv struct {
	BmwCar
}

func (bmwmpv BmwMpv) GetModel() string {
	return "BmwMpv"
}

//定义宝马汽车工厂
type BwmFatcory struct {

}

func (bmwf BwmFatcory) CreateSuv() ICar {
	return BmwSuv{}
}

func (bmwf BwmFatcory) CreateMpv() ICar {
	return BmwMpv{}
}
