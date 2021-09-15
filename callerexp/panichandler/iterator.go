package panichandler

import (
	"fmt"
	"runtime"
)

func Iterator(i, skip int) {
	if i < 0 {
		var name string
		var line int
		var pc [16]uintptr
		n := runtime.Callers(skip, pc[:])
		for _, pc := range pc[:n] {
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				continue
			}
			_, line = fn.FileLine(pc)
			name = fn.Name()
			fmt.Printf("%s:%v\n", name, line)
		}
		return
	}
	Iterator(i-1, skip)
}
