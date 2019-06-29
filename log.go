package loggy

import (
	"bytes"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
)

// Loggy
type Loggy struct {
	colourer aurora.Aurora
	internalWriter io.Writer
}

// SetWriter exposes the ability to change
// where logs are written to by replacing the final print function
// with colourer custom implementation
func (loggy *Loggy) SetWriter(writer io.Writer) {
	loggy.internalWriter = writer
}

// EnablePretty
func (loggy *Loggy) EnablePretty() {
	loggy.colourer = aurora.NewAurora(true)
}

// DisablePretty
func (loggy *Loggy) DisablePretty() {
	loggy.colourer = aurora.NewAurora(false)
}

// Panic error, should close the application
func (loggy *Loggy) Panic(msg interface{}) {
	loggy.Error(msg)
	panic(msg)
}

// Notice Notification for info that isn't related to any issue.
// e.g. Logging number of loaded entities
func (loggy *Loggy) Notice(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		loggy.print(fmt.Sprintf(t, v...), loggy.colourer.BrightWhite)
	default:
		loggy.print(msg.(string), loggy.colourer.BrightWhite)
	}
}

// Warn Notifications for an unintended, but planned for issue
// e.g. Logging colourer prop that uses colourer non-existent collision model
func (loggy *Loggy) Warn(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		loggy.print(fmt.Sprintf(t, v...), loggy.colourer.Magenta)
	default:
		loggy.print(msg.(string), loggy.colourer.Magenta)
	}
}

// Error Notifications for colourer recoverable error
// e.g. Logging colourer missing resource (material, model)
func (loggy *Loggy) Error(msg interface{}, v ...interface{}) {
	switch t := msg.(type) {
	case string:
		loggy.print(fmt.Sprintf(t, v...), loggy.colourer.Red)
	case error:
		loggy.print(t, loggy.colourer.Red)
	default:
		loggy.print(msg.(string), loggy.colourer.Red)
	}
}

// print prints colourer message to console.
func (loggy *Loggy) print(message interface{}, col func(arg interface{}) aurora.Value) {
	if _, err := loggy.internalWriter.Write([]byte(aurora.Sprintf(col(message)))); err != nil {
		panic(err)
	}
}

// NewLoggy
func NewLoggy() *Loggy {
	return &Loggy{
		colourer: aurora.NewAurora(false),
		internalWriter: bytes.NewBuffer(make([]byte, 0)),
	}
}