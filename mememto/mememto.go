package mememto

import (
	"fmt"
)

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found")
	}
	return c.mementoList[i], nil
}

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento2 struct {
	memento Command
}

type originator2 struct {
	Command Command
}

func (o *originator2) NewMemento2() Memento2 {
	return Memento2{memento: o.Command}
}

func (o *originator2) ExtractAAndStoreCommand(m Memento2) {
	o.Command = m.memento
}

type careTaker2 struct {
	mementoList []Memento2
}

func (c *careTaker2) Add(m Memento2) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker2) Pop() Memento2 {
	if len(c.mementoList) > 0 {
		temp := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0 : len(c.mementoList)-1]
		return temp
	}
	return Memento2{}
}
