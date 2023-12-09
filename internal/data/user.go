package data

import (
	"context"

	"subway_intro/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 初始化Repo
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/task")),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	u, err := ur.data.db.User.
		Create().
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}, err
}

func (ur *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	u, err := ur.data.db.User.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}, nil
}

func (ur *userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	us, err := ur.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, u := range us {
		rv = append(rv, &biz.User{
			ID:   u.ID,
			Name: u.Name,
			Age:  u.Age,
		})
	}
	return rv, nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) (*biz.User, error) {
	u, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	u, err = u.Update().
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
	return &biz.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}, nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, id int64) error {
	return ur.data.db.User.DeleteOneID(id).Exec(ctx)
}
