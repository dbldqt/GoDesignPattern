package memento

//创建备忘录类
type Memento struct {
	state string
}

func (self *Memento) GetState() string{
	return self.state
}