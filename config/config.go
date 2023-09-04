package config

var logger *Logger

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
