package main

import (
    zmq "github.com/pebbe/zmq4"
    "fmt"
    "time"
)


func main() {
    publisher, _ := zmq.NewSocket(zmq.PUB)
    defer publisher.Close()
    time.Sleep(time.Second)
    publisher.Bind("tcp://*:1123")

    for {
        var addr, content string
        fmt.Println("Enter stream to publish in:")
        fmt.Scan(&addr)
        fmt.Println("Enter the content to publish:")
        fmt.Scanln(&content)
        publisher.Send(addr, zmq.SNDMORE)
        publisher.Send(content, 0)
        //publisher.Send("C", zmq.SNDMORE)
        //publisher.Send("Publishing C", 0)
    }
}
