//nolint
package log

import (
	"github.com/gocraft/dbr/v2"
)

// EventReceiver gets events from dbr methods for logging purposes.
// @see https://github.com/gocraft/dbr/blob/master/event.go
type EventReceiver struct {
	dbr.EventReceiver
}

// NewEventReceiver initialize
func NewEventReceiver() *EventReceiver {
	return &EventReceiver{}
}

// Event receives a simple notification when various events occur.
func (e *EventReceiver) Event(eventName string) {
}

// EventKv receives a notification when various events occur along with
// optional key/value data.
func (e *EventReceiver) EventKv(eventName string, kvs map[string]string) {
	e.writeLog(kvs)
}

// EventErr receives a notification of an error if one occurs.
func (e *EventReceiver) EventErr(eventName string, err error) error {
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data.
func (e *EventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	e.writeLog(kvs)
	return err
}

// Timing receives the time an event took to happen.
func (e *EventReceiver) Timing(eventName string, nanoseconds int64) {
}

// TimingKv receives the time an event took to happen along with optional key/value data.
func (e *EventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	e.writeLog(kvs)
}

func (e *EventReceiver) writeLog(kvs map[string]string) {
	for _, v := range kvs {
		Infof("query: %s", v)
	}
}
