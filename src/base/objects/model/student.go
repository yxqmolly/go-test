package model

type Student struct {
	Id int
	Name string
	Score float64
}

//属性或方法的首字母小写只能当前文件使用，大写类似于java中public
func NewStudent(id int, n string, s float64) *Student {
	return &Student{
		Id: id,
		Name : n,
		Score: s,
	}
}