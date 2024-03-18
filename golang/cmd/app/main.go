package main

import (
	"context"
	"time"

	"nodasoft-hr/internal/task"
)

// ЗАДАНИЕ:
// * сделать из плохого кода хороший;
// * важно сохранить логику появления ошибочных тасков;
// * сделать правильную мультипоточность обработки заданий.
// Обновленный код отправить через merge-request.

// приложение эмулирует получение и обработку тасков, пытается и получать и обрабатывать в многопоточном режиме
// В конце должно выводить успешные таски и ошибки выполнены остальных тасков

const (
	createTimeoutSec = 3
	taskTimeoutSec   = 20
	workersCount     = 10
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(createTimeoutSec))
	defer cancel()

	producer := task.Producer{WorkersCount: workersCount}
	// запускаем производителя тасков, который возвращает канал с тасками для обработки,
	// вызов не блокирующий, можно сразу же начинать обработку, завершится по таймауту контекста и закроет каналы
	jobs := producer.Run(ctx)

	processor := task.Processor{WorkersCount: workersCount, TaskTimeout: taskTimeoutSec}
	// запускаем обработчик тасков, который возвращает канал с результатами обработки,
	// вызов не блокирующий, можно сразу же начинать перекладывать из каналов в нужную структуру
	results := processor.Run(jobs)

	//перекладываем, блокирующий
	results.Collect()
	results.Print()
}
