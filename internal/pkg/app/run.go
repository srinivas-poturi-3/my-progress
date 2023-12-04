package app

import (
	"strings"
	"sync"
	"templates/internal/server"
)

const templatesName = "templates"

func Make(wg *sync.WaitGroup) App {
	t := new(Worker)
	t.name = templatesName
	t.wg = wg
	return t
}

func (w *Worker) Name() string {
	if len(strings.TrimSpace(w.name)) == 0 {
		return ""
	}

	return w.name
}

func (w *Worker) Start() error {
	w.wg.Add(1)
	srv := server.Server{}
	srv.New()
	srv.Start()
	return nil
}

func (w *Worker) Stop() {
	close(w.tasks)
	w.wg.Done()

	return
}
