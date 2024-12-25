package types

import "golang.org/x/crypto/bcrypt"

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