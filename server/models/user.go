package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// User the user of our application
type User struct {
	gorm.Model

	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Profile  Profile
}

// Profile the users profile
type Profile struct {
	gorm.Model

	UserID    uint
	NickName  string `gorm:"size:128"`
	FirstName string `gorm:"size:128"`
	LastName  string `gorm:"size:128"`
	Note      string `gorm:"type:TEXT"`
}

// DisplayName returns either the nickname if it's not empty, or the First+Last Name
func (p *Profile) DisplayName() string {
	if strings.TrimSpace(p.NickName) != "" {
		return p.NickName
	}

	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
