package model

import "ttserver/model/mysql"

type User struct {
	ID int `gorm:"ID,primary_key" json:"id"`
	UserName string `gorm:"username, type:varchar(20);not null" json:"user_name"`
	PassWord string  `gorm:"password, type:varchar(256);not null" json:"pass_word"`
	UUID string `gorm:"uuid,type :varchar(128);not null" json:"uuid"`
	Phone string `gorm:"phone, type:varchar(20);not null" json:"phone"`
	Coin int `gorm:"coin, type:uint; not null" json:"coin"`
}

func CreateUserTable(){
	db := mysql.GetDB()
	db.AutoMigrate(&User{})
}

func Find(TableName string, key map[string]interface{}, fields []interface{} ){
	db  := mysql.GetDB()
	var user User
	result := db.Where(key).Find(&user)
	if result.Error != nil{
		return user
	}
	return user

}

func (u *User)Insert()

//根据key
func (u *User)Update(key interface{}, )