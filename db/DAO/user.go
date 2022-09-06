package DAO

import (
	"ttserver/db"
	"ttserver/utils"
)

type User db.User

func CreateUserTable(){
	database := db.GetDB()
	database.AutoMigrate(&db.User{})
}

//创建出User用户的增删改查
func (u *User) Find(condition map[string]interface{}, fields[]string) User {
	var user User
	err := db.Find(&user, condition, fields)
	if err != nil{
		return User{}
	}
	return user
}
//如果跟新失败返回空用户， 否则返回更新完毕的用户
func (u *User) Update(condition map[string]interface{}, key map[string]interface{}) User {
	var user User
	err := db.Update(&user, condition, key)
	if err != nil{
		return User{}
	}
	newMap := make(map[string]interface{})
	for key,val := range condition{
		newMap[key] = val
	}

	for key,val := range key{
		newMap[key] = val
	}

	return u.Find(newMap, utils.GetTagName(User{}))
}

func (u *User) Delete(condition map[string]interface{})error{
	var user User
	err := db.Delete(&user,condition)
	if err != nil{
		return err
	}
	return nil
}

func (u *User)Add(key map[string]interface{})error{
	var user User
	err := db.Add(&user, key)
	if err != nil{
		return err
	}
	return nil
}

