package biz

import (
	"context"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Age       uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	ListUser(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id int64, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) (u *User, err error) {
	u, err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	// 可能创建成功之后还会有什么操作，比如操作 Redis 或者其他的记录
	return
}

func (uc *UserUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUsecase) Update(ctx context.Context, id int64, user *User) (u *User, err error) {
	u, err = uc.repo.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (u *User, err error) {
	u, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) List(ctx context.Context) (us []*User, err error) {
	return uc.repo.ListUser(ctx)
}
