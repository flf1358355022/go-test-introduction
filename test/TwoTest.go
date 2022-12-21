package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	println("------管道测试-------")
	//简单用法
	simpleChannel := make(chan int, 2) //设定初始容量为2,不可扩容
	simpleChannel <- 100
	simpleChannel <- 200

	//取出管道数据丢弃
	<-simpleChannel
	channelInfo := <-simpleChannel
	fmt.Printf("读取simpleChannel管道数据: %v \n", channelInfo)

	intsChannel := make(chan int, 50) //设定初始容量为50
	//协程相关在后续进阶篇中介绍，相当于开启一个新的线程执行逻辑，不阻塞main()线程的逻辑执行
	//开启协程,用于写数据
	go func() {
		for true {
			//将i写入管道
			writeNum := rand.Intn(100)
			intsChannel <- writeNum
			fmt.Printf("写入管道数据: %v \n", writeNum)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	//开启协程,用于读数据
	go func() {
		for true {
			//读取管道数据
			readInfo := <-intsChannel
			fmt.Printf("读取管道数据: %v \n", readInfo)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	//防止数据还没有在协程里打印，main函数退出,main()执行结束后其他相关的协程也会结束
	time.Sleep(time.Second * 3)

	//程序结束
}

func extendTest() {
	father := Father{name: "父亲", phone: "13100001111"}
	fmt.Printf("父亲属性: %d \n", father)

	son := new(Son)
	son.hobby = "打篮球"
	son.name = "儿子"
	son.phone = "13033331111"
	fmt.Printf("儿子属性: %d \n", son)
}

type Father struct {
	name, phone string
}

type Son struct {
	Father //匿名结构体
	hobby  string
}

func interfaceTest() {
	//根据不同的示例调用对应的实例方法
	doVoice(Dark{name: "小鸭子"})
	doVoice(new(Cat))
}

// 定义一个接口
type Animal interface { // type 接口名 interface
	voice()
}

// 定义结构提
type Dark struct {
	name string
}

type Cat struct {
	name string
}

// 实现接口方法
func (dark Dark) voice() { // 接口是引用类型，所以这里传递的是变量的引用
	fmt.Printf("%v 嘎嘎叫 \n", dark.name)
}

func (cat Cat) voice() { // 接口是引用类型，所以这里传递的是变量的引用
	fmt.Printf("喵喵叫 \n")
}

func doVoice(animal Animal) {
	animal.voice()
}

/**

func timeTest()  {

	println("-----常用日期函数测试----")
	//返回当前系统时间
	currentTime := time.Now()
	fmt.Printf("当前系统时间=%v \n", currentTime)
	//返回当前系统时间年、月、日、日、时、分、秒、毫秒
	date, month, day := currentTime.Date()
	fmt.Printf("当前系统年=%v,月=%v,日=%v \n", date, month, day)
	currentTimeYear := currentTime.Year()
	currentTimeDay := currentTime.Day()
	currentTimeMonth := currentTime.Month()
	currentTimeHour := currentTime.Hour()
	currentTimeMinute := currentTime.Minute()
	currentTimeSecond := currentTime.Second()
	//时间戳
	currentTimeUnixMilli := currentTime.UnixMilli()
	//纳秒:常用于生成随机数字(如生成订单号、随机序列等)
	currentTimeUnixNano := currentTime.UnixNano()
	//方法【Unix】将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）
	currentTimeUnix := currentTime.Unix()
	fmt.Printf("当前系统年=%v,月=%v,日=%v,时=%v,分=%v,秒=%v,当前系统毫秒数=%v,当前系统纳秒数=%v,Unix时间=%v  \n", currentTimeYear, currentTimeDay, currentTimeMonth,
		currentTimeHour, currentTimeMinute, currentTimeSecond, currentTimeUnixMilli, currentTimeUnixNano, currentTimeUnix)

	//Duration 类型用于表示两个时刻 ( Time ) 之间经过的时间，以 纳秒 ( ns ) 为单位。 点击进去可以看到是time里面的自定义类型：type Duration int64
	start_time := time.Now()
	// 空循环 uint32 的最大值次数
	const UINT32_MAX uint32 = ^uint32(0)
	var i uint32
	for i = 0; i < UINT32_MAX; i += 1 {

	}
	end_time := time.Now()
	spand_time := time.Duration(end_time.Sub(start_time))
	fmt.Printf("空循环 uint32 的最大值次数耗时时间(s):%v \n", spand_time.Seconds())

	println("------时间格式化-----")
	var now = time.Now()
	// 以下的数字都是固定的值，不能更换,据说2006/01/02 15:04:05是创始人思考创建go的时间
	fmt.Println(now.Format("2006")) // 2022
	fmt.Println(now.Format("01"))   // 04
	fmt.Println(now.Format("02"))   // 30
	fmt.Println(now.Format("15"))   // 10
	fmt.Println(now.Format("04"))   // 52
	fmt.Println(now.Format("05"))   // 16

	// 数字之外的其它字符可以更换
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 2022/04/30 10:52:16
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2022-04-30 10:52:16

	println("------sleep练习-----")
	sleep_start_time := time.Now()
	var num = 1
	for {
		fmt.Printf("%v ", num)
		//休眠30ms
		time.Sleep(time.Millisecond * 30)
		if num == 5 {
			println()
			break
		}
		num++
	}
	sleep_end_time := time.Now()
	sleep_spand_time := sleep_end_time.Sub(sleep_start_time)
	fmt.Printf("sleep 测试耗费时间(ms):%v \n", sleep_spand_time.Milliseconds())

	println("------时间戳 <-互转-> 日期字符串-----")
	// 时间戳转换年月日时分秒（一个参数是秒，另一个参数是纳秒）
	//Unix返回与给定Unix时间相对应的本地时间，
	//秒和纳秒。
	var time_1 = time.Unix(1595289901, 0)
	var timeStr = time_1.Format("2006-01-02 15:04:05")
	fmt.Println("时间戳转时间字符串结果：%v \n", timeStr)

	// 日期字符串转换成时间戳
	var timeStr2 = "2022-11-25 14:44:52"
	var tmp = "2006-01-02 15:04:05" //转换模版
	timeObj5, _ := time.ParseInLocation(tmp, timeStr2, time.Local)
	fmt.Println("日期字符串转换成ms时间戳结果：%v \n", timeObj5)
}

func testPanicAndRecover() {
	//测试异常捕捉处理
	testExceptionCapture()

	fmt.Println("异常捕捉后继续处理流程")

	//测试异常panic
	testPanic()

	//上面函数执行报错后如果没有异常捕捉处理程序直接结束运行
	fmt.Println("异常捕捉处理测试")
}

func testPanic() {
	n1 := 1
	n2 := 0
	n3 := n1 / n2

	//发送异常之后，下面的输出语句不会输出,程序直接结束
	fmt.Println("res:", n3)
}

func testExceptionCapture() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到的异常信息: ", err)
			//异常之后做一些事情
			fmt.Println("发送钉钉告警或者邮件告警等")
		}
	}()

	n1 := 1
	n2 := 0
	n3 := n1 / n2

	//发送异常之后，下面的输出语句不会输出
	fmt.Println("res:", n3)
}

func testStringsUtils() {
	//统计字符串的长度,按字节 len(str)
	var str string = "hello word"
	fmt.Printf("str长度:%v \n", len(str))
	//字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str)
	for i := range r {
		fmt.Printf("遍历字符串:%v \n", i)
	}
	//字符串转整数： n , err := strconv.Atoi(“12”)
	number_1, _ := strconv.Atoi("12")
	number_2, err_2 := strconv.Atoi("test")
	fmt.Printf("字符串转整数,number_1:%v \n", number_1)
	fmt.Printf("字符串转整数,number_2:%v, 错误信息:%v \n", number_2, err_2)

	//整数转字符串: str = strconv.Itoa(12345)、strconv.FormatInt(999, 10)
	str = strconv.Itoa(12345)
	fmt.Printf("整数转字符串,str:%v \n", str)
	str = strconv.FormatInt(999, 10)
	fmt.Printf("整数转字符串,str:%v \n", str)
	//字符串转[]byte: var bytes= []byte(“hello go”)
	var bytes = []byte("hello word")
	fmt.Printf("字符串转byte数组,str:%v \n", bytes)
	//[]byte转字符串: str = string([]byte{97,98,99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("byte转字符串,str:%v \n", str)
	//10进制转2,8,16进制: str = strconv.FormatInt(123,2) // 2->8,16S
	str = strconv.FormatInt(123, 2)
	fmt.Printf("数字转2进制字符串,str:%v \n", str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("数字转8进制字符串,str:%v \n", str)
	str = strconv.FormatInt(123, 10)
	fmt.Printf("数字转10进制字符串,str:%v \n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("数字转16进制字符串,str:%v \n", str)

	//查找子串是否在指定的字符串中: strings.Contains(“seafood”, “foo”) //true
	isContains := strings.Contains("food foo eat", "foo")
	fmt.Printf("是否包含字符串foo,isContains:%v \n", isContains)

	//统计一个字符串有几个指定的子串:strings.Count(“ceheese”, “e”) //4
	countNum := strings.Count("hello word", "l")
	fmt.Printf("字符串含有%d个l:%v \n", countNum)

	//不区分大小写的字符串比较(== 是区分字母大小写的): fmt.PrintIn(strings.EqualFold(“abc”, “Abc”) // true
	isEqual := strings.EqualFold("abc", "Abc")
	fmt.Printf("字符串不区分大小写比较是否相等:%v \n", isEqual)
	fmt.Printf("字符串区分大小写比较是否相等:%v \n", "abc" == "Abc")

	//返回子串在字符串第一次出现的index值，如果没有返回-1 : strings.Index(“NLT_abc”,”abc”) //4
	firstIndex := strings.Index("hello word hello word", "ll")
	fmt.Printf("字符子串在字符串中第一次出现的地方:%v \n", firstIndex)

	//返回子串在字符串最后一次出现的index，如没有返回-1 : strings.LastIndex(“go golang” , “go”)
	lastIndex := strings.LastIndex("hello word hello word", "ll")
	fmt.Printf("字符子串在字符串中最后一次出现的地方:%v \n", lastIndex)

	//将指定的子串替换成 另外一个子串: strings.Replace(“go go hello” , “go”, “go语言”, n) n 可以指定你希望替换几个,如果n = -1表示全部替换
	str = strings.Replace("hello word hello word", "ll", "ll-", 2)
	fmt.Printf("替换后的字符串:%v \n", str)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组:
	splitData := strings.Split("Mr Li,20,北京朝阳区", ",")
	fmt.Printf("分割后的数据:%v \n", splitData)

	//将字符串的字母进行大小写的转换: strings.ToLower(“Go”) // go strings.ToUpper(“Go”) //GO
	str = strings.ToLower("HELLO Word")
	fmt.Printf("转换成小写后的数据:%v \n", str)
	str = strings.ToUpper("hello Word")
	fmt.Printf("转换成大写后的数据:%v \n", str)

	//将字符串左右两边的空格去掉 : strings.TrimSpace(“ tn a lone gopher ntrn “)
	str = strings.TrimSpace(" hello, word ")
	fmt.Printf("去掉空格后的数据:%v \n", str)

	//将字符串左右两边指定的字符去掉: strings.Trim(“! hello! “, “ !”) // [“hello”]//将左右两边!和””去掉
	str = strings.Trim("! hello, word !", "!")
	fmt.Printf("去掉!后的数据:%v \n", str)

	//将字符串左边指定的字符去掉: strings.TrimLeft(“! hello! “,” !”) // [“hello”]//将左边!和”“去掉
	str = strings.TrimLeft("! hello, word !", "!")
	fmt.Printf("去掉左边!后的数据:%v \n", str)

	//将字符串右边指定的字符去掉: strings.TrimRight(“! hello! “,” !”) // [“hello”]//将右边!和””去掉
	str = strings.TrimRight("! hello, word !", "!")
	fmt.Printf("去掉右边!后的数据:%v \n", str)

	//判断字符串是否以指定的字符串开头: strings.HasPrefix(“ftp://192.168.10.1" ,”ftp”) // true
	hasPrefix := strings.HasPrefix("! hello, word !", "!")
	fmt.Printf("是否以!开头:%v \n", hasPrefix)

	//判断字符串是否以指定的字符串结束: strings.HasSuffix(“‘NLT_abc.jpg”,”abc”) //false
	hasSuffix := strings.HasSuffix("! hello, word !", "!")
	fmt.Printf("是否以!结束:%v \n", hasSuffix)

	//判断字符串是否包含一个字符串
	hasContains := strings.Contains("hello word", "hello")
	fmt.Printf("字符串hello word是否包含hello:%v \n", hasContains)
}

type MyInt int64

func (receiver MyInt) testMethod(number_1 int64) string {
	fmt.Println("this is a data type receiver,", receiver)
	return "-------" + strconv.FormatInt(number_1, 10) + "-------"
}

type Person struct {
	Name string
	Age  int
}

func (receiver Person) testMethodV2() Person {
	fmt.Println("this is a struct receiver,", receiver)
	return receiver
}

func methodTest() {
	var number_test MyInt = 888
	result_num := number_test.testMethod(999)
	println("testMethod invoke result:", result_num)

	personInfo := Person{Age: 20, Name: "Mr F"}
	result := personInfo.testMethodV2()
	fmt.Printf("testMethodV2 invoke result:%v, type:%T", result, result)

}

func pkgTest() {
	//result_number := packageUtils.NumberUtils(10, 20, "*")
	//fmt.Println("包调用函数返回值:", result_number)
}

func testTransmit() {
	//值传递测试
	var number int64 = 999
	testV1(number)
	//从这里看出调用函数赋值后并不会改变原值
	fmt.Println("main()当前number=", number)

	//应用传递测试
	var i int = 10
	fmt.Printf("i当前的值 = %v \n", i)
	//1、ptr是一个指针变量 2、ptr存的是i变量的地址 3、类型是*int
	var ptr *int = &i
	//赋予新值
	*ptr = 20
	fmt.Printf("指针存储的值:%v, 类型:%T, 占内存字节数:%d, 指针存储地址指向的值:%d, 指针的地址:%v \n", ptr, ptr, unsafe.Sizeof(ptr), *ptr, &ptr)
	fmt.Printf("i修改后的值 = %v \n", i)
}

func testV1(number int64) {
	number = 1000
	fmt.Println("testV1()当前number=", number)
}

func testStruct() {
	//使用起了别名的类型
	var number_test myIntType = 999
	fmt.Printf("number_test = %v, 类型:%T, 大小:%v字节  \n", number_test, number_test, unsafe.Sizeof(number_test))

	myUserMsg := MyUserMsg{name: "123", phone: "13111110000"}
	fmt.Printf("myUserMsg = %v, 类型:%T, 大小:%v字节  \n", myUserMsg, myUserMsg, unsafe.Sizeof(myUserMsg))

}

// 定义类型
type UserMsg struct {
	name, phone string
}

// 类型起别名
type myIntType int64

// 结构体起别名
type MyUserMsg UserMsg


*/
