package models

import (
	"os"
	"fmt"
	// "time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	_ "github.com/lib/pq"
)

// Clientから受け取るUser情報の構造体
type Login struct {
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag"`
}

// ######################################
// Table構造を反映させるための構造体
// ######################################

// Tableへ格納するために使用する構造体
type User struct {
	// 以下を宣言することで、ID, CreatedAt, UpdatedAtを自動的に作成してくれる
	gorm.Model
	ID uint `gorm:"primary_key"`
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password []byte `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag" binding:"required"`
	Schedules []Schedule `gorm:"many2many:user_schedule_rel"`
}

type Schedule struct {
	gorm.Model
	ID uint `gorm:"primary_key"`
	Title string `form:"Title" json:"Title" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag"`
	Todos []Todo `gorm:"foreignKey:ScheduleId"`
}

// type UserScheduleRel struct {
// 	gorm.Model
// 	ID uint `gorm:"primary_key"`
// 	UserId uint
// 	ScheduleId uint
// }

type Todo struct {
	gorm.Model
	ID uint `gorm:"primary_key"`
	Content string `form:"Content" json:"Content" binding:"required"`
	ScheduleId uint
}

// ######################################
// 共通関数
// ######################################

// DBへの接続関数
func GetDB() (*gorm.DB, error) {
	postgresqlInfo := fmt.Sprintf("host=db port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	db, err := gorm.Open(postgres.Open(postgresqlInfo), &gorm.Config{})

	return db, err
}