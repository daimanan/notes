package main //包名和文件可以不一样

import (
	"fmt"
	"os"
)

//程序入口
func main()  {
	fmt.Println("Hello,world")
	os.Exit(-1)
	fmt.Println("我可以用go开发程序了！^_^")
}
