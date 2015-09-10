package main

import (
    zmq "github.com/pebbe/zmq4"
    "fmt"
)

func main() {
    subscriber, _ := zmq.NewSocket(zmq.SUB)
    defer subscriber.Close()

    subscriber.Connect("tcp://localhost:1123")
    subscriber.SetSubscribe("AB")

    for {
        address, _ := subscriber.Recv(0)
        contents, _ := subscriber.Recv(0)
        fmt.Printf("[%s] %s\n", address, contents)
    }
}
