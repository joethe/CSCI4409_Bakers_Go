package main

import "fmt"
import "os"
import "math/rand"
import "time"
import "strconv"

// channels:
// queue -- customers send work requests here. Work requests consist of a request, and a channel to send results back to. Servers read from this channel.
// ackChan -- customers send an acknowledgment when they recieve their results.

type workOrder struct {
  sender chan int
  req int
}


func main() {
  // Arguments specify the number of servers and customers
  args := os.Args[1:]

  servers, err := strconv.Atoi(args[0])
  customers, err1 := strconv.Atoi(args[1])
  if err != nil || err1 != nil {
    fmt.Println(err)
    os.Exit(2)
  }

  queue := make(chan workOrder, customers) // channel, buffered to the amount of customers we're expected to deal with...

  ackChan := make(chan bool, customers) // channel, buffered to the amount of customers we're expecting to deal with...

  // for loop creates servers (functions, each in own goroutine)
  for i := 0; i < servers; i++ {
    go server(queue)
  }

  // for loop creates customers (functions, each in own goroutine)
  for i := 0; i < customers; i++ {
    go customer(queue, ackChan)
  }

  // wait for all customers to be served.
  for i := 0; i < customers; i++ {
    <- ackChan
  }
}

// creates a channel which will contain all successive fib numbers.
// apparently doing it this way is too fast....
func fibGen() (chan int) {
  c := make(chan int)

  go func() {
    for i, j := 0, 1; ; i, j = j+i, i {
      c <- i
    }
  }()

  return c
}

// this one seems slightly slower....
func slowFib(n int)(int){
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return slowFib(n - 2) + slowFib(n - 1)
  }
}


func server(c chan workOrder) {

  //fmt.Println("New Server Created!")

  // look for work on the channel c, take work, do work, send completed work back, repeat.
  for {
    wo := <-c

    //fibChan := fibGen()

    //var fib int

    //for i := 0; i <= wo.req; i++ {
    //  fib = <- fibChan
    //}

    // also too fast...
    //  for k, i, j := 0, 0, 1; k <= wo.req ; i, j = j+i, i {
    //    k++
    //    fib = i
    //  }

    wo.sender <- slowFib(wo.req)
  }
}

func customer(c chan workOrder, ack chan bool) {

  //fmt.Println("New Customer Created!")

  // sleep for a random period of time
  time.Sleep(time.Duration(rand.Intn(15000)) * time.Millisecond)

  //fmt.Println("Customer awake!")

  resultChan := make(chan int)

  // generate work request
  wo := workOrder{sender: resultChan, req: rand.Intn(40)}
  //wo := workOrder{sender: resultChan, req: 20}

  // send work request to channel c
  c <- wo

  // wait for reply on own channel
  result := <- resultChan
  // print work result
  fmt.Println("Got fib: ", result)
  ack <- true
  // die
  return
}






//
