package main

import (
	"ttserver/config"
	"ttserver/db"
	"ttserver/model"
)

func main() {
	config.Init()
	model.InitDb()
	//model.CreateUserTable()
	var user db.User
	err := model.Add(&user, map[string]interface{}{"username":"wxl", "password":"123"})
	if err != nil{
		panic(err)
	}
	return
	//err := model.Find(&user, map[string]interface{}{"username": "liuyang"},[]string{"password","uuid","coin"})
	//if err != nil{
	//	return
	//}
	//fmt.Println(user)
	//
	//err = model.Update(&user, map[string]interface{}{"username":"liuyang"}, map[string]interface{}{"coin":789})
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(user)


}
