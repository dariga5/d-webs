package logger

var (
	logger *loggers
)

func InitLogger(path string) {
	logger = config(path)
}
