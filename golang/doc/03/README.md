# Go的变量、常量以及与其他语言的差异

### 1、Go编写测试程序

- 源码文件以_test结尾：hello_test.go
- 测试方法名以Test开头：func TestXXX(t *testing.T) {...}

```go
package test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}
```

## 2、变量的演示(斐波那契数列)

```go
//斐波那契数列：1, 1, 2, 3, 5, 8, 13, 21, 34, 55
func TestFibList(t *testing.T) {
	//第一种
	//var a int = 1
	//var b int = 1

	//第二种
	//var(
	//	a int =1
	//	b =1
	//)

	//第三得
	a:=1
	b:=1

	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
```

## 3、与其他主要编著程语言的差异

- 赋值可以进行自动类型推断

- 在一个赋值语句中可以对多个变量时行同时赋值

```go
//交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	/*
	tmp := a
	a = b
	b = tmp
	*/
	a, b = b, a //一个赋值语句里对多个变量进行赋值
	t.Log(a, b)
}
```

## 4、常量的赋值

```go
package constant_test

import "testing"

//定义星期常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConstantTry1(t *testing.T) {
	t.Log(Monday, Tuesday)
}

//定位二进制数
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry2(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable)
	t.Log(a&Writable == Writable)
	t.Log(a&Executable == Executable)
}
```

