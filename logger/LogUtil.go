package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	trace  *log.Logger
	info   *log.Logger
	error  *log.Logger
	waring *log.Logger
)

func init() {
	trace = log.New(ioutil.Discard,
		"TRACE:",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(os.Stdout,
		"INFO:",
		log.Ldate|log.Ltime|log.Lshortfile)

	error = log.New(io.MultiWriter(os.Stderr, os.Stdout),
		"ERROR:",
		log.Ldate|log.Ltime|log.Lshortfile)

	waring = log.New(os.Stdout,
		"WARING:",
		log.Ldate|log.Lmicroseconds|log.Llongfile)
}

func Info(message string) {
	info.Println(message)
}

func Error(message string) {
	error.Println(message)
}

func Waring(message string) {
	waring.Println(message)
}

func Trace(message string) {
	trace.Println(message)
}
