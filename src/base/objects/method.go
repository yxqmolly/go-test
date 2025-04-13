package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-test/src/base/objects/model"
	"io"
	"os"
)
 var FILE_PATH_ABSOLUTE_RESOURCE = "/Users/yxq/Desktop/vscode-workspace/go/src/go-test/src/resource/linshi/"
 var FILE_PATH_RELATIVE_RESOURCE = "../../../src/resource/linshi/"


type Person struct {
	Id   int
	Name string
}

func Test1() {
	var p1 Person
	p1.Id = 1
	p1.Name = "nora"
	fmt.Println(p1)
	var p2 *Person = &p1
	fmt.Println(p2)
	fmt.Println((*p2).Id)
	fmt.Println(p2.Id)
}

/**
 * 结构体的定义,首字母大写
 */
func Test2() {
	var stu = model.NewStudent(1, "nora", 100.1)
	fmt.Println(stu)
	var stu2 = &model.Student{
		Id:    2,
		Name:  "nora2",
		Score: 100.2,
	}
	fmt.Println(stu2)
	fmt.Println(stu2.Score)
}

/**
 * 结构体的定义,首字母小写
 */
func Test3() {
	var stu = model.NewStudent2(1, "nora", 100.1)
	fmt.Println(stu)
}

func (p Person) Test4(n int) {
	res := 0
	for i := 1; i < n; i++ {
		res += 1
	}
	fmt.Println(p.Name, ",计算结果：", res)
}

/**
* toString()
*/
func (p Person) toString() string {
	return fmt.Sprintf("Person{Id=%d, Name=%s}", p.Id, p.Name)
}



func PointMethod1(){
	r1 := model.Rect2{
		Leftup: &model.Point{1,200},
		RightDown: &model.Point{40,500},
	}
	fmt.Printf("%d,%d,%d,%d",&r1.Leftup.X, &r1.Leftup.Y, 
	&r1.RightDown.X, &r1.RightDown.Y)
}

func TypeJudge(items... interface{}){
	for index, x := range items {
		switch x.(type){
		case bool:
			fmt.Printf("第%d个参数是bool类型，值是%v\n", index, x)	
		default:
			fmt.Printf("第%d个参数是%T类型，值是%v\n", index, x, x)
		}
	}
}

func  Deposit(account *model.Account, accountNo string, pwd string, money float64) {
	if account == nil  {
		fmt.Println("账号不存在")
		return
	}
	if accountNo != account.AccountNo || pwd != account.Pwd {
		fmt.Println("账号或密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("存款金额必须大于0")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func TestOS(){
	fmt.Println("命令行的参数:", len(os.Args))
	for i,v :=range os.Args{
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}

func OpenFile1(fileName string){
	file, err := os.Open(FILE_PATH_RELATIVE_RESOURCE + fileName)
	if  err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if  err := scanner.Err(); err != nil {
		fmt.Println("read file error", err)
	}
}

func OpenFile2(fileName string){
	file, _ := os.Open(FILE_PATH_RELATIVE_RESOURCE + fileName)
	//defer xxx注册一个函数调用，在方法执行正常/异常结束使用
	defer file.Close()
	content,err := io.ReadAll(file)
	if  err != nil {
		fmt.Println("open file error", err)
		return
	}
	fmt.Println(string(content))
}
/** 
 * 文件不存在就创建文件、存在就追加内容
 * 
 */
func WriteAndReadFile(fileName string){
	filePath := FILE_PATH_RELATIVE_RESOURCE + fileName
	//os.O_WRONLY|os.O_APPEND //不存在就不写
	//os.O_WRONLY|os.O_CREATE //不存在则创建,存在就覆盖
	// file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// file, err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if  err!= nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()
	str := "===test WriteAndReadFile===\n"
	write := bufio.NewWriter(file)
	for i :=0; i<5; i++ {
		write.WriteString(str)
	}
	write.WriteString("===test WriteAndReadFile===\n")
	write.Flush()
	OpenFile1(fileName)
}

func ReadLion(){
		//定义几个变量，用于接收命令行的参数值
		var user string
		var pwd string
		var host string
		var port int
		
		//&user 就是接收用户命令行中输入的 -u 后面的参数值
		//"u" ,就是 -u 指定参数
		//"" , 默认值
		//"用户名,默认为空" 说明
		flag.StringVar(&user, "u", "", "用户名,默认为空")
		flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
		flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
		flag.IntVar(&port, "port", 3306, "端口号,默认为3306")
		//这里有一个非常重要的操作,转换， 必须调用该方法
		flag.Parse()
	
		//输出结果
		fmt.Printf("user=%v pwd=%v host=%v port=%v", 
			user, pwd, host, port)
}
