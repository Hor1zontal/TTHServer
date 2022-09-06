package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
	"ttserver/config"
	"github.com/mitchellh/mapstructure"
)

var db *gorm.DB
var err error

func InitDb() {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Server.Database.DBUser,
		config.Server.Database.DBPassword,
		config.Server.Database.DBHost,
		config.Server.Database.DBPort,
		config.Server.Database.DBName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	// _ = db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
func GetDB()*gorm.DB{
	return db
}

func Find(TypeName interface{}, condition map[string]interface{}, fields []string )error{
	db  := GetDB()
	if _, ok := condition["status"]; !ok{
		condition["status"] = 1
	}
	result := db.Where(condition).Select(fields).Find(TypeName)
	//fmt.Println(result.RowsAffected)
	//fmt.Println(TypeName)
	if result.Error != nil{
		return err
	}
	return nil
}

func Update(TypeName interface{}, condition map[string]interface{}, key map[string]interface{})error{
	db := GetDB()
	if _, ok := condition["status"]; !ok{
		condition["status"] = 1
	}
	result := db.Model(TypeName).Where(condition).Updates(key)
	//fmt.Println(result.RowsAffected)
	if result.Error != nil{
		return err
	}
	return nil
}

//假删除，让Status 设置为 1
func Delete(TypeName interface{}, condition map[string]interface{})error{

	if _, ok := condition["status"]; !ok{
		condition["status"] = 1
	}
	err := Update(TypeName,condition,map[string]interface{}{"status":1})
	if err != nil{
		return err
	}
	return nil
}


func Add(TypeName interface{}, key map[string]interface{})error{

	err := mapstructure.Decode(key,TypeName)
	if err != nil{
		return err
	}
	db := GetDB()
	result := db.Create(TypeName)
	fmt.Println(result.RowsAffected)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

