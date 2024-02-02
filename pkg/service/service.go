package service

import (
	"errors"
	"fmt"

	pb "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/pb"
	inter "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/repo/interfaces"
	interr "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/service/interfaces"
	user "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/user/handler"
	userpb "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/user/pb"
)

type AdminService struct {
	AdminRepo inter.AdminRepoInter
	client    userpb.UserServiceClient
}

func NewAdminService(repo inter.AdminRepoInter, client userpb.UserServiceClient) interr.AdminServiceInter {
	return &AdminService{
		AdminRepo: repo,
		client:    client,
	}
}

func (a *AdminService) AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error) {
	admin, err := a.AdminRepo.FindAdmin(p.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != p.Password {
		return nil, errors.New("incorrect password")
	}

	response := pb.AdminResponse{
		Status:  "Success",
		Error:   "",
		Message: "Logged in successfully",
	}
	return &response, nil
}

func (a *AdminService) CreateUserService(p *pb.User) (*pb.AdminResponse, error) {
	result, err := user.CreateUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	fmt.Println("helooooooooooooo")
	response := &pb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}
	return response, nil
}

func (a *AdminService) EditUserService(p *pb.User) (*pb.User, error) {
	result, err := user.EditUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	response := &pb.User{
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
	}
	return response, nil
}

func (a *AdminService) DeleteUserService(p *pb.UserID) (*pb.AdminResponse, error) {
	result, err := user.DeleteUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	respose := &pb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}

	return respose, nil
}

func (a *AdminService) FindUserByEmailService(p *pb.UserRequest) (*pb.User, error) {
	result, err := user.FindUserByEmailHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	response := &pb.User{
		Username: result.Email,
		Email:    result.Email,
		Password: result.Password,
		Role:     result.Role,
	}

	return response, nil
}

func (a *AdminService) FindAllUserService(p *pb.NoParam) (*pb.UserList, error) {
	result, err := user.FindAllUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	var users []*pb.User
	for _, i := range result.Userlist {
		users = append(users, &pb.User{
			Username: i.Username,
			Email:    i.Email,
			Password: i.Password,
			Role:     i.Role,
		})
	}
	respnse := &pb.UserList{
		Users: users,
	}
	return respnse, nil
}

func (a *AdminService) FindUserByIDService(p *pb.UserID) (*pb.User, error) {
	result, err := user.FindUserByIDHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	fmt.Println("2231")

	response := &pb.User{
		Username: result.Email,
		Email:    result.Email,
		Password: result.Password,
		Role:     result.Role,
	}

	return response, nil
}
