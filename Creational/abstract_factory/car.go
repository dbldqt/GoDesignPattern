package abstract_factory

//定义汽车接口
type ICar interface {
	GetBrand() string
	GetModel() string
}

//定义汽车工厂接口
type ICarFactory interface {
	CreateSuv() ICar
	CreateMpv() ICar
}

