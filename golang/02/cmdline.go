package main

import (
	"os"
	"fmt"
)

//传入命令行参数 单行注释

/*
这是多行注释
注释有两种：
	单行注释
	多行注释
 */

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("命令行传入参数：", os.Args[1])
	}
}
