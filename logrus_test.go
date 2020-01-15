package abstractlogger

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"testing"
)

func logrusLogger(out io.Writer,level logrus.Level) *logrus.Logger {
	return &logrus.Logger{
		Out:       out,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}
}

func TestLogrusLogger(t *testing.T){

	level := logrus.DebugLevel

	var directOut bytes.Buffer
	direct := logrusLogger(&directOut,level)

	var wrappedOut bytes.Buffer
	wrapped := logrusLogger(&wrappedOut,level)
	indirect := NewLogrusLogger(wrapped,DebugLevel)

	direct.WithField("foo","bar").Debug("baz")
	indirect.Debug("baz",String("foo","bar"))

	direct.WithField("foo","bar").Info("baz")
	indirect.Info("baz",String("foo","bar"))

	direct.WithField("foo","bar").Warn("baz")
	indirect.Warn("baz",String("foo","bar"))

	direct.WithField("foo","bar").Error("baz")
	indirect.Error("baz",String("foo","bar"))

	if directOut.String() != wrappedOut.String() {
		t.Fatal("must be the same")
	}
}