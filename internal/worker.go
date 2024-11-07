package internal

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/splashk1e/jet/internal/services"
)

type Worker struct {
	service *services.WorkerService
	mu      *sync.Mutex
}

func NewWorker(service *services.WorkerService) *Worker {
	return &Worker{
		service: service,
		mu:      &sync.Mutex{},
	}
}
func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Minute)
	wg := sync.WaitGroup{}
Loop:
	for {
		select {
		case <-ticker.C:
			wg.Wait()
			wg.Add(1)
			go func() {
				logrus.Info("worker starts file update")
				defer wg.Done()
				defer logrus.Info("worker ends file update")

				w.mu.Lock()
				defer w.mu.Unlock()
				if err := w.service.FileUpdate(); err != nil {
					logrus.Errorf("can not file update with error:%s", err.Error())
				}

			}()
		case <-ctx.Done():
			wg.Wait()
			break Loop
		}

	}
}
