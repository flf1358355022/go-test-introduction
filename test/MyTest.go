package main

import (
	"fmt"
	"strconv"
)

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d = 1, 2
var e, f = 123, "hello"

func main() {
	//g, h := 123, "hello"
	//print("----------")
	//println(x, y, a, b, c, d, e, f, g, h)

	//const (
	//	a = iota
	//	b = 0
	//	c = iota
	//)
	//
	//println(a, b, c)

	//const (
	//	a = iota   //0
	//	b          //1
	//	c          //2
	//	d = "ha"   //独立值，iota += 1
	//	e          //"ha"   iota += 1
	//	f = 100    //iota +=1
	//	g          //100  iota +=1
	//	h = iota   //7,恢复计数
	//	i          //8
	//)
	//fmt.Println(a,b,c,d,e,f,g,h,i)
	//test("ceshi", 11)
	//println(testString("ceshi", 11))

	//var i,j,k int
	//// 声明数组的同时快速初始化数组
	//balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//
	///* 输出数组元素 */
	//for i = 0; i < 5; i++ {
	//	fmt.Printf("balance[%d] = %f\n", i, balance[i] )
	//}
	//
	//balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	///* 输出每个数组元素的值 */
	//for j = 0; j < 5; j++ {
	//	fmt.Printf("balance2[%d] = %f\n", j, balance2[j] )
	//}
	//
	////  将索引为 1 和 3 的元素初始化
	//balance3 := [5]float32{1:2.0,3:7.0}
	//for k = 0; k < 5; k++ {
	//	fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
	//}

	//var a int = 10
	//
	//fmt.Printf("变量的地址: %x\n", &a  )

	//var a int= 20   /* 声明实际变量 */
	//var ip *int        /* 声明指针变量 */
	//
	//ip = &a  /* 指针变量的存储地址 */
	//
	//fmt.Printf("a 变量的地址是: %x\n", &a  )
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip )

	//var a1 int
	//var b1 *int
	//println(a1)
	//println(b1)
	//println(b1 == nil)

	//var user User= User{"123", 11, "444", "787"}
	//fmt.Println(user)
	//fmt.Printf("Get info: %s", user.address)

	//定义切片
	//s :=[] int {1,2,3 }
	//s = append(s, 2)
	//fmt.Println(s)
	//
	///* 创建map */
	//countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	//
	//countryCapitalMap ["123"] = "123"
	//fmt.Println("原始地图")
	//
	///* 打印地图 */
	//for country := range countryCapitalMap {
	//	fmt.Println(country, "首都是", countryCapitalMap [ country ])
	//}
	//
	///*删除元素*/ delete(countryCapitalMap, "France")
	//fmt.Println("法国条目被删除")
	//
	//fmt.Println("删除元素后地图")
	//
	///*打印地图*/
	//for country := range countryCapitalMap {
	//	fmt.Println(country, "首都是", countryCapitalMap [ country ])
	//}

	//fmt.Println(Factorial(3))

	//var a int = 3
	//var b string = string(a)
	//fmt.Println(a)
	//fmt.Println(b)
	//
	////os.Exit(1)
	//var e = 1 &^ 0
	//var f = 1 &^ 1
	//
	//print(e, "---", f)

	var s0 []int
	s0 = append(s0, 1, 2)

	println(len(s0))
	s1 := make([]int, 2, 4)

	println(len(s1), "-----", s1[0:2])
}
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

type User struct {
	name    string
	age     int
	address string
	phone   string
}

func test(name string, age int) {
	var result string = name + ":" + strconv.Itoa(age)
	fmt.Print(result)
}

func testString(name string, age int) string {
	var result string = name + ":" + strconv.Itoa(age)
	fmt.Println(result)
	return result
}
