# My notes on workshop

mechanics vs semantics:
m - how things are implemented
s - the behavior

garbage collection article:
https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html

## GC phases

STW - stop the word

1. mark start(STW)
2. marking (concurrent) traverse the heap and analyze if values is still in use (there is a pointer on stack that points to it)
3. mark termination (STW)

mark assist

GOGC variable

amount fo allocspace

## concurrency

if work can be done out of order it's concurrent

3GHz(3 clock cycles/ns) * 4 instructions per cycle = 12 instructions per ns!

1 ns ............. 1 ns .............. 12 instructions  (one)
1 µs .......... 1000 ns .......... 12,000 instructions  (thousand)
1 ms ..... 1,000,000 ns ...... 12,000,000 instructions  (million)
1 s .. 1,000,000,000 ns .. 12,000,000,000 instructions  (billion)

context switch 12K instructions (1 micro second)

CPU bound workload requires parallelism – sum an array of 1KK items. How many hw threads we have? let's use all of them! It'll newer go to a waiting state
IO bound workload not requires parallelism - fetch 1000 urls. How many threads should we use? It depends

coroutines - application level threads
go routines don't have priorities

global run queue
local run queue

every time threads talk to each other there are two context switches ctx/ready and ctx/waiting

go context switch - 200 ns instead of 1 micro second

go runtime transforms IO bounded workload into CPU workload for operation system

synchronization - putting events in a line
orchestration - two or more need to talk to each other. Use channel for this

Using channel for sync you doing it wrong
WaitGroup if for orchestration

## Notes

Always seek for a guarantee that concurrent code works correctly and it's not a coincident

EVERY GOROUTINE IN A RUNNABLE STATE IS RUNNING!

good practice for readability

```
mu.Lock()
{
    //do stuff
}
mu.Unlock()
```

if you use mutex as a struct attribute, you should use reference semantics

if you use RWLock() never copy-paste .Lock() and Unlock()

Channels are for signaling it's not a queue, don't treat they as a data structure

if a goroutine sends a signal does it need a guarantee that signal is received

the cost of the guarantee is latency

## channels

buffered channels gives no guarantee

fanOut pattern it's hard to scale. It's okay in cli and cron. but be aware to use it on a server

pooling needs guarantee use unbuffered channel. Otherwise, you can't handle cancelation

channels is just signaling semantics, not a queue. because concurrency assumes out of order operations


# cancellation
use buffered channel for cancellation

if you use context in a business layer, be sure it will work with empty context.
don't rely that some value will be in the context


NOTE:
for io bound jobs number of threads matters less


## mem profile

flat column shows - that allocated value in current function
cum colum shows - that value was created outside current function and returned here

inlined functions not cause allocations

