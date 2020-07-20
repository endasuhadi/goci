package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    Db *gorm.DB
    Driver = "mysql"
    User = "user"
    Password = ""
    Dbname = "test"
)