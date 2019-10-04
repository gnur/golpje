package main

import (
	"github.com/sirupsen/logrus"
)

// AddSourceHook is struct only used for adding a source
type AddSourceHook struct {
}

// Levels needs to implemented
func (h *AddSourceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire needs to implemented
func (h *AddSourceHook) Fire(e *logrus.Entry) error {
	e.Data["source"] = "booksing"
	return nil
}
