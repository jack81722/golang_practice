package fake

func (f *Fake) Exp() string {
	method, ok := f.methods[getFuncName(f.Exp)]
	if !ok {
		return ""
	}
	_func, ok := method.(func() string)
	if !ok {
		return ""
	}
	return _func()
}
