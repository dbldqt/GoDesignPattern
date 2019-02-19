package prototype

type PrototypeManager struct {
	prototypes map[string]ICloneable
}

func (pm *PrototypeManager) set(key string,pb ICloneable){
	pm.prototypes[key] = pb
}

func (pm *PrototypeManager) get(key string) ICloneable  {
	pb,ok := pm.prototypes[key]
	if ok == true{
		return pb.Clone()   //如果找到了该类型的变量，则返回一个clone对象
	}else{
		return nil
	}
}

func NewPrototypeManager() *PrototypeManager{
	return &PrototypeManager{
		make(map[string]ICloneable),
	}
}