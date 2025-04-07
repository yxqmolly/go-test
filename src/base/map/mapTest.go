package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string){
	if users[name] != nil {
		users[name]["pwd"] = "12121"
	}else {
		users[name] = make(map[string]string, 2)
		users[name]["name"] = name
		users[name]["pwd"] = "123456"
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = map[string]string{
		"name": "smith",
		"pwd": "123456",
	}
	modifyUser(users, "nora")
	for key,val := range users{
		fmt.Println("key= ", key, "val=", val)
	}
	
}