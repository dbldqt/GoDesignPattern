package prototype

type ICloneable interface {
	Clone() ICloneable
}

type Ceil struct {
	Name string
}

func (ceil *Ceil) Clone() ICloneable{
	ceil1 := *ceil
	return &ceil1
}

