package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todo_gorm/domain"
)

// DBの初期化
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	db.AutoMigrate(&domain.Todo{})
	return db
}

// DBの作成処理
func DBCreate(todo domain.Todo) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	db.Create(&todo)
} 

//DBの読込処理
func DBRead(id ...int) []domain.Todo {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

//DBの更新処理
func DBUpdate(id int, text string, status domain.Status, deadline int) domain.Todo {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
	return todo
} 

//DBの削除処理
func DBDelete(id int) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/todo_gorm?parseTime=true")
	if err != nil {
		fmt.Errorf("cannot open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}