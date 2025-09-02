package dot

import (
	"runtime"
)

func GetCallPlace(skip int) (file string, line int) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "-", 0
	}

	f := runtime.FuncForPC(pc)
	file = f.Name()
	_, line = f.FileLine(pc)

	return file, line
}
