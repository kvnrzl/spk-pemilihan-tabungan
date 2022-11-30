package services

import (
	"context"
	"errors"
	"project_spk_pemilihan_tabungan/config"
	model "project_spk_pemilihan_tabungan/models"
	repository "project_spk_pemilihan_tabungan/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminServiceImpl struct {
	adminRepository repository.AdminRepository
	db              *gorm.DB
	validate        *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, DB *gorm.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		adminRepository: adminRepository,
		db:              DB,
		validate:        validate,
	}
}

func (s *AdminServiceImpl) AdminCreate(ctx context.Context, username string, password string) (*model.Admin, error) {
	var admin model.Admin

	if err := s.validate.Var(username, "required"); err != nil {
		return &model.Admin{}, err
	}

	if err := s.validate.Var(password, "required"); err != nil {
		return &model.Admin{}, err
	}

	DB := s.db.Begin()

	if admin, _ := s.adminRepository.FindByUsername(ctx, DB, username); admin.Username != "" {
		DB.Rollback()
		return &model.Admin{}, errors.New("Username already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		DB.Rollback()
		return &model.Admin{}, err
	}

	admin = model.Admin{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  username,
		Password:  string(hashedPassword),
	}

	result, err := s.adminRepository.Create(ctx, s.db, &admin)
	if err != nil {
		DB.Rollback()
		return &model.Admin{}, err
	}

	DB.Commit()

	return result, nil
}

func (s *AdminServiceImpl) AdminUpdate(ctx context.Context, username string, password string) (*model.Admin, error) {
	return nil, nil
}
func (s *AdminServiceImpl) AdminDelete(ctx context.Context, id int) error {
	return nil
}

func (s *AdminServiceImpl) AdminLogin(ctx context.Context, username string, password string) (string, error) {
	if err := s.validate.Var(username, "required"); err != nil {
		return "", errors.New("Username is required")
	}

	if err := s.validate.Var(password, "required"); err != nil {
		return "", errors.New("Password is required")
	}

	DB := s.db.Begin()

	admin, err := s.adminRepository.FindByUsername(ctx, DB, username)
	if err != nil {
		return "", err
	}

	DB.Commit()

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", errors.New("Username or password is wrong")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: time.Now().Add(config.JWT_EXPIRE_DURATION).Unix(),
	})
	token, err := claims.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AdminServiceImpl) AdminLogout(ctx context.Context, token string) error {
	return nil
}
