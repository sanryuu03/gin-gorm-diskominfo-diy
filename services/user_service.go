package services

import (
	"errors"
	"gin-gorm-diskominfo-diy/models/domain"
	"gin-gorm-diskominfo-diy/repositories"
	"gin-gorm-diskominfo-diy/utils"
)

type UserService interface {
	Register(email, password string) (*domain.Users, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(email, password string) (*domain.Users, error) {
	// 1. Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.Users{
		Email:        &email,
		PasswordHash: hashedPassword,
	}

	// 2. Simpan ke DB
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, errors.New("email sudah terdaftar atau terjadi kesalahan database")
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	// 1. Cari user berdasarkan email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	// 2. Verifikasi password hash
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("email atau password salah")
	}

	// 3. Generate JWT Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("gagal membuat token autentikasi")
	}

	return token, nil
}
