package main

import (
	"fmt"
	"ttserver/config"
	"ttserver/model"
)

func main() {
	config.Init()
	model.InitDb()
	//model.CreateUserTable()
	var user model.User
	model.Find(&user, map[string]interface{}{"username": "liuyang"},[]string{"pass_word","UUID"})

	fmt.Println(user)
}
