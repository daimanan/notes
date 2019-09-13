# 编写第一个Go程序，Hello World

## 1.程序的结构

```go
package main //包，代码所在的模块

import "fmt" //引入代码依赖

//程序入口
func main() {
	fmt.Println("Hello World")
}
```
> “Hello, world"程序是指在计算机屏幕上输出“Hello , world”这行字符串的计算机程序，“hello, world”的中文意思是“你好，世界。”。这个例程在Brian Kernighan(布莱恩·柯林汉) 和Dennis M. Ritchie(丹尼斯·里奇)合著的《The C Programme Language》使用而广泛流行。因为它的简洁，实用，并包含了一个该版本的C程序首次在1974年Brian Kernighan所撰写的《Programming in C: A Tutorial》出现
>
> ```c
> printf("hello, world\n");
> ```
>
> 实际上将“Hello”和“World”一起使用的程序最早出现于1972年，在贝尔实验室成员Brian Kernighan撰写的内部技术文件《Introduction to the Language B》之中：
>
> ```c
> main(){
>     extern a,b,c;
>     putchar(a);putchar(b);putchar(c);putchar('!*n');
> }
> a'hell';
> b'o,w';
> c'orld';
> ```
>
> 最初的"hello, world"打印内容有个标准，即全小写，有逗号，逗号后空一格，且无感叹号。不过沿用至今，完全遵循传统标准形式的反而很少出现。

## 2.应用程序入口(开发注意问题)

- 必须是main包：package main 
- 必须是main方法：func main()
- 文件名可以不是main.go

## 3.Go的命令行

- go run 源码文件.go		运行源码文件
- go build 源码文件.go	编译生成可执行文件

## 4.退出返回值(与其他的语言有差异)

- Go中main函数不支持任何返回值
- 通过os.Exit来返回状态 （0 正常、-1 异常退出）

## 5.获取命令行参数(与其他的语言有差异)

- main函数不支持传入参数

  ```go
  func main(arg []string) //是错误的
  ```

- 在程序中直接通过os.Args获取命令行参数

  ```go
  package main
  
  import (
  	"fmt"
  	"os"
  )
  
  func main()  {
  	fmt.Println("--------args-------")
  	if len(os.Args) > 1{
  		fmt.Println(os.Args[1])
  	}
  }
  ```