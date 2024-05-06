package dao

import "github.com/reoden/todolist/app/user/repository/db/model"

func migrate() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.User{})
}
