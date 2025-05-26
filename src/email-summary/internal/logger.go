package internal

import (
	"log"
	"os"
)

const logFileName = "email_summary.log"

func Logger(err error) {

	if err == nil {
		return
	}

	logFile, errlog := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if errlog != nil {
		panic(errlog)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(err)

}
