package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ts.com/backend/model"
)

var db *gorm.DB

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(mysql:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志级别为info
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
	}

	db.Create(&model.User{
		Name:     "李白",
		Nickname: "xiaobai",
		Age:      100,
		Password: "123456",
		Gender:   true,
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		var user model.User
		db.First(&user)
		c.JSON(200, gin.H{
			"message": "success",
			"data":    user,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//r.Run(":8081")
}
