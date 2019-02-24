package command

func ExampleCommand() {
	//创建命令主体
	mb := &MotherBoard{}
	//构建具体的命令对象，命令主体作为参数
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)
	//命令的入口，触发命令的地方
	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButtion1()
	box1.PressButtion2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButtion1()
	box2.PressButtion2()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}
