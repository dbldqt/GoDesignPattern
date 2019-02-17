package simple_factory

//基础接口，其本身及其方法均为导出的方法，工厂所产生的所有子类对象都实现该接口
type ICar interface {
	GetCarBrand() string
}
