package event

type Listener interface {
	Process(event Event)
	IsAsync() bool
}
