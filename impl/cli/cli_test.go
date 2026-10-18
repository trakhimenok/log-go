package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Masterminds/log-go"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	// Test the logger meets the interface
	var _ log.Logger = new(Logger)

	black := color.New(color.FgBlack)

	buf := &bytes.Buffer{}
	lgr := &Logger{
		TraceOut:   buf,
		DebugOut:   buf,
		InfoOut:    buf,
		WarnOut:    buf,
		ErrorOut:   buf,
		PanicOut:   buf,
		FatalOut:   buf,
		Level:      log.TraceLevel,
		TraceColor: black,
		DebugColor: black,
		WarnColor:  black,
		ErrorColor: black,
		PanicColor: black,
		FatalColor: black,
	}

	lgr.Trace("test trace")
	if !strings.Contains(buf.String(), `test trace`) {
		t.Log(buf.String())
		t.Error("cli trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracef("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Error("cli trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracew("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli trace not logging correctly")
	}
	buf.Reset()

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `test debug`) {
		t.Log(buf.String())
		t.Error("cli debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Error("cli debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli debug not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `test info`) {
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `test warn`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `test error`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli error not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "PANIC: test panic\n", func() {
		lgr.Panic("test panic")
	})
	if !strings.Contains(buf.String(), `test panic`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "PANIC: Hello World\n", func() {
		lgr.Panicf("Hello %s", "World")
	})
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "PANIC: foo bar baz=qux \n", func() {
		lgr.Panicw("foo bar", log.Fields{"baz": "qux"})
	})
	if !strings.Contains(buf.String(), `foo bar baz=qux`) {
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
	if !strings.Contains(buf.String(), `FATAL: test fatal`) {
		t.Log(buf.String())
		t.Error("cli fatal not logging correctly")
	}
	buf.Reset()
	exitCode = 0

	lgr.Fatalf("Hello %s", "World")
	assert.Equal(t, 1, exitCode)
	if !strings.Contains(buf.String(), `FATAL: Hello World`) {
		t.Log(buf.String())
		t.Error("cli fatalf not logging correctly")
	}
	buf.Reset()
	exitCode = 0

	lgr.Fatalw("foo bar", log.Fields{"baz": "qux"})
	assert.Equal(t, 1, exitCode)
	if !strings.Contains(buf.String(), `FATAL: foo bar baz=qux`) {
		t.Log(buf.String())
		t.Error("cli fatalw not logging correctly")
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

func TestNewStandard(t *testing.T) {
	lgr := NewStandard()
	assert.NotNil(t, lgr)
	assert.Equal(t, log.InfoLevel, lgr.Level)
	assert.NotNil(t, lgr.TraceOut)
	assert.NotNil(t, lgr.DebugOut)
	assert.NotNil(t, lgr.InfoOut)
	assert.NotNil(t, lgr.WarnOut)
	assert.NotNil(t, lgr.ErrorOut)
	assert.NotNil(t, lgr.PanicOut)
	assert.NotNil(t, lgr.FatalOut)
	assert.NotNil(t, lgr.TraceColor)
	assert.NotNil(t, lgr.DebugColor)
	assert.NotNil(t, lgr.WarnColor)
	assert.NotNil(t, lgr.ErrorColor)
	assert.NotNil(t, lgr.PanicColor)
	assert.NotNil(t, lgr.FatalColor)
}

func TestCheckEnding(t *testing.T) {
	assert.Equal(t, "hello\n", checkEnding("hello"))
	assert.Equal(t, "hello\n", checkEnding("hello\n"))
	assert.Equal(t, "\n", checkEnding(""))
}
