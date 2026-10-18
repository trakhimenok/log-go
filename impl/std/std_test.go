package std

import (
	"bytes"
	stdlog "log"
	"strings"
	"testing"

	"github.com/Masterminds/log-go"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	// Test the logger meets the interface
	var _ log.Logger = new(Logger)

	var (
		buf    bytes.Buffer
		logger = stdlog.New(&buf, "", stdlog.Lshortfile)
	)
	lgr := New(logger)

	// Make sure levels are working
	lgr.Debug("test debug")
	if strings.Contains(buf.String(), `[DEBUG]   test debug`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	// Test all levels
	lgr.Level = log.TraceLevel

	lgr.Trace("test trace")
	if !strings.Contains(buf.String(), `[TRACE]   test trace`) {
		t.Log(buf.String())
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracef("Hello %s", "World")
	if !strings.Contains(buf.String(), `[TRACE]   Hello World`) {
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracew("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[TRACE]   foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `[DEBUG]   test debug`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[DEBUG]   Hello World`) {
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[DEBUG]   foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `[INFO]    test info`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `[INFO]    Hello World`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[INFO]    foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `[WARNING] test warn`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[WARNING] Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[WARNING] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `[ERROR]   test error`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[ERROR]   Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[ERROR]   foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC]   test panic", func() {
		lgr.Panic("test panic")
	})
	if !strings.Contains(buf.String(), `test panic`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC]   Hello World", func() {
		lgr.Panicf("Hello %s", "World")
	})
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC]   foo bar [baz=qux]", func() {
		lgr.Panicw("foo bar", log.Fields{"baz": "qux"})
	})
	if !strings.Contains(buf.String(), `foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	var exitCode int
	origExit := exit
	exit = func(code int) {
		exitCode = code
	}
	defer func() {
		exit = origExit
	}()

	lgr.Fatal("test fatal")
	assert.Equal(t, 1, exitCode)
	if !strings.Contains(buf.String(), `[FATAL]   test fatal`) {
		t.Log(buf.String())
		t.Error("stdlib fatal not logging correctly")
	}
	buf.Reset()
	exitCode = 0

	lgr.Fatalf("Hello %s", "World")
	assert.Equal(t, 1, exitCode)
	if !strings.Contains(buf.String(), `[FATAL]   Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib fatalf not logging correctly")
	}
	buf.Reset()
	exitCode = 0

	lgr.Fatalw("foo bar", log.Fields{"baz": "qux"})
	assert.Equal(t, 1, exitCode)
	if !strings.Contains(buf.String(), `[FATAL]   foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib fatalw not logging correctly")
	}
	buf.Reset()
	exitCode = 0

	// Test level above FatalLevel skips Fatal logging
	lgr.Level = log.FatalLevel + 1
	lgr.Fatal("skipped fatal")
	lgr.Fatalf("skipped %s", "fatal")
	lgr.Fatalw("skipped fatal", log.Fields{"baz": "qux"})
	assert.Equal(t, 0, exitCode)
	assert.Empty(t, buf.String())
}
