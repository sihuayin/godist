package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username               string    `gorm:"column:username;size:255"`
	IsEmailVerified        int8      `gorm:"column:is_email_verified`
	AuthKey                string    `gorm:"column:auth_key;size:32"`
	PasswordHash           string    `gorm:"column:password_hash;size:255"`
	PasswordResetToken     string    `gorm:"column:password_reset_token;size:255"`
	EmailConfirmationToken string    `gorm:"column:email_confirmation_token;size:255"`
	Email                  string    `gorm:"column:email;size:255"`
	Avatar                 string    `gorm:"column:avatar;size:100"`
	Role                   int16     `gorm:"column:role`
	Status                 int16     `gorm:"column:status"`
	CreatedAt              time.Time `gorm:"column:created_at"`
	// UpdatedAt              time.Time `gorm:"column:updated_at"`
	Realname string `gorm:"column:realname;size:32"`
}

func (u *User) TableName() string {
	return "user"
}

func FindOneByName(name string) *User {
	var u User
	// fmt.Println(db)
	globalDB.Table("user").Where("username = ?", name).Scan(&u)

	return &u
}

func FindOneByAuthKey(key string) *User {
	var u User
	// fmt.Println(db)
	globalDB.Table("user").Where("auth_key = ?", key).Scan(&u)
	return &u
}
