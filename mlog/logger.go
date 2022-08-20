package mlog

import (
	"fmt"
	"log"
	"os"
)

var stdLog, errLog *log.Logger

func InitLoggers() {
	stdLog = log.New(os.Stdout, "[INFO] ", 0)
	stdLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	errLog = log.New(os.Stderr, "[ERROR] ", 1)
	errLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	stdLog.Output(2, fmt.Sprintf(format, v...))
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) { stdLog.Output(2, fmt.Sprint(v...)) }

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) { stdLog.Output(2, fmt.Sprintln(v...)) }

func PrintErrln(v ...interface{}) {
	stdLog.Output(2, fmt.Sprintln(v...))
	errLog.Output(2, fmt.Sprintln(v...))
}

func PrintErrf(format string, v ...interface{}) {
	stdLog.Output(2, fmt.Sprintf(format, v...))
	errLog.Output(2, fmt.Sprintf(format, v...))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	stdLog.Output(2, fmt.Sprint(v...))
	errLog.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	errLog.Output(2, fmt.Sprintf(format, v...))
	stdLog.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	stdLog.Output(2, fmt.Sprintln(v...))
	errLog.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	stdLog.Output(2, s)
	errLog.Output(2, s)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	stdLog.Output(2, s)
	errLog.Output(2, s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	stdLog.Output(2, s)
	errLog.Output(2, s)
	panic(s)
}
