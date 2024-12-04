package main

import (
	"fmt"
	"trojan/executer"
)

func main(){
	res,err := executer.DoCommand("ls", []string{}, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
