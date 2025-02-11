package hooks

type Subscriber interface {
	Handle(event Event)
}
