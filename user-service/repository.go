package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/kuzicala/shippy/user-service/proto/user"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	if e := u.db.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func (u *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if e := u.db.First(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func (u *UserRepository) Create(user *pb.User) error {
	if e := u.db.Create(&user).Error; e != nil {
		return e
	}

	return nil
}

func (u *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if e := u.db.First(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}
