package main

import "fmt"

// Команда
type Command interface {
	Execute()
}

// Приемник команды
type Light struct{}

func (l *Light) On() {
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	fmt.Println("Свет выключен")
}

func (l *Light) Dim(level int) {
	fmt.Printf("Яркость света установлена на %d%%\n", level)
}

type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.On()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.Off()
}

type DimCommand struct {
	light *Light
	level int
}

func (c *DimCommand) Execute() {
	c.light.Dim(c.level)
}

// Инвокер команды
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// func main() {
// 	light := &Light{}
// 	remote := &RemoteControl{}

// 	turnOnCmd := &TurnOnCommand{light: light}
// 	turnOffCmd := &TurnOffCommand{light: light}
// 	dimCmd := &DimCommand{light: light, level: 50}

// 	remote.SetCommand(turnOnCmd)
// 	remote.PressButton()

// 	remote.SetCommand(turnOffCmd)
// 	remote.PressButton()

// 	remote.SetCommand(dimCmd)
// 	remote.PressButton()
// }
