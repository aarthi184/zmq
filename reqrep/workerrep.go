package main

import (
    zmq "github.com/pebbe/zmq4"
    "fmt"
    "time"
    "strconv"
)

func main() {
    responder, _ := zmq.NewSocket(zmq.REP)
    defer responder.Close()
    responder.Connect("tcp://localhost:5560")

    for {
        request, _ := responder.Recv(0)
        fmt.Printf("Received request: [%s]\n", request)
        time.Sleep(time.Second)
        i, _ := strconv.Atoi(request)
        resp := doWork(i)
        responder.Send(resp, 0)
    }
}

func doWork(count int) string {
    switch count {
    case 1:
        return "One"
    case 2:
        return "Two"
    case 3:
        return "Three"
    case 4:
        return "Four"
    case 5:
        return "Five"
    }
    return "Invalid"
}
