package models

import (
	"encoding/json"
	"io"
	"reflect"
)

// Payload ...
type Payload int

// Process ...
func (p *Payload) Process() ([]byte, error) {
	return nil, nil
}

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// JobQueue A buffered channel that we can send work requests on.
var JobQueue chan Job

// DoJobProcess ...
func DoJobProcess(w io.Writer, fn interface{}) {
	result := reflect.ValueOf(fn).Call([]reflect.Value{})[0].Interface()
	json.NewEncoder(w).Encode(result)
}
