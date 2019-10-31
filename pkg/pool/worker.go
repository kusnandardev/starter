package pool

import (
	"log"

	"kusnandartoni/starter/pkg/work"
)

// Work :
type Work struct {
	ID  int
	Job string
}

// Worker :
type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

// Start :
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				work.DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

// Stop :
func (w *Worker) Stop() {
	log.Printf("Worker [%d] is stopping", w.ID)
	w.End <- true
}
