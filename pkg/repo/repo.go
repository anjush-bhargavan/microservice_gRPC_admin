package repo

import (
	"github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/entities"
	inter "github.com/anjush-bhargavan/microservice_gRPC_admin/pkg/repo/interfaces"
	"gorm.io/gorm"
)


type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInter {
	return &AdminRepo{
		DB:db,
	}
}

func (a *AdminRepo)  FindAdmin(username string) (*entities.Admin,error) {
	var admin entities.Admin
	err := a.DB.Where("username = ?",username).First(&admin).Error
	return &admin,err
}