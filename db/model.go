package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	//gorm.Model
	ID int `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username"`
	Password string  `gorm:"password" json:"password"`
	UUID string `gorm:"uuid" json:"uuid"`
	Phone string `gorm:"phone" json:"phone"`
	Coin int `gorm:"coin" json:"coin"`
	Status int `gorm:"status" json:"status"`
}


type TableVideo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Id          int64  `gorm:"type:int;not null" json:"id"`
	AuthorId    int64
	Desc        string `gorm:"type:varchar(200)" json:"desc"`
	PlayUrl     string `gorm:"type:varchar(200)" json:"play_url"`
	CoverUrl    string `gorm:"type:varchar(200)" json:"cover_url"`
	PublishTime time.Time
}

type Like struct {
	Id      int64 //自增主键
	UserId  int64 //点赞用户id
	VideoId int64 //视频id
	Cancel  int8  //是否点赞，0为点赞，1为取消赞
}

// Follow 用户关系结构，对应用户关系表。
type Follow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
}