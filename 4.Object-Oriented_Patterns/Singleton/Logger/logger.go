package Logger

import (
	"log"
	"os"
	"sync"
)

type hLogger struct {
	*log.Logger
	filename string
}

var hlogger *hLogger
var once sync.Once

func GetInstance() *hLogger {
	//only 1 execution:
	once.Do(func() {
		hlogger = createLogger("hydralogger.log")
	})
	return hlogger
}

func createLogger(fname string) *hLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
