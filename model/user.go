package model

type User struct {
	//gorm.Model
	ID int `gorm:"primaryKey"`
	Username string `gorm:"username,type:varchar(20);not null"`
	Password string  `gorm:"password,type:varchar(256);not null"`
	UUID string `gorm:"uuid,type:varchar(128);not null"`
	Phone string `gorm:"phone,type:varchar(20);not null"`
	Coin int `gorm:"coin,type:uint;not null"`
}

func CreateUserTable(){
	db := GetDB()
	db.AutoMigrate(&User{})
}

