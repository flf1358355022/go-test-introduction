package main

import (
	"fmt"
	"reflect"
)

var number int

type User struct {
	Name, address string
	age           int `json:"Age"`
}

// 新增-设置名称方法
func (s *User) SetName(name string) {
	s.Name = name
	fmt.Printf("有参数方法 通过反射进行调用:%d \n", s)
}

// 新增-打印信息方法
func (s User) PrintStudent() {
	fmt.Printf("无参数方法 通过反射进行调用:%v\n", s)
}

func main() {
	println("--------基本数据类型反射--------")
	number = 100
	//获取反射类型
	reflectType := reflect.TypeOf(number)
	fmt.Println("reflectType = ", reflectType)             // int
	fmt.Println("reflectType name = ", reflectType.Name()) // int

	// 获取属性值
	reflectValue := reflect.ValueOf(number)
	fmt.Printf("reflectValue = %v,reflectValue type = %T\n", reflectValue, reflectValue) // 100, reflect.Value

	n1 := 100 + reflectValue.Int() // 获取反射值持有的整型值
	fmt.Println("n1 = ", n1)

	iV := reflectValue.Interface() // 反射值转换成空接口
	num, ok := iV.(int)            // 类型断言
	fmt.Println("num = ", num, ok)

	//注意⚠️:type和kind有时候有时候是一样的，有时候是不一样的(基本类型一样，结构体不一样)
	fmt.Printf("----%v, %v, %v \n", reflectType.Kind(), reflectValue.Type(), reflectValue.Kind())

	//获取类别
	k := reflectValue.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("number is int\n")
	case reflect.String:
		fmt.Printf("number is string\n")
	}

	println("--------结构体反射--------")
	user := User{Name: "无名", address: "山洞", age: 18}
	reflectType_2 := reflect.TypeOf(user)
	reflectvalue_2 := reflect.ValueOf(user)

	iV_2 := reflectValue.Interface() // 反射值转换成空接口
	fmt.Println("reflectType_2 = ", reflectType_2)
	fmt.Printf("reflectvalue_2 = %v,reflectvalue_2 type = %T\n", reflectvalue_2, reflectvalue_2)
	fmt.Printf("iV_2 value=%d, type=%T \n ", iV_2, iV_2)

	//获取字段数量
	numFieldCount := reflectvalue_2.NumField()
	fmt.Printf("获取到结构体字段数量:%d \n", numFieldCount)
	for i := 0; i < numFieldCount; i++ {
		field := reflectvalue_2.Field(i)
		fmt.Printf("第 %d 个字段值 = %v, 类别 = %v \n", i+1, field, field.Kind()) // 获取字段值
	}

	//获取方法数量
	numMethodCount := reflectvalue_2.NumMethod()
	fmt.Printf("获取到结构体方法数量:%d \n", numMethodCount)
	for i := 0; i < numMethodCount; i++ {
		method := reflectvalue_2.Method(i)
		fmt.Printf("第 %d 个方法地址 = %v, 类别 = %v \n", i+1, method, method.Kind()) // 获取方法相关信息
	}

	//通过reflect.Value获取对应的方法并调用
	m1 := reflectvalue_2.MethodByName("PrintStudent")
	var args []reflect.Value
	m1.Call(args)

	//修改结构体字段属性,方式1
	user_2 := User{Name: "无名-2", address: "山洞", age: 18}
	reflectvalue_3 := reflect.ValueOf(&user_2)
	m2 := reflectvalue_3.MethodByName("SetName")
	var args2 []reflect.Value
	name := "stu01"
	nameVal := reflect.ValueOf(name)
	args2 = append(args2, nameVal)
	m2.Call(args2)
	//修改结构体字段属性,方式2
	fmt.Printf("修改字段前属性:%s \n", user_2)
	//根据字段下标修改，注意⚠️:在struct中的属性，严格区分首字母大小写，大写为公有属性，外面可以访问到，小写为私有，外面访问不到。
	//reflectvalue_3.Elem().Field(2).SetString("小利") //❌
	reflectvalue_3.Elem().Field(0).SetString("小利")
	fmt.Printf("修改字段后属性:%s \n", user_2)
	//根据字段名修改
	//reflectvalue_3.Elem().FieldByName("age").SetInt(99) //❌
	reflectvalue_3.Elem().FieldByName("Name").SetString("---")
	fmt.Printf("修改字段后属性:%s \n", user_2)

	//获取字段中的tag信息 	`json:"Age"`
	numFieldCount = reflectvalue_2.NumField()
	for i := 0; i < numFieldCount; i++ {
		structField := reflectvalue_3.Type().Elem().Field(i)
		fmt.Printf("获取字段结构信息, 字段名称:%v, 字段类型:%v, 字段位置:%v, 字段包路径:%v, 字段tag:%v  \n",
			structField.Name, structField.Type, structField.Index, structField.PkgPath, structField.Tag)
	}

}

/**
func atomicTest() {
	waitGroup := sync.WaitGroup{}

	//两个协程对同一个变量进行加减，如果在没有加锁的情况下，使用atomic能够确保任一时刻只有一个goroutine对变量进行操作
	println("----------atomic.Addxxx测试加法-------------")
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			atomic.AddInt64(&number, 1)
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			atomic.AddInt64(&number, 1)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomic.Addxxx测试加法执行结束")

	println("----------atomic.Addxxx测试加减法-------------")
	number = 0
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			atomic.AddInt64(&number, 1)
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			atomic.AddInt64(&number, -1)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomic.Addxxx测试加减法执行结束")

	println("----------atomic.Loadxxx测试加减法-------------")
	//载入操作能够保证原子的读变量的值，当读取的时候，任何其他CPU操作都无法对该变量进行读写，其实现机制受到底层硬件的支持。
	number = 0
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			loadInt64 := atomic.LoadInt64(&number)
			loadInt64++
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			loadInt64 := atomic.LoadInt64(&number)
			loadInt64--
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomicLoadxxx测试加减法执行结束")
	println("----------atomic.Loadxxx测试加减法-------------")
	//载入操作能够保证原子的读变量的值，当读取的时候，任何其他CPU操作都无法对该变量进行读写，其实现机制受到底层硬件的支持。
	number = 0
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			loadInt64 := atomic.LoadInt64(&number)
			loadInt64++
		}
		fmt.Println("加逻辑结束, number:", number)
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 100*100*100*100; i++ {
			//原子操作
			loadInt64 := atomic.LoadInt64(&number)
			loadInt64--
		}
		fmt.Println("减逻辑结束, number:", number)
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomicLoadxxx测试加减法执行结束")

	println("----------atomic.CompareAndSwapxxx测试比较和替换-------------")
	//CAS操作是先比较变量的值是否等价于给定的值，如果是才进行替换
	number = 0
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 200; i++ {
			if number <= 100 {
				number++
			}
			//原子操作
			atomic.CompareAndSwapInt64(&number, 100, 999) //只有number满足100的时候才能将number替换成999
		}
		fmt.Println("加逻辑结束, number:", number)
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomic.CompareAndSwapxxx测试加减法执行结束")

	println("----------atomic.Swapxxx测试替换-------------")
	//Swap不管变量的旧值是否被改变，直接赋予新值然后返回背替换的值。
	number = 0
	waitGroup.Add(1)
	go func() {

		//原子操作
		swapInt64 := atomic.SwapInt64(&number, 999) //只有number满足100的时候才能将number替换成999
		fmt.Println("加逻辑结束, number:", number)
		fmt.Println("替换后返回的值, swapInt64:", swapInt64)
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("number:", number)
	println("atomic.Swapxxx测试替换执行结束")

	println("----------atomic.Storexxx测试替换-------------")
	//写操作，直接赋值
	number = 0
	atomic.StoreInt64(&number, 1000)

	fmt.Println("number:", number)
	println("atomic.Storexxx测试替换执行结束")
}


func tickerTest()  {
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker 1") //每秒执行一次
		counter++
		if counter > 5 {
			break
		}
	}
	ticker.Stop() //停止
}


func timerTest() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)

	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	//如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后")

	time.Sleep(time.Second * 2)
	fmt.Println("再一次2s后")

	<-time.After(time.Second * 2) //time.After函数的返回值是chan Time
	fmt.Println("再再一次2s后")

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer 3 expired")
	}()

	stop := timer3.Stop() //停止定时器
	////阻止timer事件发生，当该函数执行后，timer计时器停止，相应的事件不再执行
	if stop {
		fmt.Println("Timer 3 stopped")
	}

	fmt.Println("before")
	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
	timer4.Reset(time.Second * 1)            //重新设置时间,即修改NewTimer的时间
	<-timer4.C
	fmt.Println("after")

}

func concurrencyControl() {

	listconlimit := make(chan bool, 10) // 新建长度为10的管道
	wg := &sync.WaitGroup{}
	for n := 0; n <= 50; n++ { // 50
		listconlimit <- true // 管道写入，缓冲为10，写满10就阻塞
		fmt.Printf("当前管道长度:%v \n", len(listconlimit))
		wg.Add(1)
		go func(n int, group *sync.WaitGroup) {
			defer func() {
				data_info := <-listconlimit
				fmt.Printf("读取管道数据:%v, 当前管道长度:%v \n", data_info, len(listconlimit))
				group.Done()
			}() //释放管道资源

			time.Sleep(time.Second) // 模拟耗时操作
			//逻辑执行完上面defer释放资源将协程标记释放
		}(n, wg)

	}

	wg.Wait()
	fmt.Println("ok")
}



func flagTimeOutProcess()  {


	var syncWg sync.WaitGroup

	syncWg.Add(1) //协程计数器加1
	go func() {
		for {
			fmt.Println("协程阻塞中")
			time.Sleep(time.Second)
		}

		syncWg.Done() //上面是循环，所以在这里无法执行
	}()

	go func() {
		//此协程在3秒后执行协程计数器-1操作
		time.Sleep(time.Second * 3)
		syncWg.Done() //协程计数器减1
	}()

	//等待协程执行结束
	syncWg.Wait()

	fmt.Println("主线程执行结束")

}

func selectAndTimeAfter() {
	stopFlagChannel := make(chan string)

	go func() {
		for {
			fmt.Printf("block process,current Second:%v \n", time.Now().Second())
			time.Sleep(time.Second)
		}
		stopFlagChannel <- "processing..."
	}()

	select {
	//第一个case里阻塞的时间只有比第二个case阻塞的时间长的时候, 才能执行第二个case
	case res := <-stopFlagChannel:
		fmt.Println(res)
	case <-time.After(time.Second * 5):
		fmt.Printf("timeout control... stop,current Second:%v \n", time.Now().Second())
	}

	fmt.Println("主线程执行结束")
}

func blockTest() {
	var syncWg sync.WaitGroup

	stopFlagChannel := make(chan string)

	syncWg.Add(1)
	go func() {
		for {
			fmt.Println("协程阻塞中")
			time.Sleep(time.Second)
			//stopFlagChannel <- "processing..."
		}
		syncWg.Done()
	}()

	select {
	//第一个case里阻塞的时间只有比第二个case阻塞的时间长的时候, 才能执行第二个case
	case res := <-stopFlagChannel:
		fmt.Println(res)
	case <-time.After(time.Second * 5):
		//syncWg.Done()
		fmt.Println("timeout control... stop")
	}

	//等待协程执行结束
	syncWg.Wait()

	fmt.Println("主线程执行结束")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //写入数据成功时执行此case
			fmt.Printf("写入管道数据:%v \n", x)
			x, y = y, y+1
		case <-quit: //读数据成功时执行此case
			fmt.Println("quit")
			return //读取完了直接结束
		default:
			// 如果以上都没有符合条件，那么进入default处理流程
			//do nothing
		}
	}
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("读取管道数据:%v \n", <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}


func unlockNumTest() {

	var syncWG sync.WaitGroup

	var maxNum int64
	for i := 0; i <= 100; i++ {
		syncWG.Add(1) //执行协程+1，
		go func() {   // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
			for j := 0; j <= 100; j++ {
				num++
				fmt.Printf("当前变量值:%v \n", num)
				if num > maxNum {
					maxNum = num
				}
			}
			syncWG.Done() //执行协程-1,可认为当前协程执行结束，逻辑执行完毕
		}()
	}

	for i := 0; i <= 100; i++ {
		syncWG.Add(1) //执行协程+1，
		go func() {   // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
			for j := 0; j <= 100; j++ {
				num++
				fmt.Printf("当前变量值:%v \n", num)
				if num > maxNum {
					maxNum = num
				}
			}
			syncWG.Done() //执行协程-1,可认为当前协程执行结束，逻辑执行完毕
		}()
	}

	syncWG.Wait() //阻塞 直到协程组内协程数为0时往下执行

	fmt.Println("主线程执行结束,maxNum:", maxNum)
	fmt.Println("主线程执行结束,变量num:", num)
}

func lockTest() {

	var syncLock sync.Mutex
	var syncWG sync.WaitGroup

	var maxNum int64
	for i := 0; i <= 100; i++ {
		syncWG.Add(1) //执行协程+1，
		go func() {   // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
			syncLock.Lock() // 请求锁
			for j := 0; j <= 100; j++ {
				num++
				fmt.Printf("协程组1-%v执行.......,当前变量值:%v \n", j, num)
				if num > maxNum {
					maxNum = num
				}
			}
			syncLock.Unlock() // 释放锁
			syncWG.Done()     //执行协程-1,可认为当前协程执行结束，逻辑执行完毕
		}()
	}

	for i := 0; i <= 100; i++ {
		syncWG.Add(1) //执行协程+1，
		go func() {   // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
			syncLock.Lock() // 请求锁
			for j := 0; j <= 100; j++ {
				num++
				fmt.Printf("协程组1-%v执行.......,当前变量值:%v \n", j, num)
				if num > maxNum {
					maxNum = num
				}
			}
			syncLock.Unlock() // 释放锁
			syncWG.Done()     //执行协程-1,可认为当前协程执行结束，逻辑执行完毕
		}()
	}

	syncWG.Wait() //阻塞 直到协程组内协程数为0时往下执行

	fmt.Println("主线程执行结束,maxNum:", maxNum)
	fmt.Println("主线程执行结束,变量num:", num)
}

func syncWaitTest() {
	var syncWG sync.WaitGroup

	go func() { // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
		syncWG.Add(1) //执行协程+1
		for i := 0; i < 5; i++ {
			fmt.Println("协程1执行.......", i)
			time.Sleep(time.Second) // 休眠1秒
		}
		syncWG.Done() //执行协程-1,标识当前协程执行结束
	}()

	go func() { // 开启一个协程,main主线程执行玩的时候协程直接结束，主线程并不会等待协程执行结束才结束
		syncWG.Add(1) //执行协程+1，
		for i := 0; i < 5; i++ {
			fmt.Println("协程2执行.......", i)
			time.Sleep(time.Second) // 休眠1秒
		}
		syncWG.Done() //执行协程-1,可认为当前协程执行结束，逻辑执行完毕
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("主线程执行.......", i)
		time.Sleep(time.Second) //休眠1秒
	}

	syncWG.Wait() //阻塞 直到协程组内协程数为0时往下执行

	fmt.Println("主线程执行结束")
}

func testParam() {
	for index, arg := range os.Args {
		fmt.Println("第", (index + 1), "个参数是", arg)
	}
}
*/
