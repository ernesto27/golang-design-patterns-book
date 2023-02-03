package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(s.LogWriter, "--------------------------------\n")
	s.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

type MyList []int

func (m MyList) Len() int {
	return len(m)
}

func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyList) Less(i, j int) bool {
	return m[i] < m[j]
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func reciever(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case mgs := <-goodbyeCh:
			println(mgs)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

var wait sync.WaitGroup

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		time.Sleep(time.Second)

		ch <- 2

		close(ch)
	}()

	for v := range ch {
		println(v)
	}

	// helloCh := make(chan string, 1)
	// goodbyeCh := make(chan string, 1)
	// quitCh := make(chan bool)

	// go reciever(helloCh, goodbyeCh, quitCh)

	// go sendString(helloCh, "hello!")
	// time.Sleep(time.Second)
	// go sendString(goodbyeCh, "goodbye!")
	// <-quitCh

	// channel := make(chan string, 1)

	// go func() {
	// 	channel <- "Hello world"
	// 	println("Finish goroutiners")
	// }()

	// time.Sleep(time.Second)

	// message := <-channel
	// println(message)

	// toUpperSync("Hello callbacks", func(s string) {
	// 	fmt.Println(s)
	// })

	// wait.Add(1)
	// toUpperAsync("hello callback", func(s string) {
	// 	fmt.Println(s)
	// 	wait.Done()
	// })

	// fmt.Println("HI")
	// wait.Wait()

	// GOROUTINES
	// var wait sync.WaitGroup

	// goRoutines := 5
	// wait.Add(goRoutines)

	// for i := 0; i < goRoutines; i++ {
	// 	go func(goRoutineID int) {
	// 		fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
	// 		wait.Done()
	// 	}(i)
	// }
	// wait.Wait()

	// VISITOR
	// start := state.StartState{}
	// game := state.GameContext{
	// 	Next: &start,
	// }

	// for game.Next.ExecuteState(&game) {
	// }

	// // VISITOR
	// products := make([]visitor.Visitable2, 3)
	// products[0] = &visitor.Rice{
	// 	Product: visitor.Product{
	// 		Price: 32.0,
	// 		Name:  "Some rice",
	// 	},
	// }
	// products[1] = &visitor.Pasta{
	// 	Product: visitor.Product{
	// 		Price: 40.0,
	// 		Name:  "Some pasta",
	// 	},
	// }

	// products[2] = &visitor.Fridge{
	// 	Product: visitor.Product{
	// 		Price: 50,
	// 		Name:  "A fridge",
	// 	},
	// }

	// priceVisitor := &visitor.PriceVisitor{}

	// for _, p := range products {
	// 	p.Accept(priceVisitor)
	// }

	// fmt.Printf("Total: %f\n", priceVisitor.Sum)

	// nameVisitors := &visitor.NamePrinter{}

	// for _, p := range products {
	// 	p.Accept(nameVisitors)
	// }

	// fmt.Printf("\nProduct list:\n-------------\n%s", nameVisitors.Names)

	// ITERPRETER
	// stack := interpreter.PolishNotationStack2{}
	// operators := strings.Split("3 4 sum 2 sub", " ")

	// for _, o := range operators {
	// 	if o == interpreter.SUM || o == interpreter.SUB {
	// 		right := stack.Pop()
	// 		left := stack.Pop()
	// 		mathFunc := interpreter.OperatorFactory(o, left, right)
	// 		res := interpreter.Value(mathFunc.Read())
	// 		stack.Push(&res)
	// 	} else {
	// 		val, err := strconv.Atoi(o)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		temp := interpreter.Value(val)
	// 		stack.Push(&temp)
	// 	}
	// }
	// println(int(stack.Pop().Read()))

	// var myList MyList = []int{6, 4, 2, 8, 1}
	// fmt.Println(myList)
	// sort.Sort(myList)
	// fmt.Println(myList)

	// // COMMAND
	// queue := command.CommandQueue{}
	// queue.AddCommand(command.CreateCommand("First message"))
	// queue.AddCommand(command.CreateCommand("Second message"))
	// queue.AddCommand(command.CreateCommand("Third message"))
	// queue.AddCommand(command.CreateCommand("Fourth message"))
	// queue.AddCommand(command.CreateCommand("Fifth message"))

	// var timeCommand command.Command2
	// timeCommand = &command.TimePassed{Start: time.Now()}

	// var helloCommand command.Command2
	// helloCommand = &command.HelloMessage{}

	// time.Sleep(time.Second)
	// fmt.Println(timeCommand.Info())
	// fmt.Println(helloCommand.Info())

	// http.Handle("/", &MyServer{})
	// http.Handle("/", &LoggerServer{
	// 	LogWriter: os.Stdout,
	// 	Handler:   &MyServer{},
	// })

	// http.Handle("/", &LoggerServer{
	// 	Handler: &BasicAuthMiddleware{
	// 		Handler:  new(MyServer),
	// 		User:     "1111",
	// 		Password: "1111",
	// 	},
	// 	LogWriter: os.Stdout,
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// f, _ := os.Create("name.txt")
	// c := adapter.Counter{f}
	// c.Count(10)

	// pipeReader, pipeWriter := io.Pipe()
	// defer pipeReader.Close()
	// defer pipeWriter.Close()

	// counter := adapter.Counter{
	// 	Writer: pipeWriter,
	// }

	// s := composition.CompositeSwimmerA{
	// 	MySwim: composition.Swim,
	// }
	// s.MyAthlete.Train()
	// s.MySwim()

	// f := composition.Shark{
	// 	Swim: composition.Swim,
	// }
	// f.Eat()
	// f.Swim()

	// swimmer := composition.CompositeSwimmerB{
	// 	Trainer: &composition.Athlete{},
	// 	Swimmer: &composition.SwimmerImpl{},
	// }

	// swimmer.Train()
	// swimmer.Swim()
}
