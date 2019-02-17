package builder

//定义食物条目接口
type IItem interface {
	Name() string
	Packing() Packing
	Price() float64
}

//定义默认汉堡类
type Burger struct {

}

func (burger Burger) Name() string {
	return "Burger"
}

func (burger Burger) Packing() Packing {
	return Wrapper{}
}

func (burger Burger) Price() float64 {
	return 1.0
}

//定义默认饮料类
type Drink struct {

}

func (cd Drink) Name() string {
	return "Drink"
}

func (cd Drink) Packing() Packing {
	return Bottle{}
}

func (cd Drink) Price() float64 {
	return 1.0
}

//定义蔬菜汉堡,继承默认汉堡类
type VegBurger struct {
	Burger
}
//重写父类的方法
func (vb VegBurger) Name() string {
	return "VegBurger"
}

func (vb VegBurger) Price() float64 {
	return 2.5
}
//定义鸡肉汉堡，继承自默认汉堡
type ChickenBurger struct {
	Burger
}
//重写父类的方法
func (cb ChickenBurger) Name() string{
	return "CheckenBurger"
}

func (cb ChickenBurger) Pirce() float64{
	return 5.0
}

//定义可口饮料，继承自默认饮料类
type Coke struct {
	Drink
}

func (coke Coke) Name() string {
	return "Coke"
}

func (coke Coke) Price() float64 {
	return 3.0
}

//定义百事饮料，继承自默认饮料
type Pepsi struct {
	Drink
}

func (pepsi Pepsi) Name() string {
	return "Pepsi"
}

func (pepsi Pepsi) Price() float64 {
	return 3.0
}







