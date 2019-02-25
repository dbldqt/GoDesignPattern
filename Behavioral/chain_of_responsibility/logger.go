package chain_of_responsibility

import "fmt"

//定义记录器接口
type ILogger interface {
	SetNexLogger(log ILogger)
	GetNexLogger() ILogger
	LogMessage(level int,msg string)
	write(msg string)
}
const (
	INFO = 1
	DEBUG = 2
	ERROR = 3
)
//定义基础记录器
type BaseLogger struct {
	Level int
	nextLogger ILogger
}

func (self *BaseLogger) SetNexLogger(logger ILogger) {
	self.nextLogger = logger
}

func (self *BaseLogger) GetNexLogger() ILogger{
	return self.nextLogger
}

func (self *BaseLogger) LogMessage(level int,msg string) {
	panic("implement me")
}

func (self *BaseLogger) write(msg string) {
	panic("implement me")
}
//定义控制台记录器
type ConsoleLogger struct {
	BaseLogger
}
//此处重写write
func (self *ConsoleLogger) write(msg string){
	fmt.Println("consoleLogger log msg "+msg)
}
//此处要重写LogMessage,尽管方法本身没有任何改变，但是该函数调用了self.write,
//不覆写这个方法，调不到覆写的write方法
func (self *ConsoleLogger) LogMessage(level int,msg string) {
	if self.Level <= level{
		self.write(msg)
	}

	if self.nextLogger != nil{
		self.nextLogger.LogMessage(level,msg)
	}
}
//定义错误记录器
type ErrorLogger struct {
	BaseLogger
}

func (self *ErrorLogger) write(msg string){
	fmt.Println("ErrorLogger log msg "+msg)
}

func (self *ErrorLogger) LogMessage(level int,msg string) {
	if self.Level <= level{
		self.write(msg)
	}

	if self.nextLogger != nil{
		self.nextLogger.LogMessage(level,msg)
	}
}
//定义文件记录器
type FileLogger struct {
	BaseLogger
}

func (self *FileLogger) write(msg string){
	fmt.Println("FileLogger log msg "+msg)
}
func (self *FileLogger) LogMessage(level int,msg string) {
	if self.Level <= level{
		self.write(msg)
	}

	if self.nextLogger != nil{
		self.nextLogger.LogMessage(level,msg)
	}
}




