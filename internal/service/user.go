package service

import (
	"context"

	pb "subway_intro/api/user/v1"
	"subway_intro/internal/biz"
	"subway_intro/internal/rpc"
)

type UserService struct {
	user *biz.UserUsecase
	rpc  *rpc.RpcClient
	pb.UnimplementedUserServiceServer
}

func NewUserService(user *biz.UserUsecase, r *rpc.RpcClient) *UserService {
	return &UserService{
		user: user,
		rpc:  r,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	u, err := s.user.Create(ctx, &biz.User{
		Name: req.Name,
		Age:  req.Age,
	})
	if err != nil {
		return nil, pb.ErrorUserDomainError("user create error name=%s", req.Name)
	}
	return &pb.CreateUserReply{User: &pb.User{
		Id:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}}, err
}

func (s *UserService) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.user.Delete(ctx, req.Id)
	return &pb.DeleteUserReply{}, err
}

func (s *UserService) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.user.Get(ctx, req.Id)
	if err != nil {
		return nil, pb.ErrorUserNotFound("user not found id=%s", req.Id)
	}
	return &pb.GetUserReply{User: &pb.User{
		Id:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}}, err
}

func (s *UserService) List(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	us, err := s.user.List(ctx)
	reply := &pb.ListUserReply{}
	for _, u := range us {
		reply.UserList = append(reply.UserList, &pb.User{
			Id:   u.ID,
			Name: u.Name,
			Age:  u.Age,
		})
	}
	return reply, err
}

func (s *UserService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	u, err := s.user.Update(ctx, req.Id, &biz.User{
		Name: req.Name,
		Age:  req.Age,
	})
	if err != nil {
		return nil, pb.ErrorUserDomainError("user update error name=%s", req.Name)
	}
	return &pb.UpdateUserReply{User: &pb.User{
		Id:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}}, err
}
