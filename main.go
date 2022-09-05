package main

import (
	"ttserver/config"
	"ttserver/model/mysql"
)

func main() {
	config.Init()
	mysql.InitDb()
}
