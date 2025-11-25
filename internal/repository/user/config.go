package user

type AccessType struct {
	IsCreator bool
	IsPremium bool
	IsUser    bool
}
type User struct {
	Email    string
	Password string
	Access   AccessType
}
