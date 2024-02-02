package interfaces

import "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/entities"

type AdminRepoInter interface {
	FindAdmin(username string) (*entities.Admin,error)
}