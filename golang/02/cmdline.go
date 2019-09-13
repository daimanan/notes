package main

import (
	"os"
	"fmt"
)

//传入命令行参数

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("命令行传入参数：", os.Args[1])
	}

}
