package simpleFatory
//基础接口，其本身及其方法均为导出的方法，工厂所产生的所有子类对象都实现该接口
type Base interface {
	GetClassName() string
}

//该子类无需导出，所以小写字母开头
type classA struct {

}

func (a *classA) GetClassName() string{
	return "classA"
}

//该子类无需导出，所以小写字母开头
type classB struct {

}

func (b *classB) GetClassName() string{
	return "classB"
}


//定义一个工厂类的基础接口
type Factory interface {
	GetInstance() Base
}


//定义一个用于创建classA对象的工厂，实现产生Base对象的接口
type FactoryA struct {

}

func (fa *FactoryA) GetInstance() Base{
	return &classA{}
}


//定义一个用于创建classB对象的工厂，实现产生Base对象的接口
type FactoryB struct {

}

func (fb *FactoryB) GetInstance() Base{
	return &classB{}
}


//此时如果再增加一个base的子类，只需要添加一个子类以及一个产生该子类对象的方法即可，符合开闭原则，
//但是要在使用处指定使用的工厂类
type classC struct {

}

func (c *classC) GetClassName() string{
	return "classC"
}

type FactoryC struct {

}

func (fc *FactoryC) GetInstance() Base{
	return &classC{}
}

