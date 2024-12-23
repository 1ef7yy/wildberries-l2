package main

type CPU struct{}

func (c *CPU) Freeze() {
	println("CPU is frozen")
}

func (c *CPU) Jump(pos int) {
	println("CPU jumped to", pos)
}

func (c *CPU) Execute() {
	println("CPU executed")
}

type HardDrive struct{}

func (h *HardDrive) Read(lba, size int) []byte {
	println("Read", size, "bytes from LBA", lba)
	return make([]byte, size)
}

type Memory struct{}

func (m *Memory) Load(pos int, data []byte) {
	println("Loaded", len(data), "bytes from", pos)
}

type ComputerFacade struct {
	CPU       *CPU
	HardDrive *HardDrive
	Memory    *Memory
}

func (c *ComputerFacade) Start() {
	c.CPU.Freeze()
	c.Memory.Load(0, c.HardDrive.Read(0, 10e10))
	c.CPU.Jump(0)
	c.CPU.Execute()
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		CPU:       &CPU{},
		HardDrive: &HardDrive{},
		Memory:    &Memory{},
	}
}
