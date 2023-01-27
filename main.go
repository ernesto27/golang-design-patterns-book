package main

import "desingpatterns/composition"

func main() {

	s := composition.CompositeSwimmerA{
		MySwim: composition.Swim,
	}
	s.MyAthlete.Train()
	s.MySwim()

	f := composition.Shark{
		Swim: composition.Swim,
	}
	f.Eat()
	f.Swim()

	swimmer := composition.CompositeSwimmerB{
		Trainer: &composition.Athlete{},
		Swimmer: &composition.SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}
