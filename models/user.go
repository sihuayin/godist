package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id                     int       `orm:"column(id);auto"`
	Username               string    `orm:"column(username);size(255)"`
	IsEmailVerified        int8      `orm:"column(is_email_verified)"`
	AuthKey                string    `orm:"column(auth_key);size(32)"`
	PasswordHash           string    `orm:"column(password_hash);size(255)"`
	PasswordResetToken     string    `orm:"column(password_reset_token);size(255);null"`
	EmailConfirmationToken string    `orm:"column(email_confirmation_token);size(255);null"`
	Email                  string    `orm:"column(email);size(255)"`
	Avatar                 string    `orm:"column(avatar);size(100);null"`
	Role                   int16     `orm:"column(role)"`
	Status                 int16     `orm:"column(status)"`
	CreatedAt              time.Time `orm:"column(created_at);type(datetime)"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime)"`
	Realname               string    `orm:"column(realname);size(32)"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
