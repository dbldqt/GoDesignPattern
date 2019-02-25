package chain_of_responsibility

import (
	"testing"
)

func GetChainLogger() ILogger{
	errorLogger := ErrorLogger{BaseLogger{Level:ERROR}}
	fileLogger := FileLogger{BaseLogger{Level:DEBUG}}
	consoleLogger := ConsoleLogger{BaseLogger{Level:INFO}}

	errorLogger.SetNexLogger(&fileLogger)
	fileLogger.SetNexLogger(&consoleLogger)
	return &errorLogger
}
func TestFileLogger_Write(t *testing.T) {
	chainLogger := GetChainLogger()
	chainLogger.LogMessage(INFO,"info")
	chainLogger.LogMessage(ERROR,"error")
	chainLogger.LogMessage(DEBUG,"debug")
}


