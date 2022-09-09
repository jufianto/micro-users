package model

type Roles int

const (
	UNKNOWN Roles = iota
	ADMIN
	USER
)

type Users struct {
	UUID     string
	Username string
	Password string
	Role     Roles
}