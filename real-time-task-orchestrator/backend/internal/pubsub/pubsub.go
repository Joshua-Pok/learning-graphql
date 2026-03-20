package pubsub

type Pubsub struct {
}

func (p *Pubsub) Publish(topic string, data interface{}) {

}

func (p *Pubsub) Subscribe(topic string) <-chan interface{} {

}
