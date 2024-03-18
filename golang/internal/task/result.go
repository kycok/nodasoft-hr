package task

import (
	"fmt"
	"sync"

	"nodasoft-hr/internal/model"
)

type Result struct {
	Done    chan model.Ttype
	Undone  chan error
	wg      *sync.WaitGroup
	Success []model.Ttype
	Fail    []error
}

func NewResult() *Result {
	return &Result{
		Done:   make(chan model.Ttype),
		Undone: make(chan error),
		wg:     &sync.WaitGroup{},
	}
}

func (t *Result) Add(task model.Ttype, failed bool) {
	switch failed {
	case false:
		t.Done <- task
	case true:
		t.Undone <- fmt.Errorf("id=%d timeWork=%s result=error message=%s", task.GetId(), task.GetFT(), task.GetTaskRESULT())
	}
}

func (t *Result) Collect() {
	t.Success = make([]model.Ttype, 0, len(t.Done))
	t.Fail = make([]error, 0, len(t.Undone))

	t.wg.Add(2)

	go func() {
		defer t.wg.Done()
		for r := range t.Done {
			t.Success = append(t.Success, r)
		}
	}()

	go func() {
		defer t.wg.Done()
		for r := range t.Undone {
			t.Fail = append(t.Fail, r)
		}
	}()

	t.wg.Wait()
}

func (t *Result) Print() {
	println("Done tasks:")
	for _, r := range t.Success {
		fmt.Println(r)
	}

	println("\nErrors:")
	for _, e := range t.Fail {
		fmt.Println(e.Error())
	}
}

func (t *Result) Close() {
	close(t.Done)
	close(t.Undone)
}
