package main

import (
	pipe2 "golang-demo/ipc/pipe"
	"io"
	"time"
)

func main()  {
	// 基本管道通信
	pipe()
}

func pipe()  {
	reader,writer := io.Pipe()
	go pipe2.Read(reader)
	go pipe2.Write(writer)
	time.Sleep(time.Second * 10)
}
