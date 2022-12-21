package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type UserInfo struct {
	//用户姓名
	name string
	//年龄、身高
	age, height int16 //相同类型可以定义在一行
	//用户地址
	address string
}

// 如果里面的字段类型都一样可以使用这种方法
type UserInfoV2 struct{ name, address string }

// 组合型
type UserInfoV3 struct {
	user_1 UserInfo
	user_2 UserInfoV2
}

func main() {

	funcTest()

	/**
	//初始化结构体
	info := UserInfo{age: 18, name: "小明", height: 180, address: "浙江省杭州市"}
	fmt.Printf("结构体信息:%v  \n", info)
	//访问结构体属性，info.xxx
	fmt.Printf("访问name信息:%v, age:%v, 访问height信息:%v, 访问address信息:%v   \n", info.name, info.age, info.height, info.address)
	//分配内存地址连续
	fmt.Printf("name分配内存地址:%v, age分配内存地址:%v, height分配内存地址:%v, address分配内存地址:%v   \n", &info.name, &info.age, &info.height, &info.address)
	//组合型结构体
	user_1 := UserInfo{age: 33, name: "Mir Zhang", height: 180, address: "北京"}
	user_2 := UserInfoV2{name: "Mir Li", address: "北京"}
	user_3 := UserInfoV3{user_1, user_2}
	fmt.Printf("结构体信息:%s  \n", user_3)

	*/

}

/**   如下所示，函数定义
func 函数名 (参数列表) （返回值列表） { //返回值只有一个时可以不写（）

	//函数体，功能执行
	return 返回值列表

}
*/
// ///////////return///////////////
func returnTest(n1 int64, n2 int64) (string, int64) {
	//return:用于函数执行结果的值返回，可以返回一个或者多个参数
	result_1 := strconv.FormatInt(n1*n2, 10)
	result_2 := n1 + n2
	return result_1, result_2
}

// 上面的写法还可以这些写
func returnTestV2(n1 int64, n2 int64) (result_1 string, result_2 int64) {
	result_1 = strconv.FormatInt(n1*n2, 10)
	result_2 = n1 + n2
	return
}

// //////////////普通函数//////////////////
// 单个入参、单个出参
func testFuntion(number int) int { //()里的是入参, {左边的是返回值类型和参数， 可以写(result, int)
	return number * 10
}

func testFuntionV2(number int, value string) (a int, b string) {
	//转化成10进制，进行字符串拼接
	resultStr := strconv.FormatInt(int64(number), 10) + value
	return number, resultStr
}

// 多个入参、单个出参
func testFuntionV3(args ...int) (int, string, float64) { //(args ...int)表示可以传递多个参数，可以理解为0到多个参数
	resultNum := 0
	for arg := range args {
		fmt.Println("打印参数arg = ", arg)
		resultNum += arg
	}
	resultStr := "返回值2"
	resultFloatNum := 888.88
	fmt.Printf("最终返回值1=%v, 最终返回值2=%v, 最终返回值3=%v \n", resultNum, resultStr, resultFloatNum)
	return resultNum, resultStr, resultFloatNum
}

// init函数,最大的作用是用来初始化源文件，该函数会在main函数执行前被调用
func init() {
	//do something
	fmt.Println("初始化函数，先于main函数执行")
}

// ///////////匿名函数(全局变量)///////////////
var (
	globalVariableFunc = func(number_1 int, number_2 int) (int, string) {
		return number_2 * number_1, strconv.FormatInt(int64(number_1), 10)
	}
)

// ///////////闭包///////////////
// 含义:闭包是由函数和与其相关的引用环境组合而成的实体[抽象、难以理解！]，其实就是:匿名函数+外部引用
// 为什么使用闭包(为了避免全局变量被滥用):可以让变量常驻内存、可以让变量不污染全局
// 可参考文章:https://blog.csdn.net/qq_27654007/article/details/116667624
func closureTest() func(int) int {
	var n int = 10
	return func(x int) int {
		//每一次调用都会给n进行累加赋值，n的值在内存中伴随整个闭包实例的整个生命周期
		n = n + x
		return n
	}
}

// ///////////defer///////////////
func deferTest(strValue string) string {

	//依据打印输出可以看出来，defer这行的逻辑是在return那一刻执行的
	defer printString("defer 延迟执行测试")
	printString("deferTest 执行测试-1")
	printString("deferTest 执行测试-2")
	return strValue
}

func printString(str string) {
	fmt.Println(str)
}

// ///////////函数作为参数/////////////
func testFuncArg(param func(str string)) {
	param("函数作为参数测试")
}

// ///////////函数作为返回值/////////////
func testReturnFunc() func(result_1 int64, result_2 int64) int64 {
	return func(n1 int64, n2 int64) int64 {
		return n1 * n2
	}
}

////////////////////////////////////进阶//////////////////////////////////////

//func main() {
//
//	returnValue_1, returnValue_2 := testFunction(10)
//	//打印返回值
//	fmt.Printf("returnValue_1 = %v, returnValue_2 = %v", returnValue_1, returnValue_2)
//}

// testFunction:函数名, value:函数入参， b、c:函数执行返回参数
func testFunction(value int) (b int, c int) {
	fmt.Println("value = ", value)
	return 10, 10
}

//func main() {
//	rand.Seed(time.Now().Unix())
//	rangerNum := 0
//	for i := 0; i < 10; i++ {
//		rangerNum = rand.Intn(100) + 1
//		fmt.Println(rangerNum)
//	}
//
//}

// /////////////全局变量////////////////
//var age, name = 18, "小明"

// /////////////全局变量///////////////
func abc_1() {

	////创建变量
	//var a int
	////变量赋值
	//a = 10
	////使用变量(进行打印)
	//fmt.Println("a = ", a)

	//创建变量
	//变量=变量名+值+数据类型，这一点请大家注意，变量的三要素
	//Golang 的变量如果没有赋初值，编译器会使用默认值, 比如 int 默认值 0 string 默认值为空串， 小数默认为 0
	///////////////局部变量///////////////

	//1、创建变量但 不赋值 使用的默认值(int 默认值为0)
	var a int
	//使用变量(进行打印)
	fmt.Println("a = ", a)

	//2、创建变量但不赋予类型，根据值自行判定变量类型(类型推导)
	var b = 10
	//fmt.Printf格式化输出 %v:输出值， %T:输出值的类型 \n:换行
	fmt.Printf("b = %v, type: %T  \n", b, b)

	//3、使用:=声明变量
	//c := "Go" 等价于 var c string = "Go";
	c := "Go"
	fmt.Printf("c = %v, type: %T \n", c, c)

	//4、多变量声明,带值声明或者不带值声明
	var d, e, f int              //不带值
	var d1, e1, f1 int = 1, 2, 3 //带值(相同类型)
	d2, e2, f2 := 1, "123", 's'  //带值(相同类型)
	fmt.Println(d, e, f, "------", d1, e1, f1, "------", d2, e2, f2)

	//注意⚠️:在同一代码块中，不能对类型再次赋予其他类型
	//var testA int = 10;
	//testA  := 20; // ❌

	///////////////局部变量///////////////

	//输出全局变量
	//fmt.Println(age, name)

	//又详细又易懂的快读入门，这一篇足矣
}

func baseDataTest() {

	println("----byte类型----")
	//Golang 程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量 使用占用空间小的数据类型。【如：年龄】
	//bit: 计算机中的最小存储单位。byte:计算机中基本存储单元。[二进制再详细说] 1byte = 8 bit
	var n0 byte = 10
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n0, n0, unsafe.Sizeof(n0))

	println("----int类型----")
	//int 类型
	var n1 int = 20
	var n2 int8 = 21
	var n3 int16 = 22
	var n4 int32 = 23
	var n5 int64 = 24
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n1, n1, unsafe.Sizeof(n1))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n2, n2, unsafe.Sizeof(n2))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n3, n3, unsafe.Sizeof(n3))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n4, n4, unsafe.Sizeof(n4))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", n5, n5, unsafe.Sizeof(n5))

	println("----float类型----")
	//float 类型
	var f1 float32 = 30.23748734872346 //会损失精度，如果我们要保存一个精度高的数，则应该选用 float64
	var f2 float64 = 21.23748734872346
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", f1, f1, unsafe.Sizeof(f1))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", f2, f2, unsafe.Sizeof(f2))

	println("----字符类型----")
	//字符类型(英文字母占1 个字节 汉字占3 个字节)， 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的 unicode 字符
	//Go 语 言 的 字 符 使 用 UTF-8 编 码 ， 如 果 想 查 询 字 符 对 应 的 utf8 码 值 http://www.mytju.com/classcode/tools/encode_utf8.asp
	//不理解utf-8、ASCII和Unicode可以看考这篇文章:https://www.bilibili.com/read/cv16001311/
	var c0 int = 's'
	var c1 int = '北'
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d, 显示信息=%c \n", c0, c0, unsafe.Sizeof(c0), c0)
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d, 显示信息=%c \n", c1, c1, unsafe.Sizeof(c1), c1)

	println("----布尔类型----")
	//bool 类型占 1 个字节。适于逻辑运算，一般用于程序流程控制
	var flagControl bool = true
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", flagControl, flagControl, unsafe.Sizeof(flagControl))
	for i := 0; i < 100; i++ {
		if i == 10 {
			//do something
			flagControl = false
		}
		if flagControl == false {
			fmt.Println("结束循环, i= ", i)
			break
		}
	}

	println("----string 类型----")
	//字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本
	//Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本，这样 Golang 统一使用 UTF-8 编码,中文 乱码问题不会再困扰程序员。
	var str string = "测试字符串"
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", str, str, unsafe.Sizeof(str))
	//注意⚠️:字符串一旦赋值了，字符串就不能修改了：在 Go 中字符串是不可变的。
	//var str1 string = "abc"
	//str1[0] = 'c' //❌ 不可取，不可修改里面的内容
	//字符串的两种表示形式 (1) 双引号, 会识别转义字符 (2) 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果 【案例演示】
	//方式1
	var str2 string = "测试字符串"
	//方式2
	var str3 string = `
			测试字符串表示形式，
			我是以字符串的原生形式输出，可以实现防止攻击、输出源代码等效果
			var str4 string = "4"
	`
	//字符串拼接方式 "" + "",当一行字符串太长时，需要使用到多行字符串，可以如下处理
	var str5 string = "hello" + "word" + "hello" + "word" + "hello" + "word" +
		"test" + "str"
	fmt.Println(str2, str3, str5)

	println("----指针类型----")
	//基本数据类型，变量存的就是值，也叫值类型。 获取变量的地址，用&，比如： var num int, 获取 num 的地址：&num
	//指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	var i int = 10
	//1、ptr是一个指针变量 2、ptr存的是i变量的地址 3、类型是*int
	var ptr *int = &i
	fmt.Printf("指针存储的值:%v, 类型:%T, 占内存字节数:%d, 指针存储地址指向的值:%d, 指针的地址:%v \n", ptr, ptr, unsafe.Sizeof(ptr), *ptr, &ptr)

	println("-----基本数据类型的相互转换----")
	////////////////基本数据类型的相互转换///////////////////
	//Golang 和 java / c 不同，Go 在不同类型的变量之间赋值时需要显式转换。也就是说 Golang 中数 据类型不能自动转换。
	var int_1 int = 100
	var float1 float32 = float32(int_1)
	var int_2 int8 = int8(int_1)

	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", float1, float1, unsafe.Sizeof(float1))
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", int_2, int_2, unsafe.Sizeof(int_2))
	//在转换中，比如将 int64 转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按 溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围.
	var int_3 int64 = 999
	var int_4 int8 = int8(int_3) //转换的时候值溢出
	fmt.Printf("值:%v, 类型:%T, 占内存字节数:%d \n", int_4, int_4, unsafe.Sizeof(int_4))

	//基本类型和string互转
	//方式 1：fmt.Sprintf("%参数", 表达式)
	var number1 int = 999
	var number2 float64 = 999.999
	var bool_1 bool = false
	var byte_1 byte = 'a'
	var str_1 string

	str_1 = fmt.Sprintf("%d", number1)
	fmt.Printf("1、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	str_1 = fmt.Sprintf("%f", number2)
	fmt.Printf("2、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	str_1 = fmt.Sprintf("%t", bool_1)
	fmt.Printf("3、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	str_1 = fmt.Sprintf("%c", byte_1)
	fmt.Printf("4、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))
	//方式 2：使用 strconv 包的函数
	var number4 int32 = 888
	var number5 float64 = 888.888
	var bool_2 bool = false

	str_1 = strconv.FormatInt(int64(number4), 10)
	fmt.Printf("5、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	//FormatFloat(f float64, fmt byte, prec, bitSize int),参数依次表示为转换的值、转换的格式、保留小数后多少位、是转换成float64还是float32
	str_1 = strconv.FormatFloat(number5, 'f', 2, 64)
	fmt.Printf("6、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	str_1 = strconv.FormatBool(bool_2)
	fmt.Printf("7、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", str_1, str_1, unsafe.Sizeof(str_1))

	//string 类型转基本数据类型
	var str_2 string = "true"
	var bool_3 bool

	//strconv.ParseBool返回两个值, _标识忽略返回的第二个值
	bool_3, _ = strconv.ParseBool(str_2)
	fmt.Printf("8、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", bool_3, bool_3, unsafe.Sizeof(bool_3))
	//	strconv.ParseFloat()同理
	//	strconv.ParseInt()同理
	//	strconv.ParseUint()同理

	//注意⚠️事项:在将 String 类型转成 基本数据类型时，要确保 String 类型能够转成有效的数据
	//比如 我们可以 把 "123" ,转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 ， 其它类型也是一样的道理. float => 0 bool => false
	var str_4 string = "hello"
	var int_5 int64
	int_5, _ = strconv.ParseInt(str_4, 10, 64)
	fmt.Printf("9、转换后的值:%v, 类型:%T, 占内存字节数:%d \n", int_5, int_5, unsafe.Sizeof(int_5))

}

func abc_3() {

	var num_0 int = 100
	var num_1 int = 100
	var num_2 int = num_1 + num_0/num_0*num_1
	println(num_2)

	// 演示 % 的使用
	//看一个公式 a % b = a - a / b * b
	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)

	//i++, i--
	var i_0 int8 = 0
	i_0++
	var i_1 int8 = 0
	i_1--
	fmt.Printf("i_0 = %d, i_1 = %d \n", i_0, i_1)
	//Golang 的自增自减只能当做一个独立语言使用时，不能这样使

	//var i_3 int = i_0 ++    //❌ 错误用法
	//if i_0 ++ > 1{}    //❌ 错误用法

}

func logicalOperatorTest() {

	//逻辑运算符
	var number_0 int8 = 10
	var number_1 int8 = 15

	if number_1 == 11 || number_0 < 10 {
		println("我是1")
	} else if number_1 == 15 && number_0 == 10 {
		println("我是2")
	} else if !(number_1 == 15) {
		println("我是3")
	}

}

func scanlnTest() {

	var name string
	var age int
	var address string

	//方式1:fmt.Scanln
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入地址")
	fmt.Scanln(&address)

	fmt.Printf("姓名:%v, 年龄:%v, 地址:%v   \n", name, age, address)
	//方式2:fmt.Scanf
	println("----------------------")
	var name_1 string
	var age_1 int
	var address_1 string
	println("请输入你的姓名、年龄、地址，依次按照空格隔开")
	fmt.Scanf("%s %d %s", &name_1, &age_1, &address_1)
	fmt.Printf("姓名:%v, 年龄:%v, 地址:%v   \n", name_1, age_1, address_1)

}

func processControlTest() {

	println("-------if else流程控制------")
	//流程控制
	num := 10

	//单分支语句
	if num <= 10 {
		//满足条件进入流程
		fmt.Println("单分支测试")
	}

	//双分支语句
	if num > 10 {
		fmt.Println("双分支测试01")
	} else {
		fmt.Println("双分支测试02")
	}

	//多分支语句，注意⚠️:双分支语句或者多分支语句执行后只会进入到一个流程当中，代码由上而下执行，只要满足即出流程不会就算是其他分支满足条件也不会进入到分支里
	if num > 10 {
		fmt.Println("多分支测试01")
	} else if num == 10 {
		fmt.Println("多分支测试02")
	} else {
		fmt.Println("多分支测试03")
	}

	//嵌套语句(在一个分支结构中又完整的嵌套了另一个完整的分支结构，里面的分支的结构称为内层分 支外面的分支结构称为外层分支。)
	num_01 := 20
	if num > 10 {
		fmt.Println("多分支测试04")
	} else if num == 10 {
		if num_01 > 20 {
			fmt.Println("嵌套分支测试01")
		} else if num_01 == 20 {
			fmt.Println("嵌套分支测试02")
		}
		fmt.Println("分支测试05")
	}

	println("-------switch case流程控制------")
	//switch case流程控制
	currentTemperature := 25
	switch currentTemperature {
	case 10:
		fmt.Println("当前天气凉")
	case 25:
		fmt.Println("当前天气温暖～")
	case 40:
		fmt.Println("当前天气热死了！")
	default:
		//以上条件都不符合才会进入这里
		fmt.Println("这鬼天气多变！")
	}

}

func loopTest() {

	//简单示例，循环10次
	for i := 0; i < 10; i++ {
		fmt.Println("循环次数 i = ", i)
	}

	//上述还可以这样写
	j := 0
	for j < 10 {
		fmt.Println("循环次数 j = ", j)
		j++
	}

	println("------break用法-----")
	//还可以这样写
	k := 0
	for { //等价于 for ; ;
		if k >= 10 {
			//退出循环
			break
		}
		fmt.Println("循环次数 k = ", k)
		k++
	}
	println("------continue用法-----")
	//continue用法
	l := 0
	for l < 10 {
		if l == 2 {
			fmt.Println("我是continue测试, l = ", l)
			l++
			continue
		}
		fmt.Println("循环测试, l = ", l)
		l++
	}

	println("------goto测试-----")
	flag := false
	if flag == false {
		goto label_01
	}
	println("------goto测试----1-----")
	println("------goto测试----2-----")

label_01: //跳转标识，上述输出不会打印
	println("------goto测试----3-----")
	println("------goto测试----4-----")

	println("------return测试-----")
	returnTestNum := 0

	if returnTestNum == 0 {
		println("-----到此为止了！-----")
		return
	}
	println("-----执行动作1-----")
	println("-----执行动作2-----")

}

func arrayTest() {

	//创建数组
	var arr_0 [4]int = [4]int{1, 2, 3, 4}
	var arr_1 = [4]float64{111.000, 222, 333, 444.444}
	var arr_2 = [...]bool{false, true, true}      // 自行判断长度，中括号里...一个不能少
	var arr_3 = [...]byte{1: 'a', 0: 'b', 2: 'c'} // 指定索引和值

	fmt.Printf("arr_0= %v, 类型:%T \n", arr_0, arr_0)
	fmt.Printf("arr_1= %v, 类型:%T \n", arr_1, arr_1)
	fmt.Printf("arr_2= %v, 类型:%T \n", arr_2, arr_2)
	fmt.Printf("arr_3= %v, 类型:%T \n", arr_3, arr_3)

}

func sliceTest() {

	//数组
	array_0 := [...]int{1, 2, 3, 4}

	//切片(将数组下标为1到3的数据取出来,不含3)
	slice := array_0[1:3]
	fmt.Printf("slice 内容:%v \n", slice)
	fmt.Printf("数组下标为1的数据为:%v, 数组下标为2的数据为:%v \n", array_0[1], array_0[2])

	//通过make创建一个数组，但是过程对我们是不可见的
	sl := make([]int, 5)
	fmt.Printf("切片sl类型:%T, 数据为:%v, 长度:%v \n", sl, sl, len(sl))
	println("-------追加内容测试-------")
	//追加数据到切片中,尾部追加
	sl = append(sl, 1, 2, 3, 3, 4)
	fmt.Printf("尾部追加切片sl类型:%T, 数据为:%v 长度:%v \n", sl, sl, len(sl))
	//追加切片到切片中,尾部追加
	sl = append(sl, []int{1, 2, 3}...)
	fmt.Printf("尾部追加切片后切片sl类型:%T, 数据为:%v 长度:%v \n", sl, sl, len(sl))
	//追加数据到切片中,头部追加
	sl = append([]int{0, 1, 2}, sl...)
	fmt.Printf("头部追加切片sl类型:%T, 数据为:%v 长度:%v \n", sl, sl, len(sl))

	println("-------合并内容测试-------")
	//合并两个切片内容,注意⚠️:切片相同下标的元素会被合并的元素覆盖
	copy(sl, slice)
	fmt.Printf("合并后切片sl类型:%T, 数据为:%v 长度:%v \n", sl, sl, len(sl))

	strArray_1 := [...]string{"123", "234", "456", "999"}
	strArray_2 := [...]string{"111", "222", "333", "888"}

	strSlice_1 := strArray_1[0:4]
	strSlice_2 := strArray_2[1:4]

	fmt.Printf("原有strSlice_1 内容:%v \n", strSlice_1)
	fmt.Printf("原有strSlice_2 内容:%v \n", strSlice_2)
	copy(strSlice_1, strSlice_2)

	fmt.Printf("合并后strSlice_1 内容:%v \n", strSlice_1)
	fmt.Printf("合并后strSlice_2 内容:%v \n", strSlice_2)

	println("-------删除内容测试-------")
	//删除内容
	newSlice := []int{1, 2, 3}
	newSlice = newSlice[1:] //从下标1开始取所有数据，相当于删除了下标为0(即第一个元素)
	fmt.Printf("删除后newSlice 内容:%v \n", newSlice)
	//删除尾部元素
	newSlice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	newSlice = newSlice[0 : len(newSlice)-1] //删除最后一个元素
	fmt.Printf("删除后newSlice 内容:%v \n", newSlice)
	newSlice = newSlice[0 : len(newSlice)-3] //删除最后3个元素
	fmt.Printf("删除后newSlice 内容:%v \n", newSlice)

	//进行略复杂删除，使用append删除
	newArray_01 := [...]string{"a", "b", "c", "d"}
	stringsSlice_01 := append(newArray_01[:2], newArray_01[3]) //删除中间下标为2的元素,公式即append(newArray_01[:i], newArray_01[i + 1]),i= 2,删除多个可以参考
	fmt.Printf("删除后newArray_01 内容:%v \n", stringsSlice_01)

	//进行略复杂删除，使用copy删除
	newArray_02 := [...]string{"e", "f", "g", "h"}
	fmt.Printf("%v, ---, %v \n", newArray_02[:2], newArray_02[3:])
	//看输出内容加以理解，此处将元素["e", "f"] ["h"] 合并后 变为 ->["h", f],可以看下上面切片合并加以理解
	//按照需求切分，此处比较绕，多思考
	copyAfterLen := copy(newArray_02[:2], newArray_02[3:])
	fmt.Printf("copy 后长度:%v newArray_02内容:%v \n", copyAfterLen, newArray_02)
	stringsSlice_02 := newArray_02[:copyAfterLen]
	fmt.Printf("删除后newArray_02 内容:%v \n", stringsSlice_02)

}

func mapTest() {
	//初始化一个map,长度为10
	stringStringMap := make(map[string]string, 10)
	//向map增加元素
	stringStringMap["name"] = "Mir Li"
	stringStringMap["age"] = "18"

	fmt.Printf("stringMap 内容:%v 类型：%T, 长度:%d \n", stringStringMap, stringStringMap, len(stringStringMap))
	//覆盖元素操作
	stringStringMap["name"] = "Mir Zhang"
	fmt.Printf("stringMap 内容:%v 类型：%T, 长度:%d \n", stringStringMap, stringStringMap, len(stringStringMap))
	//删除元素操作
	delete(stringStringMap, "age")
	fmt.Printf("stringMap 内容:%v 类型：%T, 长度:%d \n", stringStringMap, stringStringMap, len(stringStringMap))
	//判断是否存在元素
	_, present := stringStringMap["name"]
	fmt.Printf("stringStringMap 元素是否存在%v元素:%v \n", "name", present)

	//其他初始化方式，初始即赋值
	stringIntMap := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Printf("stringIntMap 内容:%v 类型：%T, 长度:%d \n", stringIntMap, stringIntMap, len(stringIntMap))
}

func funcTest() {
	println("-------return测试---------")
	result_1, result_2 := returnTest(88, 99)
	result_3, result_4 := returnTestV2(88, 99)
	fmt.Printf("returnTest执行结果, result_1 = %v, result_2 = %v \n", result_1, result_2)
	fmt.Printf("returnTestV2执行结果, result_3 = %v, result_4 = %v \n", result_3, result_4)

	println("-------函数测试---------")
	testFuntion(1)
	funcResultV2_1, funcResultV2_2 := testFuntionV2(1, "123")
	fmt.Printf("funcResultV2_1 = %v, funcResultV2_2 = %v \n", funcResultV2_1, funcResultV2_2)
	testFuntionV3()
	testFuntionV3(1, 2, 3, 4, 5)
	println("-------匿名函数测试---------")
	//匿名函数，没有名字的函数，放在代码块中直接执行
	anonymousFuncResult_1, _ := func(n1 int, n2 int) (int, float64) {
		return n1 * n2, float64(n1 * n2)
	}(2, 8) //尾部的括号里传递参数

	fmt.Println("anonymousFuncResult = ", anonymousFuncResult_1)
	//也可以直接把匿名函数赋值给变量，但是赋值给变量之前不能给匿名函数传递参数
	a := func(n1 int, n2 int) (int, int) {
		return n2, n1
	} //尾部没有参数

	n1 := 10
	n2 := 29
	n1, n2 = a(n1, n2)
	fmt.Printf("匿名函数赋值给变量之后n1 = %v, n2 = %v \n", n1, n2)

	println("-------全局匿名函数测试---------")
	//调用定义好的匿名函数全局变量
	globalVariableFuncResult_1, globalVariableFuncResult_2 := globalVariableFunc(1, 3)
	fmt.Printf("调用全局匿名函数结果globalVariableFuncResult_1 = %v type: %T, globalVariableFuncResult_1 = %v type: %T\n",
		globalVariableFuncResult_1, globalVariableFuncResult_1, globalVariableFuncResult_2, globalVariableFuncResult_2)
	//闭包调用
	f2 := closureTest()
	//依下述输出来看,虽然我们没有显性的声明一个全局变量，但是我们每次调用都会进行累加
	fmt.Printf("闭包调用第1次返回结果:%v \n", f2(10))
	fmt.Printf("闭包调用第2次返回结果:%v \n", f2(10))
	fmt.Printf("闭包调用第3次返回结果:%v \n", f2(10))
	f3 := closureTest()
	fmt.Printf("新的闭包调用第1次返回结果:%v \n", f3(10)) //新的实例运行后并不会原先闭包函数中的变量

	println("-------defer测试---------")
	//defer:defer是go中一种延迟调用机制，defer后面的函数只有在当前函数执行完毕后才能执行，通常用于释放资源。
	//参考资料:https://blog.csdn.net/m0_46251547/article/details/123762669
	deferTest("字符参数")

	println("-------函数作为参数测试---------")
	testFuncArg(printString)

	println("-------函数作为返回值测试---------")
	returnFunc := testReturnFunc()
	fmt.Printf("函数作为返回值测试, 返回结果类型 = %T \n", returnFunc)
	finalResult := returnFunc(100, 200)
	fmt.Printf("执行函数, 结果 = %d, 类型 = %T \n", finalResult, finalResult)
	println("-------内置函数测试---------")
	//1. len : 用来求长度，比如string、array、slice、map、channel
	//2. new : 用来分配内存，主要用来分配值类型，比如int、float32、struct…返回的是指针
	//3. make:用来分配内存，主要用来分配引用类型，比如chan、map、slice。
	//值类型的用new，返回的是一个指针
	p := new(int)
	fmt.Println("*p = ", *p, ", p = ", p)
	*p = 29
	fmt.Println("*p = ", *p)
	//引用类型的用make
	stringStringMap := make(map[string]string, 10)
	//向map增加元素
	stringStringMap["name"] = "Mir Li"
	stringStringMap["age"] = "18"

}
