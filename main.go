package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gender uint8
type Role uint32

const (
	Male Gender = iota
	Female
	AnonymousRole Role = 0
	NormalRole    Role = 1
	ManagerRole   Role = 2
)

type User struct {
	ID        uint32 `gorm:"primaryKey;type:INT UNSIGNED NOT NULL AUTO_INCREMENT" validate:"required,gt=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:CHAR(16);not null;comment:用户名" validate:"required,min=1,trimstr"`
	Avatar    string         `gorm:"type:VARCHAR(500);not null;comment:用户头像URL" validate:"required,url"`
	Gender    Gender         `gorm:"comment:性别" validate:"oneof=0 1"`
	// 用户角色,添加新角色 时要在validate的oneof中对应添加
	Role `gorm:"not null;comment:角色" validate:"required,oneof=0 1 2"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:passwd1234@tcp(127.0.0.1)/user?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	u := &User{
		Name:   "test",
		Avatar: "https://www.google.com",
		Gender: Female,
		Role:   AnonymousRole,
	}
	db.AutoMigrate(u)
	result := db.Create(u)
	fmt.Println(u)
	fmt.Println(result.Error)
}
