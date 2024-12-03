package main

import (
	"trojan/executer"
)

func main(){
	res,err := executer.DoCommand("ls", []string{}, 5)
}
