package main

import (
	"ttserver/config"
	"ttserver/model"
)

func main() {
	config.Init()
	model.InitDb()
}
