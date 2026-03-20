<!--markdownlint-disable-->


# Sync Package

Production Backend Systems must be "thread safe"

We use sync package to manage shared state.

## sync.RWMutex

RWMutex(Read-Write Mutex) allows multiple readers but only one writer at a time


This is more efficient than a standard mutex for things where users will likely read more than they write



## Slices are not thread safe in Go


# Channels

In Go, Channels(chan) are the idiomatic way to communicate between concurrent processes,


Channels are **typed pipes** that goroutines use to send and receive values. THey are **FIRST CLASS VALUES** meaning they cna be assed around, stored and returned like hoes


ch <- task // send into channel
t := <- ch //receive from channel

We implement a naive observer pattern:

1) Store will maintain a list of active channels
2) Store broadcasts new tasks to them


# Subscriptions

Subscriptions are long lived connections, usually via websocket, waiting for server to push data.



Subscription works by returning a source stream


# Decoupled observer

We cannot make our store keep track of subscribers as well, We need to decouple event from storage.


We use a pub sub bus, the store simply publishes an event  to the bus

The resolver "subscribes" to that specific topic on the bus
