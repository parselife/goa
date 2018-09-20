package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"goa/goa/core"
)

// 系统登录用户
type User struct {
	ID          int64     `json:"id"`
	OrganId     int64     `json:"organ, omitempty"`
	Name        string    `json:"name" xorm:" notnull unique 'name'"`
	Password    string    `json:"password"`
	DisplayName string    `json:"displayName"`
	Enabled     bool      `json:"enabled" xorm:"default 1"`
	IsAdmin     bool      `json:"isAdmin" xorm:"default 0"`
	CreateAt    core.Time `json:"createAt" xorm:"created"`
	UpdateAt    core.Time `json:"updateAt" xorm:"updated"`
}

// md5 加密
func Md5Password(userPassword string) string {
	hash := md5.Sum([]byte(userPassword))
	// 将[]byte转成16进制
	return fmt.Sprintf("%x", hash)
}

// 验证密码是否正确
func ValidatePassword(userPassword string, hashed string) (bool, error) {
	if Md5Password(userPassword) == hashed {
		return true, nil
	}
	return false, errors.New("密码不正确")
}