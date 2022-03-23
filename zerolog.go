package abstractlogger

import (
	ll "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZerologLogger(zerologLogger *zerolog.Event, level Level) *ZerologLogger {
	return &ZerologLogger{
		l:          zerologLogger,
		levelCheck: NewLevelCheck(level),
	}
}

// ZerologLogger implements the Logging frontend using the popular logging backend log
// It uses the LevelCheck helper to increase performance.
type ZerologLogger struct {
	l          *zerolog.Event
	levelCheck LevelCheck
}

func (z *ZerologLogger) LevelLogger(level Level) LevelLogger {
	return &ZerologLevelLogger{
		level: level,
	}
}

func field(prev *zerolog.Event, field Field) *zerolog.Event {
	switch field.kind {
	case StringField:
		return prev.Str(field.key, field.stringValue)
	case IntField:
		return prev.Int(field.key, int(field.intValue))
	case BoolField:
		return prev.Bool(field.key, field.intValue != 0)
	case ByteStringField:
		return prev.RawJSON(field.key, field.byteValue)
	case ErrorField:
		return prev.Err(field.errorValue)
	case NamedErrorField:
		return prev.AnErr(field.key, field.errorValue)
	case StringsField:
		return prev.Strs(field.key, field.stringsValue)
	default:
		return prev.Interface(field.key, field.interfaceValue)
	}
}

func fieldsfn(init *zerolog.Event, fields []Field) *zerolog.Event {
	for i := range fields {
		field(init, fields[i])
	}
	return init
}

func (z *ZerologLogger) Debug(msg string, fields ...Field) {
	if !z.levelCheck.Check(DebugLevel) {
		return
	}
	debug := log.Debug()
	fieldsfn(debug, fields).Msg(msg)
}

func (z *ZerologLogger) Info(msg string, fields ...Field) {
	if !z.levelCheck.Check(InfoLevel) {
		return
	}
	info := log.Info()
	fieldsfn(info, fields).Msg(msg)
}

func (z *ZerologLogger) Warn(msg string, fields ...Field) {
	if !z.levelCheck.Check(WarnLevel) {
		return
	}
	warn := log.Warn()
	fieldsfn(warn, fields).Msg(msg)
}

func (z *ZerologLogger) Error(msg string, fields ...Field) {
	if !z.levelCheck.Check(ErrorLevel) {
		return
	}
	err := log.Error()
	fieldsfn(err, fields).Msg(msg)
}

func (z *ZerologLogger) Fatal(msg string, fields ...Field) {
	if !z.levelCheck.Check(FatalLevel) {
		return
	}
	fatal := log.Fatal()
	fieldsfn(fatal, fields).Msg(msg)
}

func (z *ZerologLogger) Panic(msg string, fields ...Field) {
	if !z.levelCheck.Check(PanicLevel) {
		return
	}
	panic := log.Panic()
	fieldsfn(panic, fields).Msg(msg)
}

type ZerologLevelLogger struct {
	level Level
}

func (z *ZerologLevelLogger) Println(v ...interface{}) {
	switch z.level {
	case DebugLevel:
		ll.Println(v)
	case InfoLevel:
		ll.Println(v)
	case WarnLevel:
		ll.Println(v)
	case ErrorLevel:
		ll.Println(v)
	case FatalLevel:
		ll.Println(v)
	case PanicLevel:
		ll.Println(v)
	}
}

func (z *ZerologLevelLogger) Printf(format string, v ...interface{}) {
	switch z.level {
	case DebugLevel:
		ll.Printf(format, v)
	case InfoLevel:
		ll.Printf(format, v)
	case WarnLevel:
		ll.Printf(format, v)
	case ErrorLevel:
		ll.Printf(format, v)
	case FatalLevel:
		ll.Printf(format, v)
	case PanicLevel:
		ll.Printf(format, v)
	}
}
