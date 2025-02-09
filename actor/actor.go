package actor

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type helloer struct {
	owner string
}

func newHelloer(owner string) actor.Producer {
	return func() actor.Receiver {
		return &helloer{owner: owner}
	}
}

func RunActor() {
	engine, _ := actor.NewEngine(actor.NewEngineConfig())

	pid := engine.Spawn(newHelloer("gui"), "hello")

	for i := 0; i < 3; i++ {
		engine.Send(pid, &message{data: "hello, world!"})
	}

	engine.Poison(pid).Wait()
}

type message struct {
	data string
}

func (h *helloer) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		fmt.Println("helloer has initialized")
	case actor.Started:
		fmt.Println("helloer has started")
	case actor.Stopped:
		fmt.Println("helloer has stopped")
	case *message:
		fmt.Println("hello world", msg.data)
	}
}
