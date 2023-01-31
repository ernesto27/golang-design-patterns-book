package command

import (
	"fmt"
	"time"
)

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}

type Command2 interface {
	Info() string
}

type TimePassed struct {
	Start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.Start).String()
}

type HelloMessage struct{}

func (h *HelloMessage) Info() string {
	return "Hello world!"
}
