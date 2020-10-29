package models

import "log"

// QueManager ...
type QueManager struct {
	dispatcher *Dispatcher
}

// NewQueManager ...
func NewQueManager(maxWorkers int) (*QueManager, error) {
	qm := &QueManager{}
	qm.dispatcher = NewDispatcher(maxWorkers)
	qm.dispatcher.Run()

	log.Println(len(qm.dispatcher.WorkerPool), " workers started")
	return qm, nil
}
