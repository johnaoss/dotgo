// Package foobar is an example of a plugin.
package foobar

import (
	"github.com/johnaoss/dotgo/plugins"
)

type TT struct {
	plugins.Plugin
}

func NewTT() plugins.Plugin {
	return &TT{}
}

func (t *TT) CanHandle(directive string) bool {
	return true
}

// todo
func (t *TT) Handle(directive string, data map[string]interface{}) error {
	return nil
}

func GetPlugins() []plugins.Plugin {
	pl := make([]plugins.Plugin, 1)
	pl[0] = NewTT()
	return pl
}
