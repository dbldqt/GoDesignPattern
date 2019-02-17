package simple_factory

//工厂方法根据输入条件返回子类对象
//如果增加新的子类，则要修改该方法，不符合开闭原则
func NewClass(name string) ICar{
	switch name {
		case "Bmw":
			return BmwCar{}
		case "Benz":
			return BenzCar{}
	}
	return nil
}