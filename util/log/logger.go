//nolint
package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bluele/go-timecop"
)

type (
	// Logger は、ログ出力に関する設定値を保持する
	Logger struct {
		typ        string
		level      uint32
		fields     Fields
		output     io.Writer
		timeFormat string
		levels     []string
		mutex      sync.Mutex
	}

	// Fields は、ログに付与する付加情報
	Fields map[string]interface{}
)

// Lvl は、ログレベルを定義する
type Lvl uint8

// ログレベル
const (
	DEBUG Lvl = iota
	INFO
	WARN
	ERROR
	FATAL
)

var global *Logger

func init() {
	global = New("-")
}

// OsExit は os.Exit のポインタ
var OsExit = os.Exit

// New は、初期化したロガーを返却する
func New(typ string) (l *Logger) {
	l = &Logger{
		level:  uint32(INFO),
		typ:    typ,
		fields: make(map[string]interface{}),
	}
	l.initLevels()
	l.SetOutput(os.Stdout)
	l.SetTimeFormat(time.RFC3339Nano)
	return
}

func (l *Logger) initLevels() {
	l.levels = []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
	}
}

// Typ は、設定されたログ種別を返却する
func (l *Logger) Typ() string {
	return l.typ
}

// SetTyp は、ログ種別を設定する
func (l *Logger) SetTyp(p string) {
	l.typ = p
}

// Fields は、設定されたデータフィールドを返却する
func (l *Logger) Fields() map[string]interface{} {
	return l.fields
}

// AddField は、データフィールドを追加する
func (l *Logger) AddField(key string, value interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.fields[key] = value
}

// TimeFormat は、タイムフィールドの書式を返却する
func (l *Logger) TimeFormat() string {
	return l.timeFormat
}

// SetTimeFormat は、タイムフィールドの書式を設定する
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// Level は、ログレベルを返却する
func (l *Logger) Level() Lvl {
	return Lvl(atomic.LoadUint32(&l.level))
}

// SetLevel は、ログレベルを設定する
func (l *Logger) SetLevel(level Lvl) {
	atomic.StoreUint32(&l.level, uint32(level))
}

// Output は、出力先を返却する
func (l *Logger) Output() io.Writer {
	return l.output
}

// SetOutput は、出力先を設定する
func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
}

// Debug output log message
func (l *Logger) Debug(v ...interface{}) {
	l.log(DEBUG, fmt.Sprint(v...))
}

// Debugf output log message
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(DEBUG, escape(fmt.Sprintf(format, v...)))
}

// Info output log message
func (l *Logger) Info(v ...interface{}) {
	l.log(INFO, fmt.Sprint(v...))
}

// Infof output log message
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(INFO, escape(fmt.Sprintf(format, v...)))
}

// Warn output log message
func (l *Logger) Warn(v ...interface{}) {
	l.log(WARN, fmt.Sprint(v...))
}

// Warnf output log message
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(WARN, escape(fmt.Sprintf(format, v...)))
}

// Error output log message
func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, fmt.Sprint(v...))
}

// Errorf output log message
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(ERROR, escape(fmt.Sprintf(format, v...)))
}

// Fatal output log message
func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, fmt.Sprint(v...))
	OsExit(1)
}

// Fatalf output log message
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(FATAL, escape(fmt.Sprintf(format, v...)))
	OsExit(1)
}

// Typ output log message
func Typ() string {
	return global.Typ()
}

// SetTyp output log message
func SetTyp(p string) {
	global.SetTyp(p)
}

// Level は、グローバルログのレベルを返却する
func Level() Lvl {
	return global.Level()
}

// SetLevel は、グローバルログのレベルを設定する
func SetLevel(level Lvl) {
	global.SetLevel(level)
}

// Output は、ログの出力先を返却する
func Output() io.Writer {
	return global.Output()
}

// SetOutput は、ログの出力先を指定する
func SetOutput(w io.Writer) {
	global.SetOutput(w)
}

// Debug は、DEBUG レベルのログを出力する
func Debug(i ...interface{}) {
	global.Debug(i...)
}

// Debugf は、DEBUG レベルのログを出力する
func Debugf(format string, args ...interface{}) {
	global.Debugf(format, args...)
}

// Info は、INFO レベルのログを出力する
func Info(i ...interface{}) {
	global.Info(i...)
}

// Infof は、INFO レベルのログを出力する
func Infof(format string, args ...interface{}) {
	global.Infof(format, args...)
}

// Warn は、WARN レベルのログを出力する
func Warn(i ...interface{}) {
	global.Warn(i...)
}

// Warnf は、WARN レベルのログを出力する
func Warnf(format string, args ...interface{}) {
	global.Warnf(format, args...)
}

// Error は、ERROR レベルのログを出力する
func Error(i ...interface{}) {
	global.Error(i...)
}

// Errorf は、ERROR レベルのログを出力する
func Errorf(format string, args ...interface{}) {
	global.Errorf(format, args...)
}

// Fatal は、FATAL レベルのログを出力する
func Fatal(i ...interface{}) {
	global.Fatal(i...)
}

// Fatalf は、FATAL レベルのログを出力する
func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

func (l *Logger) log(level Lvl, msg string) {
	if level < l.Level() {
		return
	}
	time := timecop.Now().Format(l.TimeFormat())
	fields := map[string]interface{}{
		"type":  l.typ,
		"level": l.levels[level],
		"time":  time,
	}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields["message"] = msg

	b, _ := json.Marshal(fields)

	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.output.Write(b)
	l.output.Write([]byte("\n"))
}

// escape as json string value
func escape(s string) string {
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	return s
}
