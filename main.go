package main

import (
	"desingpatterns/interpreter"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func main() {

	// ITERPRETER
	stack := interpreter.PolishNotationStack2{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, o := range operators {
		if o == interpreter.SUM || o == interpreter.SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := interpreter.OperatorFactory(o, left, right)
			res := interpreter.Value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(o)
			if err != nil {
				panic(err)
			}
			temp := interpreter.Value(val)
			stack.Push(&temp)
		}
	}
	println(int(stack.Pop().Read()))

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
