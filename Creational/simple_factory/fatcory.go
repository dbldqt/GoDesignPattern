package simple_factory
//基础接口，其本身及其方法均为导出的方法，工厂所产生的所有子类对象都实现该接口
type Base interface {
	GetClassName() string
}

//该子类无需导出，所以小写字母开头
type classA struct {

}

func (a *classA) GetClassName() string{
	return "ClassA"
}

//该子类无需导出，所以小写字母开头
type classB struct {

}

func (b *classB) GetClassName() string{
	return "ClassB"
}


//工厂方法根据输入条件返回子类对象
//如果增加新的子类，则要修改该方法，不符合开闭原则
func NewClass(name string) Base{
	switch name {
		case "A":
			return &classA{}
		case "B":
			return &classB{}
	}
	return nil
}