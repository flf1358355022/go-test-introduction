package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println("第", (index + 1), "个参数是", arg)
	}
}

/**

type Person struct {
	Name string
	Age  int
}

type PersonV2 struct {
	//如果需要在转换成json的时候属性名为小写需要打个``标签
	Name        string `json:"name"`
	Age         int    `json:"age"`
	UserAddress string `json:"user_address"`
}

func jsonTest() {
	println("-------序列化测试:结构体转json-------")
	person := Person{Name: "szc", Age: 23}
	json_bytes, error_ := json.Marshal(&person)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}
	fmt.Printf("person结构体转json:%v \n", string(json_bytes))

	person_2 := PersonV2{Name: "szc", Age: 23, UserAddress: "北京朝阳区"}
	json_bytes, error_ = json.Marshal(&person_2)
	if error_ != nil {
		fmt.Println("Json error:", error_)
		return
	}
	fmt.Printf("person_2结构体转json:%v \n", string(json_bytes))

	println("-------反序列化测试:json转结构体-------")
	jsonStr_1 := "{\"Name\":\"szc\",\"Age\":23}"
	jsonStr_2 := "{\"name\":\"szc\",\"age\":23,\"user_address\":\"北京朝阳区\"} "
	var person1 Person
	var person2 PersonV2

	err_1 := json.Unmarshal([]byte(jsonStr_1), &person1)
	err_2 := json.Unmarshal([]byte(jsonStr_2), &person2)
	fmt.Printf("json转Person结构体:%v \n", person1)
	fmt.Printf("json转PersonV2结构体:%v \n", person2)
	if err_1 != nil || err_2 != nil {
		fmt.Println("Json error:", err_1)
		fmt.Println("Json error:", err_2)
		return
	}
}

func fileIsExit() {
	//文件路径
	filePath := "/Users/dasouche/go/src/MyTest/file_dir/MyFile"
	_, err := os.Stat(filePath)
	var isExit bool
	if err == nil {
		isExit = true
	}
	if os.IsNotExist(err) {
		isExit = false
	}
	fmt.Printf("文件是否存在:%v \n", isExit)
}

func writeFile() {
	//需要创建文件的路径
	filePath := "/Users/dasouche/go/src/MyTest/file_dir/CreateFileTest.txt"
	//参数1:文件路径  参数2:读模式或者创建文件模式 参数3:赋予文件777权限，linux系统777代表所有用户可读写
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("New content" + fmt.Sprintf("%d", i) + "\n") // 写入一行数据
	}

	writer.Flush() // 把缓存数据刷入文件中

	file.Close()
}

func readFile() {
	filePath := "/Users/dasouche/go/src/MyTest/file_dir/MyFile"

	//Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
	file, err := os.Open(filePath)
	fmt.Printf("%v \n", err)
	if nil != err { //如果文件不存在则会报错:open file error = open xxx/xxx.txt: The system cannot find the file specified.
		fmt.Println("文件读取失败")
		//结束运行
		return
	}

	//可以file结构体里存放着一个指针
	fmt.Println("file = ", *file)

	println("-----------读取文件内容方式1-------------")
	//方式1:使用buffer.ReadLine()
	//将读取的文件放入缓冲区， 注意⚠️ :bufio.NewReader(rd io.Reader) 函数内部调用了 NewReaderSize(rd, defaultBufSize)，而这个defaultBufSize的值就是4096。
	//br := bufio.NewReader(file) //建议使用下面自定义大小的缓冲区
	buffer := bufio.NewReaderSize(file, 10240)
	var resultBuffer []byte
	for {
		line, prefix, err := buffer.ReadLine()
		fmt.Printf("读取一行内容:%c , prefix:%v, err:%v \n", line, prefix, err)
		if err == io.EOF { // 读到文件尾会返回一个EOF异常
			break
		}

		// 追加到自定义缓冲区内
		resultBuffer = append(resultBuffer, line...)
		// 如果prefix为真，则代表该行还有尚未读取完的数据，跳过后续具体操作，继续读取完该行剩余内容
		if prefix {
			continue
		}
		str := string(resultBuffer)
		fmt.Printf("--------------------\n")
		fmt.Println("len(buf) = ", len(resultBuffer))
		fmt.Println("len(str) = ", len(str))
		fmt.Println(str)
		fmt.Printf("--------------------\n\n")
		// 清空切片
		resultBuffer = append(resultBuffer[:0], resultBuffer[len(resultBuffer):]...)
	}

	println("-----------读取文件内容方式2-------------")
	//方式2:  file.Read
	file, err = os.Open(filePath)
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF { // 读到文件尾会返回一个EOF异常
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

	println("-----------读取文件内容方式3-------------")
	//方式3： reader.ReadString
	file, err = os.Open(filePath)
	reader := bufio.NewReaderSize(file, 10240)
	for {
		str, err := reader.ReadString('\n') // 一次读取一行
		if err == nil {
			fmt.Print(str) // reader会把分隔符\n读进去，所以不用Println
		} else if err == io.EOF { // 读到文件尾会返回一个EOF异常
			fmt.Println("文件读取完毕")
			break
		} else {
			fmt.Println("read error: ", err)
		}
	}

	println("-----------读取文件内容方式4-------------")
	//方式4：  ioutil.ReadAll、io.ReadAll(file)
	file, err = os.Open(filePath)
	// return 之前记得关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}

	//context, _ := ioutil.ReadAll(file) //此方式不建议使用了
	context, err := io.ReadAll(file)
	fmt.Println(string(context))

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file error = ", err)
	}
}

func instanceOfTest() {
	var str interface{}
	str = "测试"
	str = str.(string)
	fmt.Println(str)
	//str = str.(int) //类型对不上会抛panic: interface conversion: interface {} is string, not int
	//基于上述问题，可以按照下述写法来写
	str, isInt := str.(int) //返回两个，一个是原值，一个是:是否是括号里的类型[即str这个变量是不是int类型]
	println(isInt)
	if isInt {
		fmt.Printf("类型为:%T \n", str)
	} else {
		fmt.Println("str变量不是int类型")
	}
}
*/
