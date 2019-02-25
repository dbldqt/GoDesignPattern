package iterator
//定义迭代器接口
type IIterator interface {
	HasNext() bool
	Next() interface{}
}

//定义实现迭代接口的类
type IntegerContainer struct {
	list []int
	curIndex int
}

func (self *IntegerContainer) HasNext() bool {
	return (len(self.list)-self.curIndex) > 0
}

func (self *IntegerContainer) Next() int {
	next := self.list[self.curIndex]
	self.curIndex += 1
	return next
}

func (self *IntegerContainer) Append(num int){
	self.list = append(self.list,num)
}

func NewContainer() *IntegerContainer{
	return &IntegerContainer{
		make([]int,0),
		0,
	}
}



