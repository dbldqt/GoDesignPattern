package main

import (
	chian "GoDesignPattern/Behavioral/chain_of_responsibility"
)

func GetChainLogger() chian.ILogger{
	errorLogger := chian.ErrorLogger{chian.BaseLogger{Level:chian.ERROR}}
	fileLogger := chian.FileLogger{chian.BaseLogger{Level:chian.DEBUG}}
	consoleLogger := chian.ConsoleLogger{chian.BaseLogger{Level:chian.INFO}}

	errorLogger.SetNexLogger(&fileLogger)
	fileLogger.SetNexLogger(&consoleLogger)
	return &errorLogger
}

func main() {
	chainLogger := GetChainLogger()
	chainLogger.LogMessage(chian.INFO,"info")
	chainLogger.LogMessage(chian.ERROR,"error")
	chainLogger.LogMessage(chian.DEBUG,"debug")
}





