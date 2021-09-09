package async

import "sync"

type Task struct {
	Parent *Task
	mu     *sync.Mutex
	wg     *sync.WaitGroup
	action func()
	err    error
}

func NewTask(action func()) *Task {
	return &Task{
		Parent: nil,
		mu:     &sync.Mutex{},
		wg:     &sync.WaitGroup{},
		action: action,
	}
}

func (t *Task) Wait() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.wg.Wait()
}

func (t *Task) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.wg = &sync.WaitGroup{}
	t.wg.Add(1)
	go func() {
		defer func() {
			if err, ok := recover().(error); ok {
				t.err = err
			}
			t.wg.Done()
		}()
		if t.Parent != nil {
			t.Parent.Start()
			t.Parent.Wait()
		}
		t.action()

	}()
}

func (t *Task) ContinueWith(action func()) *Task {
	return &Task{
		Parent: t,
		mu:     &sync.Mutex{},
		wg:     &sync.WaitGroup{},
		action: action,
	}
}
