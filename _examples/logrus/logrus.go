package main

import (
	"github.com/mattfarina/log-go"
	"github.com/mattfarina/log-go/impl/logrus"
)

func main() {
	log.Current = logrus.NewStandard()

	// A basic message
	log.Info("Hello,", "World")

	// Use Go formatting on a warning
	log.Warnf("Foo %s", "bar")

	// An error with context
	log.Errorw("foo, bar", log.Fields{"baz": "qux"})
}
