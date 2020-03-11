package pool

// WorkerChannel :
var WorkerChannel = make(chan chan Work)

// Collector :
type Collector struct {
	Work chan Work
	End  chan bool
}

// StartDispatcher :
func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker

	input := make(chan Work)
	end := make(chan bool)
	collector := Collector{Work: input, End: end}

	// log.Printf("Starting Worker : %v", workerCount)
	for i < workerCount {
		i++
		worker := Worker{
			ID:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker)
	}
	// log.Print("\n\n")

	go func() {
		for {
			select {
			case work := <-input:
				worker := <-WorkerChannel
				worker <- work
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			}

		}
	}()

	return collector
}
