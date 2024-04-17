package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kiyor/env"
	"github.com/kiyor/terminal/color"
)

var verbose bool

func init() {
	env.BoolVar(&verbose, "VERBOSE", false, "verbose log")
}

const (
	Info    = "info"
	Error   = "error"
	Warning = "warning"
	Success = "success"
	Debug   = "info"
)

type Logger struct {
	Info    ILogger
	Success ILogger
	Warn    ILogger
	Error   ILogger
	Debug   ILogger
}

type ILogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
}

type DefaultLogger struct {
	*log.Logger
	name  string
	level string
}

func NewDefaultLog(name string, flag ...int) *Logger {
	prefix := name
	if len(name) > 0 {
		if !strings.HasSuffix(name, "-") {
			prefix = name + "-"
		}
	}
	var lflag int
	if len(flag) == 0 {
		lflag = log.Lshortfile | log.Lshortfile
	} else {
		lflag = flag[0]
	}

	info := NewDefaultLogger(name, "|", Info, os.Stderr, lflag)
	success := NewDefaultLogger(name, "g", Success, os.Stderr, lflag)
	warn := NewDefaultLogger(name, "y", Warning, os.Stderr, lflag)
	errlog := NewDefaultLogger(name, "r", Error, os.Stderr, lflag)
	debug := io.Discard
	if verbose {
		debug = os.Stderr
	}
	return &Logger{
		Info:    info,
		Success: success,
		Debug:   log.New(debug, color.Sprint("@{c}["+prefix+"debug]@{|} "), lflag),
		Warn:    warn,
		Error:   errlog,
	}
}

func NewDefaultLogger(name, colorCode, level string, output *os.File, flag int) ILogger {
	l := &DefaultLogger{
		name:   name,
		level:  level,
		Logger: log.New(output, color.Sprint("@{"+colorCode+"}["+strings.Join([]string{name, level}, "-")+"]@{|} "), flag),
	}
	return l
}

func (l *DefaultLogger) Output(calldepth int, s string) error {
	return l.Logger.Output(calldepth, s)
}

func (l *DefaultLogger) Print(v ...interface{}) {
	l.Output(3, fmt.Sprint(v...))
}

func (l *DefaultLogger) Printf(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Println(v ...interface{}) {
	l.Output(3, fmt.Sprintln(v...))
}

func (l *DefaultLogger) SetFlags(flag int) {
	l.Logger.SetFlags(flag)
}
func (l *DefaultLogger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}
func (l *DefaultLogger) SetPrefix(prefix string) {
	l.Logger.SetPrefix(prefix)
}
