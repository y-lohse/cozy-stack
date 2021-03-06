package jobs

import (
	"fmt"

	"github.com/cozy/cozy-stack/pkg/config"
)

// WorkersList is a map associating a worker type with its acutal
// configuration.
type WorkersList []*WorkerConfig

// WorkersList is the list of available workers with their associated Do
// function.
var workersList WorkersList

// GetWorkersList returns a globally defined worker config list.
func GetWorkersList() ([]*WorkerConfig, error) {
	if config.GetConfig().Jobs.NoWorkers {
		return nil, nil
	}

	workersConfs := config.GetConfig().Jobs.Workers
	workers := make(WorkersList, 0, len(workersList))
	for _, w := range workersList {
		for _, c := range workersConfs {
			if c.WorkerType == w.WorkerType {
				w = applyWorkerConfig(w, c)
			}
		}
		workers = append(workers, w)
	}

	for _, c := range workersConfs {
		_, found := findWorkerByType(c.WorkerType)
		if !found {
			return nil, fmt.Errorf("Defined configuration for the worker %q that does not exist",
				c.WorkerType)
		}
	}

	return workers, nil
}

func applyWorkerConfig(w *WorkerConfig, c config.Worker) *WorkerConfig {
	w = w.Clone()
	if c.Concurrency != nil {
		w.Concurrency = *c.Concurrency
	}
	if c.MaxExecCount != nil {
		w.MaxExecCount = *c.MaxExecCount
	}
	if c.Timeout != nil {
		w.Timeout = *c.Timeout
	}
	return w
}

func findWorkerByType(workerType string) (*WorkerConfig, bool) {
	for _, w := range workersList {
		if w.WorkerType == workerType {
			return w, true
		}
	}
	return nil, false
}

// AddWorker adds a new worker to global list of available workers.
func AddWorker(conf *WorkerConfig) {
	if conf.WorkerType == "" {
		panic("Missing worker type field")
	}
	for _, w := range workersList {
		if w.WorkerType == conf.WorkerType {
			panic(fmt.Errorf("A worker with of type %q is already defined", conf.WorkerType))
		}
	}
	workersList = append(workersList, conf)
}
