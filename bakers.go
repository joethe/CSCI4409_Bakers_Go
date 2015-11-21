package main

import "os"
import "fmt"

// channels:
// queue -- customers send work requests here. Work requests consist of a request, and a channel to send results back to. Servers read from this channel.

type workOrder struct {
  sender chan
  request int
}

queue := make(chan workOrder)

func main() {
  args := os.Args[1:]

  // setup

}


func server(c chan) {

  // look for work on the channel c, take work, do work, send completed work back, repeat.

}

func customer(c chan) {
  // sleep for a random period of time

  // generate work request

  // send work request to channel c

  // wait for reply on own channel
  // close own channel
  // print work result
  // die
}

// Arguments specify the number of servers and customers

// for loop creates servers (functions, each in own goroutine)

// for loop creates customers (functions, each in own goroutine)

//
