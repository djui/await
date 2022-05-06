package main

import (
	"log"
	"os"
)

// An enumeration of log levels.
const (
	debugLevel = iota
	infoLevel
	errorLevel
	silentLevel
)

// LevelLogger defines a logger which is aware of log levels and only logs
// messages when below or equal to a given log level.
type LevelLogger struct {
	*log.Logger
	level int
}

// NewLogger instantiates a new LevelLogger with a given log level.
func NewLogger(level int) *LevelLogger {
	return &LevelLogger{
		Logger: log.New(os.Stderr, "", log.LstdFlags),
		level:  level,
	}
}

// Debug logs a message if the current log level is below or equal to
// debugLevel.
func (l *LevelLogger) Debug(v ...interface{}) {
	if l.level <= debugLevel {
		log.Print(v...)
	}
}

// Debugln logs a message if the current log level is below or equal to
// debugLevel.
func (l *LevelLogger) Debugln(v ...interface{}) {
	if l.level <= debugLevel {
		log.Println(v...)
	}
}

// Debugf logs a message if the current log level is below or equal to
// debugLevel.
func (l *LevelLogger) Debugf(format string, v ...interface{}) {
	if l.level <= debugLevel {
		log.Printf(format, v...)
	}
}

// Info logs a message if the current log level is below or equal to infoLevel.
func (l *LevelLogger) Info(v ...interface{}) {
	if l.level <= infoLevel {
		log.Print(v...)
	}
}

// Infoln logs a message if the current log level is below or equal to
// infoLevel.
func (l *LevelLogger) Infoln(v ...interface{}) {
	if l.level <= infoLevel {
		log.Println(v...)
	}
}

// Infof logs a message if the current log level is below or equal to infoLevel.
func (l *LevelLogger) Infof(format string, v ...interface{}) {
	if l.level <= infoLevel {
		log.Printf(format, v...)
	}
}

// Error logs a message if the current log level is below or equal to
// errorLevel.
func (l *LevelLogger) Error(v ...interface{}) {
	if l.level <= errorLevel {
		log.Print(v...)
	}
}

// Errorln logs a message if the current log level is below or equal to
// errorLevel.
func (l *LevelLogger) Errorln(v ...interface{}) {
	if l.level <= errorLevel {
		log.Println(v...)
	}
}

// Errorf logs a message if the current log level is below or equal to
// errorLevel.
func (l *LevelLogger) Errorf(format string, v ...interface{}) {
	if l.level <= errorLevel {
		log.Printf(format, v...)
	}
}

// Fatal logs a message and exit the process.
func (l *LevelLogger) Fatal(v ...interface{}) {
	l.Error(v...)
	os.Exit(1)
}

// Fatalln logs a message and exit the process.
func (l *LevelLogger) Fatalln(v ...interface{}) {
	l.Errorln(v...)
	os.Exit(1)
}

// Fatalf logs a message and exit the process.
func (l *LevelLogger) Fatalf(format string, v ...interface{}) {
	l.Errorf(format, v...)
	os.Exit(1)
}
