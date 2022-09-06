package model

import (
	"time"

	"gorm.io/gorm"
)

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

// GetVideosByAuthorId
// 根据作者的id来查询对应数据库数据，并返回Video切片
func GetVideosByAuthorId(authorId int64) ([]TableVideo, error) {
	//建立结果集接收
	var data []TableVideo
	//初始化db
	//Init()
	result := db.Where(&TableVideo{AuthorId: authorId}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// GetVideoIdsByAuthorId
// 通过作者id来查询发布的视频id切片集合
func GetVideoIdsByAuthorId(authorId int64) ([]int64, error) {
	var id []int64
	//通过pluck来获得单独的切片
	result := db.Model(&TableVideo{}).Where("author_id", authorId).Pluck("id", &id)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return id, nil
}

// GetVideoByVideoId
// 依据VideoId来获得视频信息
func GetVideoByVideoId(videoId int64) (TableVideo, error) {
	var tableVideo TableVideo
	tableVideo.Id = videoId
	//Init()
	result := db.First(&tableVideo)
	if result.Error != nil {
		return tableVideo, result.Error
	}
	return tableVideo, nil

}
