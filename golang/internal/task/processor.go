package task

import (
	"sync"
	"time"

	"nodasoft-hr/internal/model"
)

type Processor struct {
	wg           sync.WaitGroup
	WorkersCount int
	TaskTimeout  int
}

func (t *Processor) Run(jobs <-chan model.Ttype) *Result {
	tResult := NewResult()

	t.wg.Add(t.WorkersCount)

	for i := 0; i < t.WorkersCount; i++ {
		go func() {
			defer t.wg.Done()
			for task := range jobs {
				failed := task.IsTaskFailed(t.TaskTimeout)
				if failed {
					task.SetTaskRESULT([]byte("something went wrong"))
				} else {
					task.SetTaskRESULT([]byte("task has been successed"))
				}

				task.SetFT(time.Now().Format(time.RFC3339Nano))
				time.Sleep(time.Millisecond * 150)
				tResult.Add(task, failed)
			}
		}()
	}

	go func() {
		t.wg.Wait()
		tResult.Close()
	}()

	return tResult
}
