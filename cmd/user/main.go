package main

import (
	"fmt"
	"go-shop-microservice/internal/pkg/event"
)

type RegisterEvent struct {
}

type SendListener struct {
}

func (listener *SendListener) Process(event event.Event) {
	fmt.Println("执行send" + event.Name())
}

func (listener *SendListener) IsAsync() bool {
	return true
}

func (e *RegisterEvent) Name() string {
	return "register"
}

func main() {
	sendListener := &SendListener{}
	registerEvent := &RegisterEvent{}
	em := event.Manager{EventListenerMap: make(map[string][]event.Listener)}
	em.Add("register", sendListener)
	em.Dispatch(registerEvent)
}
