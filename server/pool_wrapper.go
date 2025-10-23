package server

import (
	"context"
	"time"

	"github.com/alitto/pond"
)

type PondPoolWrapper struct {
	pool *pond.WorkerPool
}

func (p PondPoolWrapper) Submit(ctx context.Context, task func()) {
	p.pool.Submit(task)
}

func (p PondPoolWrapper) StopAndWaitFor(deadline time.Duration) {
	p.pool.StopAndWaitFor(deadline)
}

func (p PondPoolWrapper) Stop() context.Context {
	return p.pool.Stop()
}

func (p PondPoolWrapper) StopAndWait() {
	p.pool.StopAndWait()
}
