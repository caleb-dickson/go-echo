package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" validation:"required"`
	LastName  string `json:"lastName" validation:"required"`
	Email     string `json:"email" validation:"required,email" gorm:"unique"`
	Username  string `json:"username" validation:"required" gorm:"unique"`
	Password  string `json:"password"`
	RoleId    uint   `json:"roleId"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// SetPasswordHashed sets the User.Password to a bcrypt hashed value of the string input
func (user *User) SetPasswordHashed(pw string) {
	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	user.Password = string(encryptedPass)
}

// AuthenticatePassword compares the string argument to the User.Password.
// If they do not match, an error is returned
func (user *User) AuthenticatePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
}

func (user *User) RedactPassword() {
	user.Password = ""
}

func (user *User) InsertRolePlaceholderText() {
	user.Role.Name = "This is a placeholder. Retrieve this User directly to access the Role."
}

func (user *User) Count(db *gorm.DB) int64 {
	var totalUsers int64
	db.Model(&User{}).Count(&totalUsers)

	return totalUsers
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Limit(limit).Offset(offset).Preload("Role.Permissions").Find(&users)

	var redactedUsers []User
	for _, user := range users {
		user.RedactPassword()
		redactedUsers = append(redactedUsers, user)
	}

	return redactedUsers
}
