package model

type User struct {
	//gorm.Model
	ID int `gorm:"primaryKey" json:"id"`
	Username string `gorm:"username,type:varchar(20);not null" json:"username"`
	Password string  `gorm:"password,type:varchar(256);not null" json:"password"`
	UUID string `gorm:"uuid,type:varchar(128);not null" json:"uuid"`
	Phone string `gorm:"phone,type:varchar(20);not null" json:"phone"`
	Coin int `gorm:"coin,type:uint;not null" json:"coin"`
	Status int `gorm:"status,type:uint;not null" json:"status"`
}

func CreateUserTable(){
	db := GetDB()
	db.AutoMigrate(&User{})
}

