package main

import (
	"desingpatterns/command"
	"fmt"
	"io"
	"net/http"
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

func main() {

	// // COMMAND
	// queue := command.CommandQueue{}
	// queue.AddCommand(command.CreateCommand("First message"))
	// queue.AddCommand(command.CreateCommand("Second message"))
	// queue.AddCommand(command.CreateCommand("Third message"))
	// queue.AddCommand(command.CreateCommand("Fourth message"))
	// queue.AddCommand(command.CreateCommand("Fifth message"))

	var timeCommand command.Command2
	timeCommand = &command.TimePassed{Start: time.Now()}

	var helloCommand command.Command2
	helloCommand = &command.HelloMessage{}

	time.Sleep(time.Second)
	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())

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
