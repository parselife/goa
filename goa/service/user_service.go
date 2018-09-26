package service

import (
	"errors"
	"goa/goa/model"
	"goa/goa/repo"
)

type UserService interface {
	GetAll() []model.User
	GetByID(id int64) (model.User, bool)
	GetByOrgan(organ int64) ([]model.User, bool)
	GetByUsername(username string) (model.User, bool)
	DeleteByID(id int64) bool

	UpdatePassword(id int64, oldPassword string, newPassword string) (model.User, error)
	Save(user model.User) (model.User, error)
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo repo.UserRepository
}

// 获取所有用户
func (s *userService) GetAll() []model.User {
	return s.repo.FindAll()
}

// 获取单个用户
func (s *userService) GetByID(id int64) (model.User, bool) {
	return s.repo.FindOne(id)
}

func (s *userService) GetByOrgan(organ int64) ([]model.User, bool) {
	users, err := s.repo.FindByOrgan(organ)
	if err != nil {
		return nil, false
	}
	return users, true
}

func (s *userService) GetByUsername(username string) (model.User, bool) {

	return s.repo.FindByUserName(username)
}

func (s *userService) DeleteByID(id int64) bool {

	return s.repo.DeleteOne(id)
}

func (s *userService) UpdatePassword(id int64, oldPassword string, newPassword string) (model.User, error) {
	user, found := s.repo.FindOne(id)
	if !found {
		return user, errors.New("用户不存在")
	}
	if ok, _ := model.ValidatePassword(oldPassword, user.Password); ok == false {
		return user, errors.New("用户密码输入错误")
	}
	user.Password = model.Md5Password(newPassword)
	return s.repo.Save(user)
}

func (s *userService) Save(user model.User) (model.User, error) {
	return s.repo.Save(user)
}
