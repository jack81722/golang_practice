package rx

type Observable interface {
	Subscribe(Observer)
}

type Observer interface {
	OnError(error)
	OnNext(interface{})
	OnCompleted()
}

type just struct {
	v interface{}
}

func Just(value interface{}) *just {
	return &just{value}
}

func (j *just) Subscribe(o Observer) {
	defer func() {
		if _err := recover(); _err != nil {
			if err, ok := _err.(error); ok && err != nil {
				o.OnError(err)
			}
		}
	}()
	defer o.OnCompleted()
	o.OnNext(j.v)

}

type Subject struct {
	NextHandler     func(interface{})
	ErrorHandler    func(error)
	CompleteHandler func()
}

func (s *Subject) OnError(err error) {
	if s.ErrorHandler != nil {
		s.ErrorHandler(err)
	}
}

func (s *Subject) OnCompleted() {
	if s.CompleteHandler != nil {
		s.CompleteHandler()
	}
}

func (s *Subject) OnNext(obj interface{}) {
	if s.NextHandler != nil {
		s.NextHandler(obj)
	}
}
