package types

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct{
	UserName string `json:"username"`
	PasswordHash string`json:"password"`
}

func NewUser(registerUser RegisterUser) (User, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)
	if err != nil {
		return User{}, err
	}
	return User{UserName: registerUser.Username, PasswordHash: string(hashedPassword)}, nil
}

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))

	return err == nil
}

func CreateToken(user User) string {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user": user.UserName,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := "secret"

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenString
}