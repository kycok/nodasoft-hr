package task

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"nodasoft-hr/internal/model"
)

type Producer struct {
	wg           sync.WaitGroup
	WorkersCount int
}

func (t *Producer) Run(ctx context.Context) chan model.Ttype {
	jobs := make(chan model.Ttype, 10)

	var id atomic.Int64
	id.Store(1)

	t.wg.Add(t.WorkersCount)

	for i := 0; i < t.WorkersCount; i++ {

		go func() {
			defer t.wg.Done()

			for loop := true; loop; {
				select {
				case <-ctx.Done():
					loop = false
				default:
					timeNow := time.Now()
					formatTime := timeNow.Format(time.RFC3339Nano)

					if timeNow.Nanosecond()%2 > 0 { // вот такое условие появления ошибочных тасков
						// Зависит от хоста, точности до наносекунд может не быть
						// if timeNow.UnixMicro()%2 > 0 { // вот такое условие появления ошибочных тасков
						formatTime = "Some error occured"
					}

					jobs <- model.NewTtype(int(id.Load()), formatTime) // передаем таск на выполнение
					id.Add(1)
				}
			}
		}()
	}

	go func() {
		t.wg.Wait()
		close(jobs)
	}()

	return jobs
}
