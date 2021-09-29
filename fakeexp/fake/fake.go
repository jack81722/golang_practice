package fake

import (
	"reflect"
	"runtime"
)

type Fake struct {
	methods map[string]interface{}
}

func NewFake() *Fake {
	return &Fake{
		methods: make(map[string]interface{}),
	}
}

func (f *Fake) FakeReturn(method interface{}, result interface{}) {
	f.methods[getFuncName(method)] = result
}

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
