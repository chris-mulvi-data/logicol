package logicol_test

import (
	"github.com/chris-mulvi-data/logicol"
	"testing"
)

func TestWrite(t *testing.T) {
	logicol.Write(logicol.Info, "Hello, World!")
	logicol.Write(logicol.Debug, "Hello, World!")
	logicol.Write(logicol.Warning, "Hello, World!")
	logicol.Write(logicol.Error, "Hello, World!")
	logicol.Write(logicol.Event, "Hello, World!")
}
