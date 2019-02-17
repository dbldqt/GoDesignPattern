package simple_factory

//定义宝马车
type BmwCar struct {

}

func (bmw BmwCar) GetCarBrand() string{
	return "Bmw"
}