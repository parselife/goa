package model

import (
	"crypto/md5"
	"errors"
	"time"
)

// 系统登录用户
type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" xorm:" notnull unique 'name'"`
	Password    [16]byte    `json:"-"`
	DisplayName string    `json:"display_name"`
	Enabled     bool      `json:"enabled"`
	IsAdmin     bool      `json:"is_admin"`
	CreateAt    time.Time `json:"create_at" xorm:"created"`
	UpdateAt    time.Time `json:"update_at" xorm:"updated"`
}

// md5 加密
func Md5Password(userPassword string) [16]byte {
	return md5.Sum([]byte(userPassword))
}

// 验证密码是否正确
func ValidatePassword(userPassword string, hashed [16]byte) (bool, error) {
	if md5.Sum([]byte(userPassword)) == hashed {
		return true, nil
	}
	return false, errors.New("密码不正确")
}
