package memento

type CareTaker struct {
	list []Memento
}

func (self *CareTaker) Add(memento Memento){
	self.list = append(self.list,memento)
}

func (self *CareTaker) Get(index int) Memento{
	return self.list[index]
}

func NewCareTaker() *CareTaker{
	return &CareTaker{
		[]Memento{},
	}
}