package interfaces

import "github.com/jinzhu/gorm"

type Validation struct {
	Value string
	Valid string
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}
