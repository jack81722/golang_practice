package panichandler

import (
	"fmt"
	"runtime"
	"strings"
)

func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		// fmt.Println(name)
		// skip until not runtime func
		// means that the first panic calling
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}

func RecoverPanic() {
	r := recover()
	if r == nil {
		return
	}
	fmt.Println(identifyPanic())
}
