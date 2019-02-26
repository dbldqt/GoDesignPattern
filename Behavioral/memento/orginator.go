package memento

type Originator struct {
	state string
}

func (self *Originator) SetState(state string){
	self.state = state
}

func (self *Originator) Getstate() string{
	return self.state
}

func (self *Originator) SaveStateToMemento() Memento{
	return Memento{self.state}
}

func (self *Originator) GetStateFromMemento(memento Memento){
	self.state = memento.GetState()
}


