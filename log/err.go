package log

func (logger *Logger) CheckErr(err error) {
	logger.CheckErrLog(err)
}

func (logger *Logger) CheckErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (logger *Logger) CheckErrFatal(err error) {
	if err != nil {
		logger.logger.Fatal(err)
	}
}

func (logger *Logger) CheckErrLog(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
