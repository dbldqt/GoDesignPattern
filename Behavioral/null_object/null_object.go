package null_object
//定义顾客接口
type ICustomer interface {
	IsNil() bool
	GetName() string
}
//定义真实的顾客
type RealCustomer struct {
	name string
}

func (self *RealCustomer) IsNil() bool {
	return false
}

func (self *RealCustomer) GetName() string {
	return self.name
}
//定义空顾客
type NilCustomer struct {

}

func (self *NilCustomer) IsNil() bool {
	return true
}

func (self *NilCustomer) GetName() string {
	return "no name data"
}


func NewCustomer(name string)ICustomer{
	names := []string{"Rob","Bob","Julie","Laura"}
	for _,curName := range names{
		if curName == name{
			return &RealCustomer{name}
		}
	}

	return &NilCustomer{}
}

