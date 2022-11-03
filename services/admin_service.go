package services

import (
	"context"
	model "project_spk_pemilihan_tabungan/models"
)

type AdminService interface {
	AdminCreate(ctx context.Context, username string, password string) (*model.Admin, error)
	AdminUpdate(ctx context.Context, username string, password string) (*model.Admin, error)
	AdminDelete(ctx context.Context, id int) error
	AdminLogin(ctx context.Context, username string, password string) (string, error)
	AdminLogout(ctx context.Context, token string) error
}
