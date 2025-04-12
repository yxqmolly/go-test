package main

import (
	"fmt"
	"go-test/src/base/objects/model"
)

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

