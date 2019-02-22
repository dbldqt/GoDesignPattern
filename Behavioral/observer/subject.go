package observer

type Subject struct {
	observerList []IObserver
	state int
}

func (sub *Subject) GetState() int{
	return sub.state
}
//每次设置state都通知观察者
func (sub *Subject) SetState(state int){
	sub.state = state
	sub.NotifyAllObservers()
}
//添加观察者
func (sub *Subject) Attach(obesrver IObserver){
	sub.observerList = append(sub.observerList,obesrver)
}
//通知观察者方法
func (sub *Subject) NotifyAllObservers(){
	for _,observer := range sub.observerList{
		observer.Update()
	}
}

func NewSubject() *Subject{
	return &Subject{
		observerList:make([]IObserver,0),
	}
}
