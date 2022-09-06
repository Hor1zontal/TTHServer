package main

import (
	"ttserver/config"
	"ttserver/model"
)

func main() {
	config.Init()
	mysql.InitDb()
	//model.CreateUserTable()
}
