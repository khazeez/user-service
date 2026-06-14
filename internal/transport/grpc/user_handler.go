package grpc

import (
	"context"

	pb "github.com/khazeez/user-service/proto"

	"github.com/khazeez/user-service/internal/service"
)

type UserHandler struct {
    pb.UnimplementedUserServiceServer

    service *service.UserService
}

func NewUserHandler(
    service *service.UserService,
) *UserHandler {

    return &UserHandler{
        service: service,
    }
}


func (h *UserHandler) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.UserResponse, error) {

	user, err := h.service.CreateUser(
		req.Name,
		req.Email,
	)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.UserResponse, error) {

	user, err := h.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) ListUsers(
	ctx context.Context,
	req *pb.Empty,
) (*pb.UserListResponse, error) {

	users, err := h.service.ListUsers()
	if err != nil {
		return nil, err
	}

	resp := &pb.UserListResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, &pb.UserResponse{
			Id: u.ID,
			Name: u.Name,
			Email: u.Email,
		})
	}
	return resp, nil
}

func (h *UserHandler) DeleteUser(
	ctx context.Context,
	req *pb.DeleteUserRequest,
) (*pb.Empty, error) {

	err := h.service.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}