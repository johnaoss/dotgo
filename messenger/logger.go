package messenger

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"sync"
)

type Color string

const (
	None    Color = ""
	Reset   Color = "\033[0m"
	Red     Color = "\033[91m"
	Green   Color = "\033[92m"
	Yellow  Color = "\033[93m"
	Blue    Color = "\033[94m"
	Magenta Color = "\033[95m"
)

type Level int8

const (
	NOTSET  Level = 0
	DEBUG   Level = 10
	LOWINFO Level = 15
	INFO    Level = 20
	WARNING Level = 30
	ERROR   Level = 40
)

// Our singleton for the messenger.
var msg = &messenger{
	level:    LOWINFO,
	useColor: true,

	out: os.Stdout,
}

// Messenger returns the package default instance of the logger.
func Messenger() *messenger {
	return msg
}

func New(lvl Level, useColor bool) *messenger {
	return &messenger{
		level:    lvl,
		useColor: useColor,
		out:      os.Stdout,
	}
}

// Messenger is a copy of the messenger package in the O.G. dotbot.
type messenger struct {
	level    Level
	useColor bool

	mu sync.RWMutex

	// Used for testing
	out io.Writer
}

// SetLevel sets the output level of the logger.
func (m *messenger) SetLevel(l Level) {
	m.mu.Lock()
	m.level = l
	m.mu.Unlock()
}

func (m *messenger) Level() Level {
	m.mu.RLock()
	l := m.level
	m.mu.RUnlock()
	return l
}

// SetUseColor sets whether or not this logger should colorify output.
func (m *messenger) SetUseColor(b bool) {
	m.mu.Lock()
	m.useColor = b
	m.mu.Unlock()
}

// UseColor returns whether or not this logger can use color for output.
func (m *messenger) UseColor() bool {
	m.mu.RLock()
	b := m.useColor
	m.mu.RUnlock()
	return b
}

// bug: this only works if we use os.Stdout as an output.
func (m *messenger) getColor(lvl Level) Color {
	shouldColor := m.UseColor()

	if lvl < DEBUG || !shouldColor || !terminal.IsTerminal(int(os.Stdout.Fd())) {
		return ""
	}

	switch {
	case DEBUG <= lvl && lvl < LOWINFO:
		return Yellow
	case LOWINFO <= lvl && lvl < INFO:
		return Blue
	case INFO <= lvl && lvl < WARNING:
		return Green
	case WARNING <= lvl && lvl < ERROR:
		return Magenta
	case ERROR <= lvl:
		return Red
	default:
		return ""
	}
}

// reset returns the escape sequence for a reset colour.
func (m *messenger) reset() Color {
	if !m.UseColor() {
		return ""
	}
	return Reset
}

func (m *messenger) Log(lvl Level, str string) {
	if lvl >= m.Level() {
		fmt.Fprintf(m.out, "%s%s%s\n", m.getColor(lvl), str, m.reset())
	}
}

func (m *messenger) Debug(str string) {
	m.Log(DEBUG, str)
}

func (m *messenger) LowInfo(str string) {
	m.Log(LOWINFO, str)
}

func (m *messenger) Info(str string) {
	m.Log(INFO, str)
}

func (m *messenger) Warn(str string) {
	m.Log(WARNING, str)
}

func (m *messenger) Error(str string) {
	m.Log(ERROR, str)
}
