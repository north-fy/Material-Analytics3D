package user

import "errors"

const (
	AccessUnknown = iota
	AccessUser
	AccessVIP
	AccessExecutor
)

var (
	errWrongData = errors.New("wrong login or password")
)

type AccessType struct {
	// 0 - unknown
	// 1 - user
	// 2 - vip
	// 3 - executor
	Access int
}
type User struct {
	Login    string
	Password string
	Access   AccessType
}
