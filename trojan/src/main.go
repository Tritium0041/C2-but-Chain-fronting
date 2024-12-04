package main

import (
	"executer"
)

func main(){
	res,err := executer.DoCommand("ls", []string{}, 5)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
