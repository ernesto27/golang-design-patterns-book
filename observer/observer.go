package observer

import "fmt"

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range p.ObserverList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	p.ObserverList = append(p.ObserverList[:indexToRemove], p.ObserverList[indexToRemove+1:]...)
}

func (p *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, o := range p.ObserverList {
		o.Notify(m)
	}
}

type TestObserver struct {
	ID      int
	Message string
}

func (t *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received \n", t.ID, m)
	t.Message = m
}
