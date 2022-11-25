package infra

import (
	"database/sql"
	"fmt"
	"os"
	"todo_gorm/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	sqlDB, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+database+"?parseTime=true&loc=Local")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	return sqlDB
}

// DBの初期化
func DBInit() *gorm.DB {
	sqlDB := connectDB()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.AutoMigrate(&domain.Todo{})
	// defer sqlDB.Close()
	return db
}

// DBの作成処理
func DBCreate(todo domain.Todo) {
	sqlDB := connectDB()
	defer sqlDB.Close()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.Create(&todo)
}

// DBの読込処理
func DBRead(id ...int) []domain.Todo {
	sqlDB := connectDB()
	defer sqlDB.Close()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

// DBの更新処理
func DBUpdate(id int, text string, status domain.Status, deadline int) domain.Todo {
	sqlDB := connectDB()
	defer sqlDB.Close()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
	return todo
}

// DBの削除処理
func DBDelete(id int) {
	sqlDB := connectDB()
	defer sqlDB.Close()
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
