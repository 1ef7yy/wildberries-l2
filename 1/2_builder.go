package main

type Computer struct {
	CPU       *CPU
	HardDrive *HardDrive
	Memory    *Memory
	HasWifi   bool
	HasUSB    bool
}

type ComputerBuilder interface {
	InstallCPU(*CPU) ComputerBuilder
	InstallHardDrive(*HardDrive) ComputerBuilder
	InstallMemory(*Memory) ComputerBuilder
	InstallWifi() ComputerBuilder
	InstallUSB() ComputerBuilder
	Build() *Computer
}

func NewComputerBuilder() ComputerBuilder {
	return &computerBuilder{
		Computer: &Computer{},
	}
}

type computerBuilder struct {
	Computer *Computer
}

func (b *computerBuilder) InstallCPU(cpu *CPU) ComputerBuilder {
	b.Computer.CPU = cpu
	return b
}

func (b *computerBuilder) InstallHardDrive(hd *HardDrive) ComputerBuilder {
	b.Computer.HardDrive = hd
	return b
}

func (b *computerBuilder) InstallMemory(mem *Memory) ComputerBuilder {
	b.Computer.Memory = mem
	return b
}

func (b *computerBuilder) InstallWifi() ComputerBuilder {
	b.Computer.HasWifi = true
	return b
}

func (b *computerBuilder) InstallUSB() ComputerBuilder {
	b.Computer.HasUSB = true
	return b
}

func (b *computerBuilder) Build() *Computer {
	return b.Computer
}

type Director struct {
	builder ComputerBuilder
}

func (d *Director) ConstructComputer(cpu *CPU, hd *HardDrive, mem *Memory, hasWifi, hasUSB bool) *Computer {
	d.builder.InstallCPU(cpu).
		InstallHardDrive(hd).
		InstallMemory(mem).
		InstallUSB().
		InstallWifi()
	return d.builder.Build()
}
