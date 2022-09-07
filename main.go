package main

import (
	"ttserver/config"
	"ttserver/db"
)

func main() {
	config.Init()
	db.InitDb()
	//model.CreateUserTable()
	var user db.User
	err := db.Add(&user, map[string]interface{}{"username": "wxl", "password": "123"})
	if err != nil {
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
