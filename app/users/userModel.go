package users

type Roles int

const(
	ROLES_UNKNOWN Roles = iota
	ROLES_ADMIN
	ROLES_USER
)

type UserModel struct{
	UUID string
	Username string
	Password string
}