package serviceworker

import (
	"sync"
)

type Worker struct {
	TotalGoroutines int32
	Working int32
	wg sync.WaitGroup
}



func (this *Worker)Run() {
	
}