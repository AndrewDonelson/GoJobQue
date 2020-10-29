package models

import (
	"fmt"
	"log"
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker ...
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				log.Println("Starting Job")
				// we have received a work request.
				result, err := job.Payload.Process()
				if err != nil {
					log.Println(err)
				} else {
					fmt.Println(result)
				}
				// if DoJobProcess()
				// if err := job.Payload.UploadToS3(); err != nil {
				// 	log.Errorf("Error uploading to S3: %s", err.Error())
				// }

			case <-w.quit:
				// we have received a signal to stop
				log.Println("Quitting Job")
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
