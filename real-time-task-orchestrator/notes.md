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

We implement a naive observer pattern:

1) Store will maintain a list of active channels
2) Store broadcasts new tasks to them
