package stack

type Memory struct {
	Usage float32
}

func NewMemoryStack() Stack {
	return &Memory{0.0}
}

func (memory *Memory) GetUsage() float32 {
	return memory.Usage
}

func (memory *Memory) IncreaseResourcesUsage(qty float32) {
	memory.Usage += qty
}
