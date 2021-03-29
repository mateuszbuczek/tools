package pool

import (
	"sync"
)

type goRoutinePool struct {
	coordinatingChan      chan struct{}
	confirmationWaitGroup *sync.WaitGroup
}

func New(size int, group *sync.WaitGroup) *goRoutinePool {
	return &goRoutinePool{
		coordinatingChan:      make(chan struct{}, size),
		confirmationWaitGroup: group,
	}
}

func (p *goRoutinePool) Execute(fun func(a ...interface{}), args ...interface{}) {
	go func() {
		defer p.confirmationWaitGroup.Done()
		p.coordinatingChan <- struct{}{}
		fun(args...)
		<-p.coordinatingChan
	}()
}
