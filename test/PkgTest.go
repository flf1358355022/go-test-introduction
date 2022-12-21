package main

import (
	"MyTest/packageUtils"
	"fmt"
)

func main() {
	resultNum := packageUtils.NumberUtils(100, 200, "*")
	fmt.Printf("执行包调用函数结果:%v", resultNum)
}
