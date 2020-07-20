package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"app/conf"
	"app/routes"
	"app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.New()
	var err error
	connectString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=False&loc=Local", conf.User, conf.Password, conf.Dbname)
	conf.Db, err = gorm.Open(conf.Driver, connectString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conf.Db.Close()
	conf.Db.SingularTable(true)
	conf.Db.AutoMigrate(&model.Siswa{})
	routes.Routes(r)
	r.Run(":8080")
}
