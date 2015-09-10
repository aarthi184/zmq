package main

import (
    zmq "github.com/pebbe/zmq4"
    "fmt"
    "strconv"
)

func main() {
    requester, _ := zmq.NewSocket(zmq.REQ)
    defer requester.Close()
    requester.Connect("tcp://localhost:5559")

    for request := 1; request <= 5; request++ {
        count := strconv.Itoa(request)
        requester.Send(count, 0)
        reply, _ := requester.Recv(0)
        fmt.Printf("Received reply %d [%s]\n", request, reply)
    }
}
