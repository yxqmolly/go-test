package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "12121"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["name"] = name
		users[name]["pwd"] = "123456"
	}
}

func main() {
	test3()

}

/**
 *map的value是map类型时，不能使用make创建，必须使用make(map[string]map[string]string)
 */
func test1() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = map[string]string{
		"name": "smith",
		"pwd":  "123456",
	}
	modifyUser(users, "nora")
	for key, val := range users {
		fmt.Println("key= ", key, "val=", val)
	}
}
/** 
 *map的增删改查
 */
func test2() {
	//增加
	cites := make(map[string]string)
	cites["address1"] = "beijing"
	cites["address2"] = "shanghai"
	cites["address3"] = "shenzhen"
	fmt.Println(cites)
	//修改
	cites["address1"] = "ningde"
	fmt.Println(cites)
	//判断是否存在
	if val ,exists:= cites["address11"]; exists{
		fmt.Printf("beijing:%s", val)
	}
	//删除
	delete(cites, "address21")
	cites =make(map[string]string)
	fmt.Println(cites)
}
func test3(){
	var testMap map[string]string
	testMap = make(map[string]string,10)
	fmt.Println(testMap)
	for i:=0; i<120; i++{
		key := fmt.Sprintf("stu%d", i)
		value := fmt.Sprintf("stu%d", i)
		testMap[key] = value
	}
	fmt.Println(testMap)
	fmt.Println("testMap的长度为：%d", len(testMap))
}
