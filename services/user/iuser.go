package user

type IUser interface {
	Login(username string, password string) error
}
