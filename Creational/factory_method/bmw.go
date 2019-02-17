package factory_method

//定义宝马车
type BmwCar struct {

}

func (bmw BmwCar) GetCarBrand() string{
	return "Bmw"
}

//定义宝马工厂
type BmwFactory struct {

}

func (bmwf BmwFactory) CreateCar() ICar {
	return BmwCar{}
}
