
func modifyUser(users map[string]map[string]string, name string){
	if users[name] != nil {
		users[name]["pwd"] = "12121"
	}else {
		users[name] = map(map[string])
	}
}

func ()  {
	
}