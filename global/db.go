package global

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	if db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/office-go")); err != nil {
		log.Fatalf("create db error : %s", err.Error())
	} else {
		Db = db
	}
}
