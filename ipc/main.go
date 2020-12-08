package main

import (
	"fmt"
	pipe2 "golang-demo/ipc/pipe"
	"io"
	"strconv"
	"time"
)

func main()  {
	// 基本管道通信
	//pipe()

	for i := 13; i<300 ;i ++ {
		fmt.Printf("https://www.bilibili.com/video/BV1HE411L7bz?p=" + strconv.Itoa(i) + "分隔")
	}
}

func pipe()  {
	reader,writer := io.Pipe()
	go pipe2.Read(reader)
	go pipe2.Write(writer)
	time.Sleep(time.Second * 10)
}
