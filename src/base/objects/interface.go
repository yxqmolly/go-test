package main

import (
	"fmt"
	"time"
)

type Phone struct{
	name string
}

func (p *Phone) start(){
	fmt.Printf("手机:%s开始工作", p.name)
}

func (p *Phone) stop(){
	fmt.Printf("手机:%s工作: ", p.name)
}

func (p *Phone) charge(){
	fmt.Printf("手机:%s充电: ", p.name)
	time.Sleep(1000)
	fmt.Printf("手机:%s充电完成: ", p.name)
}




