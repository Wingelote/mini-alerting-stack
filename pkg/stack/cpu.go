package stack

type CPU struct {
	Usage float32
}

func NewCPUStack() Stack {
	return &CPU{0.0}
}

func (cpu *CPU) GetUsage() float32 {
	return cpu.Usage
}

func (cpu *CPU) IncreaseResourcesUsage(qty float32) {
	cpu.Usage += qty
}
