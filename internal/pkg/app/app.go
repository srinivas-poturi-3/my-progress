package app

import "sync"

// Worker represents a worker in the pool.
type Worker struct {
	name  string
	tasks chan func()
	wg    *sync.WaitGroup
}

type App interface {
	Name() string
	Start() error
	Stop()
}
