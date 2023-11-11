package stack

import "sync"

func NewStackService(mut *sync.Mutex) StackService {
	return StackService{
		CPU:    NewCPUStack(),
		Memory: NewMemoryStack(),
	}
}

type StackService struct {
	CPU    Stack
	Memory Stack
}

type Stack interface {
	GetUsage() float32
	IncreaseResourcesUsage(qty float32)
}
