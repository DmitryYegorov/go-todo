package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"

	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
)

const (
	salt       = "dwjfhwdkj"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)

	hashPassword := generatePasswordHash(password)

	if err != nil || user.Password != hashPassword {
		logrus.Error("User not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
