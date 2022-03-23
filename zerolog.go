package abstractlogger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZerologLogger(level Level) *ZerologLogger {
	return &ZerologLogger{
		levelCheck: NewLevelCheck(level),
	}
}

// ZerologLogger implements the Logging frontend using the popular logging backend log
// It uses the LevelCheck helper to increase performance.
type ZerologLogger struct {
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
		debug := log.Debug()
		for _, e := range v {
			debug.Interface("", e)
		}
		debug.Send()
	case InfoLevel:
		info := log.Info()
		for _, e := range v {
			info.Interface("", e)
		}
		info.Send()
	case WarnLevel:
		warn := log.Debug()
		for _, e := range v {
			warn.Interface("", e)
		}
		warn.Send()
	case ErrorLevel:
		debug := log.Debug()
		for _, e := range v {
			debug.Interface("", e)
		}
		debug.Send()
	case FatalLevel:
		fatal := log.Debug()
		for _, e := range v {
			fatal.Interface("", e)
		}
		fatal.Send()
	case PanicLevel:
		panic := log.Debug()
		for _, e := range v {
			panic.Interface("", e)
		}
		panic.Send()
	}
}

func (z *ZerologLevelLogger) Printf(format string, v ...interface{}) {
	switch z.level {
	case DebugLevel:
		log.Debug().Msgf(format, v...)
	case InfoLevel:
		log.Info().Msgf(format, v...)
	case WarnLevel:
		log.Warn().Msgf(format, v...)
	case ErrorLevel:
		log.Error().Msgf(format, v...)
	case FatalLevel:
		log.Fatal().Msgf(format, v...)
	case PanicLevel:
		log.Panic().Msgf(format, v...)
	}
}
