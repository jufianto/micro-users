package users

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/google/uuid"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) CreateUser(user UserModel) (UserModel, error) {

	user.UUID = uuid.New().String()
	user.Password = generateHash(user.Password)

	// TODO: save to db

	return user, nil
}

func generateHash(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))

	hashPassString := hex.EncodeToString(hasher.Sum(nil))
	return hashPassString
}
