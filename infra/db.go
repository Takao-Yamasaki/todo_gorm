package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todo_gorm/domain"
)

// DB初期化
func dbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	db.AutoMigrate(&domain.Todo{})
	return db
}
