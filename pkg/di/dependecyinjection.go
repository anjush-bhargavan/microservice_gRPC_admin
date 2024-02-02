package di

import (
	"log"

	"github.com/anjush-bhargavan/microservice_gRPC_admin/config"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/db"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/handler"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/repo"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/server"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/service"
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/user"
)



func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	client, err := user.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong %v",err)
	}

	adminRepo := repo.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo,client)
	adminHandler := handler.NewAdminHandler(adminService)
	server.NewGrpcServer(adminHandler)
}
