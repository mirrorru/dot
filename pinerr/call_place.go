package pinerr

import (
	"runtime"
)

type callPlace struct {
	fileName string
	line     int
}

func getCallPlace(skip int) callPlace {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return callPlace{"-", 0}
	}

	f := runtime.FuncForPC(pc)
	name := f.Name()
	_, line := f.FileLine(pc)

	return callPlace{name, line}
}

func (cp *callPlace) empty() bool {
	return len(cp.fileName) == 0
}
