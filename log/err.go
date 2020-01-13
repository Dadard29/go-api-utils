package log

func (logger *Logger) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}