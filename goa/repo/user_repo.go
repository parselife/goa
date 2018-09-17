package repo

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"goa/goa/model"
)

type UserRepository interface {
	FindOne(id int64) (user model.User, found bool)
	FindByUserName(name string) (user model.User, found bool)
	FindAll() (users []model.User)
	//FindPage() (users []model.User)
	Save(user model.User) (updatedUser model.User, err error)
	DeleteOne(id int64) (ok bool)
}

func NewUserRepo(engine *xorm.Engine) UserRepository {
	return &userRepo{engine: engine}
}

type userRepo struct {
	engine *xorm.Engine
}

// find one
func (r *userRepo) FindOne(id int64) (user model.User, found bool) {
	ret := new(model.User)
	has, err := r.engine.Id(id).Get(ret)
	if err != nil {
		fmt.Println(err)
		return *ret, false
	}
	return *ret, has
}

func (r *userRepo) FindByUserName(userName string) (user model.User, found bool) {
	ret := model.User{Name: userName}
	has, err := r.engine.Get(&ret)
	if err != nil {
		fmt.Println(err)
		return ret, false
	}
	return ret, has
}

func (r *userRepo) FindAll() (arr []model.User) {
	users := make([]model.User, 0)
	err := r.engine.Find(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	return users
}

func (r *userRepo) Save(user model.User) (updatedUser model.User, err error) {
	if user.ID > 0 {
		// update
		_, err = r.engine.Where("id = ?", user.ID).Update(&user)
		if err != nil {
			return
		}
		return user, nil
	} else {
		//	insert
		id, err := r.engine.Insert(&user)
		if err != nil {
			fmt.Println(err)
			return user, err
		}
		u, ok := r.FindOne(id)
		if ok {
			return u, nil
		}
		return user, errors.New("save error")
	}
}

func (r *userRepo) DeleteOne(id int64) (ok bool) {
	_, err := r.engine.Delete(&model.User{ID: id})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
